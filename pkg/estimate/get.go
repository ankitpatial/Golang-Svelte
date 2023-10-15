package estimate

import (
	"context"
	"roofix/ent"
	"roofix/ent/apiuser"
	"roofix/ent/estimate"
	"roofix/ent/user"
)

func GetByID(ctx context.Context, id string) (*ent.Estimate, error) {
	db := ent.GetClient()
	defer db.CloseClient()

	return db.Estimate.Query().Where(estimate.ID(id)).First(ctx)
}

// Get estimate  by its id linked to current session user.
func Get(ctx context.Context, id string) (*ent.Estimate, error) {
	db := ent.GetClient()
	defer db.CloseClient()

	qry := db.Debug().Estimate.Query().
		WithSalesRep().
		WithHomeOwner().
		WithPartner().
		WithCreator(func(c *ent.UserQuery) { c.Select(user.FieldID, user.FieldFirstName, user.FieldLastName) }).
		WithCreatorAPI(func(api *ent.ApiUserQuery) { api.Select(apiuser.FieldUsername) }).
		WithPdf().
		Where(estimate.ID(id))
	return accessFilter(ctx, qry).First(ctx)
}
