package partner

import (
	"context"
	"roofix/ent"
	"roofix/ent/partner"
	"roofix/pkg/enum"
)

// checkNameAvailable on Create or Update partner record
func checkNameAvailable(ctx context.Context, t enum.Partner, name, skipID string) (bool, error) {
	db := ent.GetClient()
	defer db.CloseClient()

	qry := db.Partner.Query().Where(partner.NameEqualFold(name), partner.TypeEQ(t))
	if skipID != "" {
		qry.Where(partner.IDNEQ(skipID))
	}

	return qry.Exist(ctx)
}
