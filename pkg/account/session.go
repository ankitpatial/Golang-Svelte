package account

import (
	"context"
	"fmt"
	"roofix/config"
	"roofix/ent"
	"roofix/ent/partner"
	"roofix/ent/partnercontact"
	"roofix/ent/usersession"
	"roofix/pkg/enum"
	"roofix/pkg/msg"
	"roofix/pkg/util/crypt"
	"roofix/pkg/util/log"
	"time"
)

func CtxWithUser(ctx context.Context, u *User) context.Context {
	return context.WithValue(ctx, config.UserCtxKey, u)
}

// CtxUser will read the session user info save in ctx
func CtxUser(ctx context.Context) *User {
	u, _ := ctx.Value(config.UserCtxKey).(*User)
	return u
}

func CtxUserID(ctx context.Context) string {
	u := CtxUser(ctx)
	if u != nil {
		return u.ID
	}

	return ""
}

func CtxUserNameEmail(ctx context.Context) string {
	u := CtxUser(ctx)
	if u != nil {
		return fmt.Sprintf("%s %s<%s>", u.FirstName, u.LastName, u.Email)
	}

	return ""
}

func CtxIP(ctx context.Context) string {
	ip, _ := ctx.Value("RealIP").(string)
	return ip
}

func CtxApiUsername(ctx context.Context) string {
	claims, _ := ctx.Value(config.APIUserCtxKey).(map[string]string)
	if name, ok := claims["apiUName"]; ok {
		return name
	}

	return ""
}

func CtxApiUserID(ctx context.Context) string {
	claims, _ := ctx.Value(config.APIUserCtxKey).(map[string]string)
	if id, ok := claims["apiUID"]; ok {
		return id
	}

	return ""
}

func AuthToken(sid string, d time.Duration) (string, error) {
	// sign token
	return crypt.JWTSign(map[string]interface{}{
		"sid": sid,
	}, d)
}

func ParseAuthToken(t string) (string, error) {
	claims, err := crypt.JWTParse(t)
	if err != nil {
		return "", err
	}

	return claims["sid"].(string), nil
}

// NewSession will save user-session details to DB
func NewSession(ctx context.Context, userId string, role enum.Role, expiresAt time.Time, ip string) (string, error) {
	db := ent.GetClient()
	defer db.CloseClient()

	ins := db.UserSession.
		Create().
		SetUserID(userId).
		SetExpiresAt(expiresAt).
		SetIP(ip)

	if role != enum.RoleAdmin { // pull linked partner
		p := linkedPartnerInfo(ctx, userId)
		if p != nil {
			if p.Status == enum.PartnerStatusInActive {
				return "", msg.AsError(msg.PartnerAccountNotActive)
			}

			ins.SetPartnerID(p.ID).SetPartnerContactID(p.ContactID)
		}
	}

	s, err := ins.Save(ctx)

	if err != nil {
		log.Error(err)
		return "", msg.AsError(msg.FailedToSave, "UserSession")
	}

	return s.ID, nil
}

func SessionUser(ctx context.Context, sid string) *User {
	db := ent.GetClient()
	s, err := db.UserSession.Query().
		WithUser(func(u *ent.UserQuery) {
			u.Select(sessionUserField...)
		}).
		WithPartner(func(p *ent.PartnerQuery) {
			p.Select(
				partner.FieldID,
				partner.FieldName,
				partner.FieldType,
				partner.FieldStatus,
				partner.FieldMobileAppSettings,
			)
		}).
		WithPartnerContact(func(pc *ent.PartnerContactQuery) {
			pc.Select(partnercontact.FieldID, partnercontact.FieldRole, partnercontact.FieldType)
		}).
		Where(usersession.IDEQ(sid)).
		Select(usersession.FieldExpiresAt).
		First(ctx)
	if s == nil || ent.IsNotFound(err) { // user session info not found
		return nil
	}

	if time.Now().After(s.ExpiresAt) { // session expired
		_ = db.UserSession.DeleteOneID(sid) // remove expired entry from DB.
		return nil
	}

	u := s.Edges.User
	// check account is still active
	if enum.AccountStatusActive != u.Status {
		// remove session as account is not active
		_ = db.UserSession.DeleteOneID(sid)
		return nil
	}

	var info *UserPartnerInfo
	if s.Edges.Partner != nil {
		p := s.Edges.Partner
		info = &UserPartnerInfo{
			ID:                p.ID,
			Type:              p.Type,
			PartnerName:       p.Name,
			Status:            p.Status,
			MobileAppSettings: p.MobileAppSettings,
		}

		pc := s.Edges.PartnerContact
		if pc != nil {
			info.ContactID = pc.ID
			info.IsPrimary = pc.Type == enum.PartnerContactPrimary
			info.ContactType = pc.Type
			info.Role = pc.Role
		}
	}

	t, _ := crypt.Encrypt(sid)

	res := &User{
		ID:        u.ID,
		Email:     u.Email,
		FirstName: u.FirstName,
		LastName:  u.LastName,
		Phone:     u.Phone,
		Role:      u.Role,
		Status:    u.Status,
		Picture:   u.Picture,
		Partner:   info,
		Token:     t,
		IsAdmin:   u.Role == enum.RoleAdmin,
	}

	if res.Partner != nil {
		res.IsCompanyAdmin = res.Partner.Role == enum.PartnerContactRoleAdmin
	}

	return res
}

// Info will return context user.ID
// isAdmin = true if ctx USER role is ADMIN, Admin user's will not have any link with partner account sot pID will nil.
// pID is linked to a partner as core contact then partner.ID
func Info(ctx context.Context) (isAdmin bool, uID, pID *string, pType *enum.Partner) {
	u := CtxUser(ctx)
	if u == nil {
		return
	}

	isAdmin = u.IsAdmin
	uID = &u.ID
	if u.IsCompanyAdmin && u.Partner != nil { // requester
		pID = &u.Partner.ID
		pType = &u.Partner.Type
	}
	return
}

func linkedPartnerInfo(ctx context.Context, uID string) *UserPartnerInfo {
	// pull partner info
	o, err := ent.GetClient().PartnerContact.Query().
		WithPartner(func(p *ent.PartnerQuery) {
			p.Select(
				partner.FieldType,
				partner.FieldName,
				partner.FieldStatus,
				partner.FieldMobileAppSettings,
			)
		}).
		Where(partnercontact.UserID(uID)).
		Only(ctx)

	if err != nil {
		if !ent.IsNotFound(err) {
			log.Error(err)
		}

		return nil
	}

	p := o.Edges.Partner
	return &UserPartnerInfo{
		ID:                o.PartnerID,
		Type:              p.Type,
		PartnerName:       p.Name,
		Status:            p.Status,
		MobileAppSettings: p.MobileAppSettings,

		ContactID:   o.ID,
		ContactType: o.Type,
		IsPrimary:   o.Type == enum.PartnerContactPrimary,
	}
}
