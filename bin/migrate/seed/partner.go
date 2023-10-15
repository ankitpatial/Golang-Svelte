package seed

import (
	"context"
	"roofix/ent"
	entPartner "roofix/ent/partner"
	"roofix/pkg/enum"
	"roofix/pkg/partner"
	"roofix/pkg/util/log"
)

// Partners will seed few of the required partners into the system
func Partners(ctx context.Context) {
	//partner.ResideoID

	db := ent.GetClient()
	defer db.CloseClient()

	if exist, err := db.Partner.Query().Where(entPartner.ID(partner.EcostellaID)).Exist(ctx); err != nil {
		log.PrintError(err)
		return
	} else if !exist {
		err = db.Partner.
			Create().
			SetID(partner.EcostellaID).
			SetType(enum.PartnerIntegration).
			SetName("Ecostella").
			Exec(ctx)
		if err != nil {
			log.PrintError(err)
			return
		}
	}

	if exist, err := db.Partner.Query().Where(entPartner.ID(partner.ResideoID)).Exist(ctx); err != nil {
		log.PrintError(err)
		return
	} else if !exist {
		err = db.Partner.
			Create().
			SetID(partner.ResideoID).
			SetType(enum.PartnerIntegration).
			SetName("Resideo").
			Exec(ctx)
		if err != nil {
			log.PrintError(err)
			return
		}
	}
}
