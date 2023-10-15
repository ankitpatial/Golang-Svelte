package partner

import (
	"context"
	"roofix/ent"
	"roofix/ent/partner"
	pkgModel "roofix/pkg/model"
	"roofix/server/model"
)

func GetMobileSettings(ctx context.Context, partnerID string) (*pkgModel.MobileAppSettings, error) {
	db := ent.GetClient()
	defer db.CloseClient()

	qry := db.Partner.Query().Where(partner.ID(partnerID)).Select(partner.FieldMobileAppSettings)
	if p, err := qry.Only(ctx); err != nil {
		return nil, err
	} else {
		return &p.MobileAppSettings, nil
	}
}

func SaveMobileSettings(ctx context.Context, partnerID string, inp *model.InputMobileAppSettings) error {
	db := ent.GetClient()
	defer db.CloseClient()

	qry := db.Partner.Query().Where(partner.ID(partnerID)).Select(partner.FieldMobileAppSettings)
	p, err := qry.Only(ctx)
	if err != nil {
		return err
	}

	mas := p.MobileAppSettings
	if len(inp.HideTabs) > 0 {
		mas.HideTabs = inp.HideTabs
	}

	if inp.LogoURL != nil {
		mas.LogoURL = *inp.LogoURL
	}

	if inp.PrimaryColor != nil {
		mas.PrimaryColor = *inp.PrimaryColor
	}

	return db.Partner.
		UpdateOneID(partnerID).
		SetMobileAppSettings(mas).
		Exec(ctx)
}
