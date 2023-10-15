package partner

import (
	"context"
	"roofix/ent"
	"roofix/pkg/util/boolean"
	"roofix/pkg/util/num"
	"roofix/server/model"
)

func SaveOperation(ctx context.Context, partnerID string, inp *model.PartnerOperationInput) error {
	db := ent.GetClient()
	defer db.CloseClient()

	err := db.Partner.UpdateOneID(partnerID).ClearFinanceOptions().ClearEpcOptions().Exec(ctx)
	if err != nil {
		return err
	}

	qry := db.Partner.UpdateOneID(partnerID)

	// What finance options do you use?
	cv := num.ReadIntPtr(inp.SalesVolume)
	if cv > 0 {
		qry.SetSalesVolume(cv)
	} else {
		qry.ClearSalesVolume()
	}

	// Down Payment %
	dp := num.ReadIntPtr(inp.DownPayment)
	if dp > 0 {
		qry.SetDownPayment(dp)
	} else {
		qry.ClearDownPayment()
	}

	// PIF Date, 1-30 days
	pif := num.ReadIntPtr(inp.PifDate)
	if pif > 0 {
		qry.SetPif(pif)
	} else {
		qry.ClearPif()
	}

	// Do you install your solar projects in house?
	if inp.InstallInHouse != nil {
		qry.SetInstallInHouse(*inp.InstallInHouse)
	} else {
		qry.ClearInstallInHouse()
	}

	// Finance options
	if len(inp.FinanceOptions) > 0 {
		qry.AddFinanceOptionIDs(inp.FinanceOptions...)
	}

	// EPC options
	if len(inp.EpcOptions) > 0 && !boolean.Read(inp.InstallInHouse) { // save only if "install your solar projects in house" = No
		qry.AddEpcOptionIDs(inp.EpcOptions...)
	}

	return qry.Exec(ctx)
}
