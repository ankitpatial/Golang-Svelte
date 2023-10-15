package partner

import (
	"context"
	"roofix/ent"
	"roofix/ent/partner"
	"roofix/pkg/enum"
	"roofix/pkg/msg"
	"roofix/pkg/util/validate"
)

func SaveOnboardingDone(ctx context.Context, partnerID string, step uint8) (bool, error) {
	db := ent.GetClient()
	defer db.CloseClient()

	// update partner steps completed
	c, err := db.Partner.Update().
		Where(partner.ID(partnerID), partner.StatusEQ(enum.PartnerStatusOnboarding)).
		SetSetupStepsCompleted(step).
		SetStatus(enum.PartnerStatusOnboardingDone).
		Save(ctx)

	return c > 0, err
}

func SaveStepsCompleted(ctx context.Context, partnerID string, step uint8) error {
	db := ent.GetClient()
	defer db.CloseClient()
	// update partner steps completed
	return db.Partner.Update().
		Where(partner.ID(partnerID), partner.SetupStepsCompletedLT(step)).
		SetSetupStepsCompleted(step).
		Exec(ctx)
}

// SaveBasicDetail is step one detail for partner onboarding
func SaveBasicDetail(ctx context.Context, userID, apiUserID *string, inp *BasicDetail) error {
	if err := validate.Struct(inp); err != nil {
		return err
	}

	// check name available
	if exist, err := checkNameAvailable(ctx, inp.Type, inp.Name, inp.ID); err != nil {
		return err
	} else if exist {
		return msg.AsError(msg.AlreadyExists, "Partner with same name")
	}

	db := ent.GetClient()
	defer db.CloseClient()

	exist, _ := db.Partner.Query().Where(partner.ID(inp.ID)).Exist(ctx)
	if exist { // update
		qry := db.Partner.UpdateOneID(inp.ID).
			SetName(inp.Name).
			SetNillableAddress(inp.Address).
			SetNillablePhone(inp.Phone).
			SetNillableLatitude(inp.Latitude).
			SetNillableLongitude(inp.Longitude).
			SetNillableWebsite(inp.Website).
			SetIsNationWide(inp.IsNationWide).
			SetCrewCount(uint16(inp.CrewCount)).
			SetJobCapacity(uint16(inp.JobCapacity)).
			SetNillableYearsInBusiness(inp.YearsInBusiness)

		return qry.Exec(ctx)
	}

	// create new
	qry := db.Partner.Create().
		SetID(inp.ID).
		SetNillableExternalID(inp.ExternalID).
		SetType(inp.Type).
		SetName(inp.Name).
		SetNillableAddress(inp.Address).
		SetNillablePhone(inp.Phone).
		SetNillableLatitude(inp.Latitude).
		SetNillableLongitude(inp.Longitude).
		SetNillableWebsite(inp.Website).
		SetIsNationWide(inp.IsNationWide).
		SetCrewCount(uint16(inp.CrewCount)).
		SetJobCapacity(uint16(inp.JobCapacity)).
		SetNillableCreatorID(userID).
		SetNillableCreatorAPIID(apiUserID)

	if inp.Status != nil {
		qry.SetStatus(*inp.Status)
	}

	return qry.Exec(ctx)
}
