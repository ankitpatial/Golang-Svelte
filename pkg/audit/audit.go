package audit

import (
	"context"
	"roofix/ent"
	"roofix/pkg/account"
	"roofix/pkg/util/log"
)

type Action string

const (
	AccLogin               Action = "account:login"
	AccLogout              Action = "account:logout"
	AccCreate              Action = "account:create"
	AccUpdate              Action = "account:update"
	AccForgotPassword      Action = "account:forgot-pwd"
	AccChangePassword      Action = "account:change-pwd"
	ApiUserCreate          Action = "apiuser:create"
	ApiUserChangePwd       Action = "apiuser:change-pwd"
	ApiAccessSave          Action = "apiaccess:save"
	ApiAccessUpdateKey     Action = "apiaccess:update-key"
	ApiAccessUpdateSecret  Action = "apiaccess:update-secret"
	OAuthPasswordGrant     Action = "oauth:token"
	OAuthRefreshTokenGrant Action = "oauth:refresh"
	JobDetail              Action = "job:detail"
	JobCreate              Action = "job:create"
	JobUpdate              Action = "job:update"
	JobUpdateStatus        Action = "job:update-status"
	FileDelivery           Action = "file:delivery"

	UserSave Action = "user:save"

	EstRequest         Action = "est:request"
	EstApproved        Action = "est:approved"
	EstEagleView       Action = "est:eagleview"
	EstEagleViewUpdate Action = "est:eagleview-update"
	EstRFXUpdate       Action = "est:rfx-update"

	PartnerCreate       Action = "partner:create"
	PartnerSave         Action = "partner:save"
	PartnerActive       Action = "partner:active"
	PartnerInActive     Action = "partner:inactive"
	PartnerSaveContacts Action = "partner:save-contacts"
)

func Operation(ctx context.Context, action Action, description string, err error) {
	op(ctx, action, account.CtxUserID(ctx), account.CtxApiUserID(ctx), description, err)
}

func OpSuccess(ctx context.Context, action Action, desc string) {
	op(ctx, action, account.CtxUserID(ctx), account.CtxApiUserID(ctx), desc, nil)
}

func OpFailed(ctx context.Context, action Action, err error) {
	op(ctx, action, account.CtxUserID(ctx), account.CtxApiUserID(ctx), "", err)
}

// UserOpSuccess is supposed to where we have user ID but not the user session
func UserOpSuccess(ctx context.Context, userID string, action Action, desc string) {
	op(ctx, action, userID, "", desc, nil)
}

// UserOpFailed is supposed to where we have user ID but not the user session
func UserOpFailed(ctx context.Context, userID string, action Action, err error) {
	op(ctx, action, userID, "", "", err)
}

func ApiOpSuccess(ctx context.Context, action Action, apiUserID string, description string) {
	op(ctx, action, "", apiUserID, description, nil)
}

func ApiOpFailed(ctx context.Context, action Action, apiUserID string, err error) {
	op(ctx, action, "", apiUserID, "", err)
}

func op(ctx context.Context, action Action, uID, apiUID string, description string, err error) {
	cl := ent.GetClient().AuditLog
	save(ctx, cl, action, uID, apiUID, description, err)
}

func save(ctx context.Context, cl *ent.AuditLogClient, action Action, uID, apiUID string, description string, err error) {
	qry := cl.Create().
		SetAction(string(action)).
		SetDescription(desc(description, err))

	if uID != "" {
		qry.SetUserID(uID)
	}

	if apiUID != "" {
		qry.SetAPIUserID(apiUID)
	}

	if ip := account.CtxIP(ctx); ip != "" {
		qry.SetIP(ip)
	}

	if _, err2 := qry.Save(ctx); err2 != nil {
		log.Error(err2)
	}
}

func desc(d string, err error) string {
	var res string
	if err != nil {
		res = "ERROR " + err.Error()
	} else {
		res = d
	}

	if len(res) > 500 { // trim
		return res[:497] + "..."
	}

	return res
}
