package server

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.32

import (
	"context"
	"fmt"
	"roofix/ent"
	"roofix/mailer"
	"roofix/pkg/account"
	"roofix/server/model"
	"roofix/template"
)

// ContactUs is the resolver for the contactUs field.
func (r *mutationResolver) ContactUs(ctx context.Context, reason string) (bool, error) {
	db := ent.GetClient()
	defer db.CloseClient()

	// session user
	u := account.CtxUser(ctx)
	name := fmt.Sprintf("%s %s", u.FirstName, u.LastName)
	// session user-partner

	// creat
	qry := db.ContactUs.
		Create().
		SetReason(reason).
		SetCreatorID(u.ID)
	// set partner
	if u.Partner != nil {
		qry.SetPartnerID(u.Partner.ID)
		name = fmt.Sprintf("%s (%s)", name, u.Partner.PartnerName)
	}
	//
	// save
	if err := qry.Exec(ctx); err != nil {
		return false, err
	}
	//
	// admin notification
	recipients := account.ToNotify(ctx, model.AdminNotifyTopicContactUsRequest.String())
	msg := &mailer.Message{
		Subject: fmt.Sprintf("Contact Us, %s", name),
		Tpl:     template.EmailContactUs,
	}
	for _, r := range recipients {
		msg.To = []string{r.Email}
		msg.Modal = map[string]interface{}{
			"name":      r.Name,
			"requester": name,
			"reason":    reason,
		}
		mailer.Send(ctx, msg)
	}

	return true, nil
}