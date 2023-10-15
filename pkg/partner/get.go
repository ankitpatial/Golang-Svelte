/*
 * Copyright (c) Simsaw Software Pvt. Ltd. 2022.
 * Author: Ankit Patial <ankit@simsaw.com>
 */

package partner

import (
	"context"
	"errors"
	"roofix/ent"
	"roofix/ent/optionlist"
	"roofix/ent/partner"
	"roofix/ent/partnercontact"
	"roofix/ent/user"
	"roofix/pkg/enum"
	"roofix/pkg/msg"
	"roofix/server/model"
)

// ByID a partner detail
func ByID(ctx context.Context, id string, pType *enum.Partner) (*ent.Partner, error) {
	qry := ent.GetClient().Partner.Query().
		WithPartnerContacts(func(q *ent.PartnerContactQuery) {
			q.WithUser()
			q.Order(ent.Asc(partnercontact.FieldCreatedAt))
		}).
		Where(partner.ID(id))

	if pType != nil && *pType == enum.PartnerSolar {
		qry.WithFinanceOptions(func(op *ent.OptionListQuery) {
			op.Select(optionlist.FieldDisplayName)
		})
		qry.WithEpcOptions(func(op *ent.OptionListQuery) {
			op.Select(optionlist.FieldDisplayName)
		})
	}

	return qry.Only(ctx)
}

// WithFinanceAndEpcOptions will fetch partner edges: FinanceOptions & EpcOptions
func WithFinanceAndEpcOptions(ctx context.Context, id string) (*ent.Partner, error) {
	qry := ent.GetClient().Partner.Query().Where(partner.ID(id)).
		WithFinanceOptions(func(op *ent.OptionListQuery) {
			op.Select(optionlist.FieldDisplayName)
		}).
		WithEpcOptions(func(op *ent.OptionListQuery) {
			op.Select(optionlist.FieldDisplayName)
		})
	return qry.Select(partner.FieldID).Only(ctx)
}

// Kind will return partner's type for given ID
func Kind(ctx context.Context, id string) (enum.Partner, error) {
	p, err := ent.GetClient().Partner.Query().Where(partner.ID(id)).Select(partner.FieldType).Only(ctx)
	if err != nil && ent.IsNotFound(err) {
		return "", msg.AsError(msg.NotFound, "partner")
	}

	return p.Type, err
}

// List partner order by name
func List(ctx context.Context, page model.PageInput, where *ent.PartnerWhereInput) (*ent.PartnerConnection, error) {
	return ent.GetClient().Partner.Query().
		Paginate(
			ctx,
			page.After, page.First, page.Before, page.Last,
			ent.WithPartnerFilter(where.Filter),
			ent.WithPartnerOrder(&ent.PartnerOrder{
				Direction: "ASC",
				Field:     ent.PartnerOrderFieldName,
			}),
		)
}

// ByUserID will return associated partner detail
func ByUserID(ctx context.Context, userID string) (*ent.Partner, error) {
	p, err := ent.GetClient().Partner.Query().
		Where(partner.HasContactsWith(user.ID(userID))).
		Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, errors.New("failed to find Partner Details")
		}
		if ent.IsNotSingular(err) {
			return nil, errors.New("found more than one Partner Details")
		}
		return nil, err
	}

	return p, nil
}
