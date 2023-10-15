/*
 * Copyright (c) Simsaw Software Pvt. Ltd. 2023.
 * Author: Ankit Patial <ankit@simsaw.com>
 */

package partner

import (
	"context"
	"roofix/ent"
	"roofix/ent/partner"
	"roofix/ent/partnerservice"
	"roofix/ent/partnerservicecity"
	"roofix/ent/partnerservicestate"
	"roofix/pkg/util/validate"
)

func SaveService(ctx context.Context, id, partnerID string, service Service, active bool) error {
	db := ent.GetClient()
	defer db.CloseClient()

	ps, err := db.PartnerService.Query().
		Where(partnerservice.ID(id)).
		Only(ctx)
	if err != nil && !ent.IsNotFound(err) {
		return err
	}

	if ps != nil {
		return db.PartnerService.UpdateOneID(id).
			SetActive(active).
			Exec(ctx)
	}

	return db.PartnerService.Create().
		SetID(id).
		SetPartnerID(partnerID).
		SetServiceID(service.ToUInt8()).
		SetActive(active).
		Exec(ctx)
}

func SaveServiceState(ctx context.Context, inp *ServiceState) error {
	if err := validate.Struct(inp); err != nil {
		return err
	}

	db := ent.GetClient()
	defer db.CloseClient()

	pss, err := db.PartnerServiceState.Query().
		Where(
			partnerservicestate.HasPartnerWith(partner.ID(inp.PartnerID)),
			partnerservicestate.Country(inp.Country),
			partnerservicestate.State(inp.State),
		).Only(ctx)
	if err != nil && !ent.IsNotFound(err) {
		return err
	}

	if pss != nil { // update
		var in ent.UpdatePartnerServiceStateInput
		in.LicenseNo = inp.LicenseNo
		in.LicenseExpDate = inp.LicenseExpDate
		in.ProofDocID = inp.ProofDocID
		return db.PartnerServiceState.UpdateOneID(pss.ID).SetInput(in).Exec(ctx)
	}

	// create
	var in ent.CreatePartnerServiceStateInput
	in.PartnerID = inp.PartnerID
	in.Country = inp.Country
	in.State = inp.State
	in.LicenseNo = inp.LicenseNo
	in.LicenseExpDate = inp.LicenseExpDate
	in.ProofDocID = inp.ProofDocID
	return db.PartnerServiceState.Create().SetInput(in).Exec(ctx)

}

func SaveServiceCity(ctx context.Context, inp *ServiceCity) error {
	if err := validate.Struct(inp); err != nil {
		return err
	}

	db := ent.GetClient()
	defer db.CloseClient()

	pss, err := db.PartnerServiceCity.Query().Where(
		partnerservicecity.HasPartnerWith(partner.ID(inp.PartnerID)),
		partnerservicecity.PostalID(inp.PostalID),
	).Only(ctx)
	if err != nil && !ent.IsNotFound(err) {
		return err
	}

	if pss != nil { // update
		var in ent.UpdatePartnerServiceCityInput
		in.Active = inp.Active
		in.LicenseNo = inp.LicenseNo
		in.ProofDocID = inp.ProofDocID
		return db.PartnerServiceCity.UpdateOneID(pss.ID).SetInput(in).Exec(ctx)
	}

	// create
	var in ent.CreatePartnerServiceCityInput
	in.PartnerID = inp.PartnerID
	in.PostalID = inp.PostalID
	in.Active = inp.Active
	in.LicenseNo = inp.LicenseNo
	in.ProofDocID = inp.ProofDocID
	return db.PartnerServiceCity.Create().SetInput(in).Exec(ctx)
}
