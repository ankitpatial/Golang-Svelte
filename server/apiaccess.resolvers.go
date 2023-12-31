package server

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.32

import (
	"context"
	"fmt"
	"roofix/ent"
	entApiAcceess "roofix/ent/apiaccess"
	"roofix/pkg/apiaccess"
	"roofix/pkg/audit"
	"roofix/pkg/util/str"
	"roofix/server/model"
	"strings"
)

// SaveAPIAccess is the resolver for the saveApiAccess field.
func (r *mutationResolver) SaveAPIAccess(ctx context.Context, input apiaccess.Input) (bool, error) {
	success, err := apiaccess.Save(ctx, &input)
	desc := fmt.Sprintf("saved info for %s", input.Name)
	audit.Operation(ctx, audit.ApiAccessSave, desc, err)
	return success, err
}

// UpdateAPIAccessKey is the resolver for the updateApiAccessKey field.
func (r *mutationResolver) UpdateAPIAccessKey(ctx context.Context, id string, key string) (bool, error) {
	err := apiaccess.UpdateApiAccessKey(ctx, id, key)
	desc := fmt.Sprintf("updated key for %s", id)
	audit.Operation(ctx, audit.ApiAccessUpdateKey, desc, err)
	return err == nil, err
}

// UpdateAPIAccessSecret is the resolver for the updateApiAccessKey field.
func (r *mutationResolver) UpdateAPIAccessSecret(ctx context.Context, id string, secret string) (bool, error) {
	err := apiaccess.UpdateApiAccessSecret(ctx, id, secret)
	desc := fmt.Sprintf("updated secret for %s", id)
	audit.Operation(ctx, audit.ApiAccessUpdateSecret, desc, err)
	return err == nil, err
}

// APIAccess is the resolver for the apiAccess field.
func (r *queryResolver) APIAccess(ctx context.Context, search *string, page model.PageInput) (*ent.ApiAccessConnection, error) {
	db := ent.GetClient()
	defer db.CloseClient()

	q := strings.TrimSpace(str.Val(search))
	order := &ent.ApiAccessOrder{Direction: "DESC", Field: ent.ApiAccessOrderFieldCreatedAt}
	qry := db.ApiAccess.Query()
	if q != "" {
		qry.Where(entApiAcceess.Or(
			entApiAcceess.IDEQ(q),
			entApiAcceess.UsernameContainsFold(q),
			entApiAcceess.URLContainsFold(q),
		))
	}
	return qry.Paginate(ctx, page.After, page.First, page.Before, page.Last, ent.WithApiAccessOrder(order))
}
