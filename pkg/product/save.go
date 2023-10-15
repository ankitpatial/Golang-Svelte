package product

import (
	"context"
	"roofix/ent"
	"roofix/ent/product"
	"roofix/ent/productpackage"
	"roofix/pkg/util/log"
	"roofix/server/model"
)

func Save(ctx context.Context, uID string, inp *model.ProductInput) error {
	db := ent.GetClient()
	defer db.CloseClient()

	exist, err := db.Product.Query().Where(product.ID(inp.ID)).Exist(ctx)
	if err != nil {
		log.Error(err)
		return err
	}

	if exist {
		return db.Product.UpdateOneID(inp.ID).
			SetName(inp.Name).
			SetType(inp.Category).
			SetType(inp.Category).
			SetDescription(inp.Description).
			SetUnitPrice(inp.Price).
			SetFeatures(inp.Features).
			SetNillableSpecialNote(inp.SpecialNote).
			SetImageID(inp.ImageID).
			Exec(ctx)
	}

	return db.Product.Create().
		SetID(inp.ID).
		SetType(inp.Category).
		SetName(inp.Name).
		SetDescription(inp.Description).
		SetUnitPrice(inp.Price).
		SetFeatures(inp.Features).
		SetNillableSpecialNote(inp.SpecialNote).
		SetImageID(inp.ImageID).
		SetCreatorID(uID).
		Exec(ctx)

}

func SavePackage(ctx context.Context, uID string, inp *model.ProductPackageInput) error {
	db := ent.GetClient()
	defer db.CloseClient()

	exist, err := db.ProductPackage.Query().Where(productpackage.ID(inp.ID)).Exist(ctx)
	if err != nil {
		log.Error(err)
		return err
	}

	if exist {
		// remove old item entries
		err = db.ProductPackage.UpdateOneID(inp.ID).ClearItems().Exec(ctx)
		if err != nil {
			return err
		}

		return db.ProductPackage.UpdateOneID(inp.ID).
			SetName(inp.Name).
			SetType(inp.Category).
			SetSoldAs(inp.SoldAs).
			SetDescription(inp.Description).
			SetPrice(inp.Price).
			SetFeatures(inp.Features).
			AddItemIDs(inp.ProductIDs...).
			Exec(ctx)
	}

	return db.ProductPackage.Create().
		SetID(inp.ID).
		SetType(inp.Category).
		SetSoldAs(inp.SoldAs).
		SetName(inp.Name).
		SetDescription(inp.Description).
		SetPrice(inp.Price).
		SetFeatures(inp.Features).
		AddItemIDs(inp.ProductIDs...).
		SetCreatorID(uID).
		Exec(ctx)

}
