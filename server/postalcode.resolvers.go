package server

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.32

import (
	"context"
	"fmt"
	"roofix/ent"
	"roofix/ent/postalcode"
	"roofix/pkg/postal"
	"roofix/server/model"
	"strings"

	"entgo.io/ent/dialect/sql"
)

// MarkServiceArea is the resolver for the markServiceArea field.
func (r *mutationResolver) MarkServiceArea(ctx context.Context, id string, value bool) (bool, error) {
	db := ent.GetClient()
	defer db.CloseClient()

	err := db.PostalCode.UpdateOneID(id).SetServiceArea(value).Exec(ctx)
	return err == nil, err
}

// States is the resolver for the states field.
func (r *queryResolver) States(ctx context.Context, q string) ([]*model.State, error) {
	var result []*model.State
	res := postal.Search(q, 10)
	if len(res) > 0 {
		for _, s := range res {
			if len(s.Cities) == 0 {
				continue
			}
			var cities []*model.City
			for _, c := range s.Cities {
				cities = append(cities, &model.City{
					ID:   c.Zip,
					Name: &c.Name,
					Zip:  &c.Zip,
				})
			}

			result = append(result, &model.State{
				ID:     s.ID,
				Name:   s.Name,
				Cities: cities,
			})
		}
	}

	return result, nil
}

// ServiceStates is the resolver for the serviceStates field.
func (r *queryResolver) ServiceStates(ctx context.Context, q string) ([]*model.State, error) {
	db := ent.GetClient()
	defer db.CloseClient()

	postalList, err := db.PostalCode.
		Query().
		Where(
			postalcode.Country(postal.CountryUS),
			postalcode.ServiceArea(true),
			postalcode.Or(
				postalcode.StateAbr(q),
				postalcode.StateContainsFold(q),
			),
		).
		Limit(10).
		Order(ent.Asc(postalcode.FieldState), ent.Asc(postalcode.FieldCity)).
		All(ctx)
	if err != nil {
		return nil, err
	}

	return toStateList(postalList), err
}

// Cities is the resolver for the cities field.
func (r *queryResolver) Cities(ctx context.Context, state string, q string, skip int, take int) ([]*model.City, error) {
	db := ent.GetClient()
	defer db.CloseClient()

	qry := db.PostalCode.Query().Unique(true).Where(postalcode.State(state))

	if strings.TrimSpace(q) != "" {
		qry.Where(postalcode.Or(
			postalcode.CityContainsFold(q),
			postalcode.CodeContainsFold(q),
		))
	}

	var res []*model.City
	err := qry.
		Where(func(pc *sql.Selector) {
			pc.Select(
				pc.C(postalcode.FieldID),
				fmt.Sprintf("%s as name", pc.C(postalcode.FieldCity)),
				fmt.Sprintf("%s as zip", pc.C(postalcode.FieldCode)),
			)
		}).
		Offset(skip).
		Limit(take).
		Order(ent.Asc(postalcode.FieldCity)).
		Select().
		Scan(ctx, &res)
	return res, err
}

// AllServiceAreas is the resolver for the allServiceAreas field.
func (r *queryResolver) AllServiceAreas(ctx context.Context) ([]*model.State, error) {
	db := ent.GetClient()
	defer db.CloseClient()

	postalList, err := db.PostalCode.
		Query().
		Where(
			postalcode.Country(postal.CountryUS),
			postalcode.ServiceArea(true),
		).
		All(ctx)
	if err != nil {
		return nil, err
	}

	return toStateList(postalList), nil
}
