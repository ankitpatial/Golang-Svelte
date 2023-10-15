/*
 * Copyright (c) Simsaw Software Pvt. Ltd. 2023.
 * Author: Ankit Patial <ankit@simsaw.com>
 */

package server

import (
	"context"
	"roofix/ent"
	"roofix/pkg/account"
	"roofix/server/model"
	"time"
)

type ServiceArea struct {
	StateAbr            string     `json:"state_abr"`
	State               string     `json:"state"`
	StateLicenseNo      *string    `json:"state_lic_no"`
	StateLicenseExpDate *time.Time `json:"state_lic_exp"`
	StateProofID        *string    `json:"state_proof"`

	ID           string  `json:"id"` // postalID
	City         string  `json:"city"`
	CityZip      string  `json:"city_zip"`
	Active       bool    `json:"active"`
	LicenseNo    *string `json:"license_no"`
	LicenseProof *string `json:"proof_doc_id"`
}

func isAdminOrCompanyAdmin(ctx context.Context, partnerID string) bool {
	u := account.CtxUser(ctx)
	if u == nil {
		return false
	}

	// must match session user's partnerID
	if u.IsCompanyAdmin && u.Partner.ID != partnerID {
		return false
	}

	return u.IsAdmin || u.IsCompanyAdmin
}

func optionListIDs(options []*ent.OptionList) []string {
	ids := make([]string, 0, len(options))
	for _, i := range options {
		ids = append(ids, i.ID)
	}
	return ids
}

func optionListAsEntities(options []*ent.OptionList) []*model.Entity {
	ids := make([]*model.Entity, 0, len(options))
	for _, i := range options {
		ids = append(ids, &model.Entity{ID: i.ID, Name: i.DisplayName})
	}
	return ids
}
