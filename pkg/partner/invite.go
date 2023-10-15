/*
 * Copyright (c) Simsaw Software Pvt. Ltd. 2023.
 * Author: Ankit Patial <ankit@simsaw.com>
 */

package partner

import (
	"context"
	"roofix/ent"
	"roofix/mailer"
	"roofix/pkg/account"
	"roofix/pkg/enum"
	"roofix/template"
	"time"
)

// SendInvite to on-board a partner to our system
func SendInvite(ctx context.Context, ctxUserID string, inp *Invite) (*Invite, error) {
	// create partner
	s := enum.PartnerStatusOnboarding
	detail := &BasicDetail{
		ID:     inp.ID,
		Type:   inp.Type,
		Name:   inp.CompanyName,
		Status: &s,
	}
	if err := SaveBasicDetail(ctx, &ctxUserID, nil, detail); err != nil {
		return inp, err
	} else {
		now := time.Now()
		inp.CreatedAt = &now
	}

	// create partner contact(Account Manager), activate to use and fill partner info
	contacts := []*ContactUserInput{
		{
			ID:        inp.ContactID,
			Type:      enum.PartnerContactPrimary,
			UserID:    inp.UserID,
			FirstName: inp.FirstName,
			LastName:  inp.LastName,
			Phone:     inp.Phone,
			Email:     inp.Email,
		},
	}

	if pc, err := SaveContacts(ctx, inp.ID, contacts); err != nil {
		return inp, err
	} else {
		var mgr *ent.PartnerContact
		for _, c := range pc {
			if c.Type == enum.PartnerContactPrimary {
				mgr = c
				break
			}
		}

		if mgr != nil {
			inp.ContactID = &mgr.ID
			inp.UserID = mgr.UserID
		}
	}

	// send invite email with details & ability to set password and login to system to complete onboarding.
	mailer.Send(ctx, &mailer.Message{
		To:      []string{inp.Email},
		Subject: "Roofix Partner Account",
		Tpl:     template.EmailPartnerOnBoarding,
		Modal: map[string]interface{}{
			"companyName": inp.CompanyName,
			"firstName":   inp.FirstName,
			"lastName":    inp.LastName,
			"email":       inp.Email,
			"link":        account.SetPasswordLink(ctx, inp.UserID), // link to set acc pwd
		},
	})

	return inp, nil
}
