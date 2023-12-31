package server

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.32

import (
	"context"
	"fmt"
	"roofix/ent"
	"roofix/ent/notifysetting"
	entPartner "roofix/ent/partner"
	"roofix/ent/partnerservice"
	"roofix/ent/partnerservicecity"
	"roofix/ent/partnerservicestate"
	"roofix/ent/postalcode"
	"roofix/ent/user"
	"roofix/mailer"
	"roofix/pkg/account"
	"roofix/pkg/audit"
	"roofix/pkg/document"
	"roofix/pkg/enum"
	"roofix/pkg/msg"
	"roofix/pkg/partner"
	"roofix/pkg/postal"
	"roofix/pkg/util/crypt"
	"roofix/pkg/util/log"
	"roofix/pkg/util/str"
	"roofix/pkg/util/uid"
	"roofix/server/generated"
	"roofix/server/model"
	"roofix/template"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
)

// SavePartner is the resolver for the savePartner field.
func (r *mutationResolver) SavePartner(ctx context.Context, input partner.BasicDetail) (bool, error) {
	if !isAdminOrCompanyAdmin(ctx, input.ID) {
		return false, msg.AsError(msg.NotAuthorized)
	}

	u := account.CtxUser(ctx)
	action := audit.PartnerSave

	if u.Role == enum.RoleAdmin {
		s := enum.PartnerStatusActive
		input.Status = &s // will be used on create in partner.SaveBasicDetail
	}

	if err := partner.SaveBasicDetail(ctx, &u.ID, nil, &input); err != nil {
		audit.OpFailed(ctx, action, err)
		return false, err
	}

	audit.OpSuccess(ctx, action, "saved ID: "+input.ID)
	return true, nil
}

// InvitePartner is the resolver for the invitePartner field.
func (r *mutationResolver) InvitePartner(ctx context.Context, input partner.Invite) (*partner.Invite, error) {
	return partner.SendInvite(ctx, account.CtxUserID(ctx), &input)
}

// SetPartnerActive is the resolver for the setPartnerActive field.
func (r *mutationResolver) SetPartnerActive(ctx context.Context, partnerID string, value bool) (bool, error) {
	err := partner.SetActiveStatus(ctx, partnerID, value)
	return err == nil, err
}

// SavePartnerOperations is the resolver for the savePartnerOperations field.
func (r *mutationResolver) SavePartnerOperations(ctx context.Context, partnerID string, inp model.PartnerOperationInput) (bool, error) {
	if !isAdminOrCompanyAdmin(ctx, partnerID) {
		return false, msg.AsError(msg.NotAuthorized)
	}

	if err := partner.SaveOperation(ctx, partnerID, &inp); err != nil {
		log.Error(err)
		return false, err
	}

	return true, nil
}

// SavePartnerCompletedSteps is the resolver for the savePartnerCompletedSteps field.
func (r *mutationResolver) SavePartnerCompletedSteps(ctx context.Context, partnerID string, step int, done *bool) (bool, error) {
	if !isAdminOrCompanyAdmin(ctx, partnerID) {
		return false, msg.AsError(msg.NotAuthorized)
	}

	var err error
	var sendMail bool
	doneOnBoarding := done != nil && *done
	if doneOnBoarding {
		sendMail, err = partner.SaveOnboardingDone(ctx, partnerID, uint8(step))
	} else {
		err = partner.SaveStepsCompleted(ctx, partnerID, uint8(step))
	}

	if err != nil {
		return false, err
	}

	// Email notification
	if sendMail {
		db := ent.GetClient()
		defer db.CloseClient()

		topic := crypt.MD5Hash(model.AdminNotifyTopicPartnerOnboardingDone.String())
		// let admin know about the process is done
		notifyUsers, err := db.User.Query().
			Where(user.HasNotifyWith(notifysetting.TopicID(topic), notifysetting.ReceiveEmail(true))).
			Select(user.FieldID, user.FieldEmail).
			All(ctx)

		if err != nil {
			log.Error(err)
		} else if len(notifyUsers) > 0 { // send email notification
			var to []string
			for _, u := range notifyUsers {
				to = append(to, u.Email)
			}

			u := account.CtxUser(ctx)
			if u.Partner != nil {
				mailer.Send(ctx, &mailer.Message{
					To:      to,
					Subject: "Partner Completed Onboarding",
					Tpl:     template.EmailPartnerOnBoardingDone,
					Modal: map[string]interface{}{
						"partner":  u.Partner.PartnerName,
						"pathname": fmt.Sprintf("/partner/%s/%s/save", strings.ToLower(u.Partner.Type.String()), u.Partner.ID),
					},
				})
			}
		}
	}

	return true, nil
}

// SaveServiceState is the resolver for the saveServiceState field.
func (r *mutationResolver) SaveServiceState(ctx context.Context, partnerID string, state string, licNo *string, expDate *time.Time, proofDocID *string) (bool, error) {
	if !isAdminOrCompanyAdmin(ctx, partnerID) {
		return false, msg.AsError(msg.NotAuthorized)
	}

	inp := &partner.ServiceState{
		PartnerID:      partnerID,
		Country:        postal.CountryUS,
		State:          state,
		LicenseNo:      licNo,
		LicenseExpDate: expDate,
		ProofDocID:     proofDocID,
	}

	err := partner.SaveServiceState(ctx, inp)
	return err == nil, err
}

// SaveServiceCity is the resolver for the saveServiceCity field.
func (r *mutationResolver) SaveServiceCity(ctx context.Context, partnerID string, postalID string, active *bool, licNo *string, proofDocID *string) (bool, error) {
	if !isAdminOrCompanyAdmin(ctx, partnerID) {
		return false, msg.AsError(msg.NotAuthorized)
	}

	inp := &partner.ServiceCity{
		PartnerID:  partnerID,
		PostalID:   postalID,
		Active:     active,
		LicenseNo:  licNo,
		ProofDocID: proofDocID,
	}

	err := partner.SaveServiceCity(ctx, inp)
	return err == nil, err
}

// SaveService is the resolver for the saveService field.
func (r *mutationResolver) SaveService(ctx context.Context, id string, partnerID string, service partner.Service, active bool) (bool, error) {
	err := partner.SaveService(ctx, id, partnerID, service, active)
	return err == nil, err
}

// SaveLeadTime is the resolver for the saveLeadTime field.
func (r *mutationResolver) SaveLeadTime(ctx context.Context, partnerID string, asphalt *string, metal *string, tile *string) (bool, error) {
	if !isAdminOrCompanyAdmin(ctx, partnerID) {
		return false, msg.AsError(msg.NotAuthorized)
	}

	db := ent.GetClient()
	defer db.CloseClient()

	err := db.Partner.
		UpdateOneID(partnerID).
		SetNillableAsphaltLeadT(asphalt).
		SetNillableMetalLeadT(metal).
		SetNillableTileLeadT(tile).
		Exec(ctx)

	return err == nil, err
}

// Contacts is the resolver for the contacts field.
func (r *partnerResolver) Contacts(ctx context.Context, obj *ent.Partner) ([]*ent.PartnerContact, error) {
	return obj.Edges.PartnerContacts, nil
}

// CrewCount is the resolver for the crewCount field.
func (r *partnerResolver) CrewCount(ctx context.Context, obj *ent.Partner) (int, error) {
	if obj == nil {
		return 0, nil
	}

	return int(obj.CrewCount), nil
}

// JobCapacity is the resolver for the jobCapacity field.
func (r *partnerResolver) JobCapacity(ctx context.Context, obj *ent.Partner) (int, error) {
	if obj == nil {
		return 0, nil
	}

	return int(obj.JobCapacity), nil
}

// SetupStepsCompleted is the resolver for the setupStepsCompleted field.
func (r *partnerResolver) SetupStepsCompleted(ctx context.Context, obj *ent.Partner) (int, error) {
	if obj == nil {
		return 0, nil
	}

	return int(obj.SetupStepsCompleted), nil
}

// IsActive is the resolver for the isActive field.
func (r *partnerResolver) IsActive(ctx context.Context, obj *ent.Partner) (bool, error) {
	if obj == nil {
		return false, nil
	}
	s := obj.Status

	return s == enum.PartnerStatusActive, nil
}

// FinanceOptions is the resolver for the financeOption field.
func (r *partnerResolver) FinanceOptions(ctx context.Context, obj *ent.Partner) ([]string, error) {
	if obj == nil || obj.Edges.FinanceOptions == nil {
		return []string{}, nil
	}

	return optionListIDs(obj.Edges.FinanceOptions), nil
}

// PifDate is the resolver for the pifDate field.
func (r *partnerResolver) PifDate(ctx context.Context, obj *ent.Partner) (*int, error) {
	if obj == nil {
		return nil, nil
	}

	return obj.Pif, nil
}

// EpcOptions is the resolver for the epcOption field.
func (r *partnerResolver) EpcOptions(ctx context.Context, obj *ent.Partner) ([]string, error) {
	if obj == nil || obj.Edges.EpcOptions == nil {
		return []string{}, nil
	}

	return optionListIDs(obj.Edges.EpcOptions), nil
}

// Status is the resolver for the status field.
func (r *partnerJobResolver) Status(ctx context.Context, obj *partner.Job) (string, error) {
	panic(fmt.Errorf("not implemented: Status - status"))
}

// Service is the resolver for the service field.
func (r *partnerServiceResolver) Service(ctx context.Context, obj *ent.PartnerService) (partner.Service, error) {
	if obj == nil {
		return partner.NoService, nil
	}
	var s partner.Service
	s.FromUInt8(obj.ServiceID)
	return s, nil
}

// Description is the resolver for the description field.
func (r *partnerServiceResolver) Description(ctx context.Context, obj *ent.PartnerService) (string, error) {
	if obj == nil {
		return "", nil
	}
	var s partner.Service
	s.FromUInt8(obj.ServiceID)
	return s.Description(), nil
}

// PartnerNameAvailable is the resolver for the partnerNameAvailable field.
func (r *queryResolver) PartnerNameAvailable(ctx context.Context, id string, name string, typeArg enum.Partner) (bool, error) {
	if name == "" { // empty name is not available
		return false, nil
	}

	db := ent.GetClient()
	defer db.CloseClient()

	exist, err := db.Partner.
		Query().
		Where(
			entPartner.IDNEQ(id),
			entPartner.TypeEQ(typeArg),
			entPartner.NameEqualFold(name),
		).
		Exist(ctx)
	return !exist, err
}

// Partner is the resolver for the partner field.
func (r *queryResolver) Partner(ctx context.Context, id string, typeArg *enum.Partner) (*ent.Partner, error) {
	if !isAdminOrCompanyAdmin(ctx, id) {
		return nil, msg.AsError(msg.NotAuthorized)
	}
	return partner.ByID(ctx, id, typeArg)
}

// PartnerDocs is the resolver for the partnerDocs field.
func (r *queryResolver) PartnerDocs(ctx context.Context, partnerID string, section enum.DocSection) ([]*document.Info, error) {
	if !isAdminOrCompanyAdmin(ctx, partnerID) {
		return nil, msg.AsError(msg.NotAuthorized)
	}

	return document.ByFolderDirSection(ctx, enum.FolderPartnerDocs, &partnerID, &section)
}

// Partners is the resolver for the partners field.
func (r *queryResolver) Partners(ctx context.Context, search *string, partnerType *enum.Partner, status *string, page model.PageInput) (*ent.PartnerConnection, error) {
	where := new(ent.PartnerWhereInput)

	if q := str.TrimSpace(search); q != "" {
		where.NameContainsFold = &q
	}

	if partnerType != nil {
		where.Type = partnerType
	}

	if s := str.Val(status); s != "" {
		status := enum.PartnerStatus(s)
		where.Status = &status
	}

	return partner.List(ctx, page, where)
}

// PartnerServiceStates is the resolver for the partnerServiceStates field.
func (r *queryResolver) PartnerServiceStates(ctx context.Context, partnerID string) ([]*model.ServiceState, error) {
	if !isAdminOrCompanyAdmin(ctx, partnerID) {
		return nil, msg.AsError(msg.NotAuthorized)
	}

	db := ent.GetClient()
	defer db.CloseClient()

	country := postal.CountryUS
	var areas []*ServiceArea
	err := db.Debug().PostalCode.
		Query().
		Where(
			postalcode.Country(country),
			postalcode.ServiceArea(true),
		).
		Where(func(p *sql.Selector) {
			s := sql.Table(partnerservicestate.Table)
			c := sql.Table(partnerservicecity.Table)

			// join
			p.LeftJoin(s).
				On(p.C(postalcode.FieldState), s.C(partnerservicestate.FieldState)).
				OnP(sql.P(func(builder *sql.Builder) {
					builder.WriteString(fmt.Sprintf("%s = '%s'", s.C(partnerservicestate.PartnerColumn), partnerID))
				}))

			p.LeftJoin(c).
				On(p.C(postalcode.FieldID), c.C(partnerservicecity.FieldPostalID)).
				OnP(sql.P(func(builder *sql.Builder) {
					builder.WriteString(fmt.Sprintf("%s = '%s'", c.C(partnerservicecity.PartnerColumn), partnerID))
				}))

			p.Select(
				p.C(postalcode.FieldID),
				p.C(postalcode.FieldStateAbr),
				p.C(postalcode.FieldState),
				p.C(postalcode.FieldCity),
				fmt.Sprintf("%s as city_zip", p.C(postalcode.FieldCode)),
				fmt.Sprintf("%s as state_lic_no", s.C(partnerservicestate.FieldLicenseNo)),
				fmt.Sprintf("%s as state_lic_exp", s.C(partnerservicestate.FieldLicenseExpDate)),
				c.C(partnerservicecity.FieldActive),
				c.C(partnerservicecity.FieldLicenseNo),
				c.C(partnerservicecity.FieldProofDocID),
			)
			p.OrderBy(postalcode.FieldState, postalcode.FieldCity)
		}).
		Select().
		Scan(ctx, &areas)
	if err != nil {
		return nil, err
	}

	var result []*model.ServiceState

	// enumerate
	for i, a := range areas {
		var s *model.ServiceState
		for _, r := range result {
			if r.Name == a.State {
				s = r
				break
			}
		}

		if s == nil {
			s = &model.ServiceState{
				ID:             a.StateAbr,
				Name:           a.State,
				LicenseNo:      a.StateLicenseNo,
				LicenseExpDate: a.StateLicenseExpDate,
			}
			result = append(result, s)
		}

		if i == 0 {
			s.Expand = true
		}

		s.Cities = append(s.Cities, &model.ServiceCity{
			ID:           a.ID,
			CityName:     a.City,
			CityZip:      a.CityZip,
			Active:       a.Active,
			LicenseNo:    a.LicenseNo,
			LicenseProof: a.LicenseProof,
		})

	}

	return result, nil
}

// PartnerServices is the resolver for the partnerServices field.
func (r *queryResolver) PartnerServices(ctx context.Context, partnerID string) ([]*ent.PartnerService, error) {
	if !isAdminOrCompanyAdmin(ctx, partnerID) {
		return nil, msg.AsError(msg.NotAuthorized)
	}

	db := ent.GetClient()
	defer db.CloseClient()

	res, err := db.PartnerService.
		Query().
		Where(partnerservice.HasPartnerWith(entPartner.ID(partnerID))).
		All(ctx)
	if err != nil {
		return nil, err
	}

	var result []*ent.PartnerService
	for _, s := range partner.AllServices {
		if s == partner.NoService {
			continue
		}

		var ps *ent.PartnerService
		for _, pss := range res {
			if pss.ServiceID == s.ToUInt8() {
				ps = pss
				break
			}
		}

		if ps == nil {
			ps = &ent.PartnerService{
				ID:        uid.ULID(),
				ServiceID: s.ToUInt8(),
				Active:    false,
			}
		}

		result = append(result, ps)
	}

	return result, err
}

// PartnerOptions is the resolver for the partnerOptions field.
func (r *queryResolver) PartnerOptions(ctx context.Context, partnerID string) ([]*model.Options, error) {
	if !isAdminOrCompanyAdmin(ctx, partnerID) {
		return nil, msg.AsError(msg.NotAuthorized)
	}
	// pull partner info
	p, err := partner.WithFinanceAndEpcOptions(ctx, partnerID)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return []*model.Options{
		{Type: enum.OptionListFinance, Options: optionListAsEntities(p.Edges.FinanceOptions)},
		{Type: enum.OptionListEPC, Options: optionListAsEntities(p.Edges.EpcOptions)},
	}, nil
}

// Partner returns generated.PartnerResolver implementation.
func (r *Resolver) Partner() generated.PartnerResolver { return &partnerResolver{r} }

// PartnerJob returns generated.PartnerJobResolver implementation.
func (r *Resolver) PartnerJob() generated.PartnerJobResolver { return &partnerJobResolver{r} }

// PartnerService returns generated.PartnerServiceResolver implementation.
func (r *Resolver) PartnerService() generated.PartnerServiceResolver {
	return &partnerServiceResolver{r}
}

type partnerResolver struct{ *Resolver }
type partnerJobResolver struct{ *Resolver }
type partnerServiceResolver struct{ *Resolver }
