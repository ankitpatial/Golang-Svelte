package product

import (
	"context"
	"entgo.io/contrib/entgql"
	"entgo.io/ent/dialect/sql"
	"roofix/ent"
	"roofix/ent/product"
	"roofix/ent/productpackage"
	"roofix/pkg/enum"
	"roofix/pkg/util/str"
	"roofix/server/model"
)

func List(ctx context.Context, category *enum.Product, search *string, page model.PageInput) (*ent.ProductConnection, error) {
	db := ent.GetClient()
	defer db.CloseClient()

	by := &ent.ProductOrder{
		Direction: entgql.OrderDirectionAsc,
		Field:     ent.ProductOrderFieldName,
	}

	return db.Product.Query().
		Where(func(p *sql.Selector) {
			if category != nil {
				p.Where(sql.EQ(p.C(product.FieldType), *category))
			}
			q := str.Val(search)
			if q != "" {
				p.Where(sql.ContainsFold(p.C(product.FieldName), q))
			}
		}).
		Paginate(ctx, page.After, page.First, page.Before, page.Last, ent.WithProductOrder(by))
}

func Packages(ctx context.Context, category *enum.Product, search *string, page model.PageInput) (*ent.ProductPackageConnection, error) {
	db := ent.GetClient()
	defer db.CloseClient()

	by := &ent.ProductPackageOrder{
		Direction: entgql.OrderDirectionAsc,
		Field:     ent.ProductPackageOrderFieldName,
	}

	return db.ProductPackage.Query().
		Where(func(p *sql.Selector) {
			if category != nil {
				p.Where(sql.EQ(p.C(productpackage.FieldType), *category))
			}
			q := str.Val(search)
			if q != "" {
				p.Where(sql.ContainsFold(p.C(productpackage.FieldName), q))
			}
		}).
		Paginate(ctx, page.After, page.First, page.Before, page.Last, ent.WithProductPackageOrder(by))
}
