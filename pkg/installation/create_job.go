package installation

import (
	"context"
	"errors"
	"roofix/ent"
	"roofix/ent/productpackage"
	"roofix/pkg/enum"
	"roofix/pkg/util/log"
	"roofix/server/model"
	"strings"
)

func CreateJob(
	ctx context.Context, installType enum.InstallationType, pkgID string, productID *string, owner *model.InstallationOwnerInput,
	requestingPartnerID *string, repUserID, creatorID string,
) (string, error) {
	if strings.TrimSpace(pkgID) == "" {
		return "", errors.New("packageID is missing")
	}

	if owner == nil {
		return "", errors.New("owner information is missing")
	}

	db := ent.GetClient()
	defer db.CloseClient()

	// pull package info
	pkg, err := db.ProductPackage.Query().
		WithItems(func(i *ent.ProductQuery) {
			i.WithImage()
		}).
		Where(productpackage.ID(pkgID)).
		Only(ctx)
	if err != nil {
		return "", err
	}

	// some package is marker as sold as individual item, packageID will be passed on
	// and productID must be there in package items
	var product *ent.Product
	if productID != nil {
		for _, p := range pkg.Edges.Items {
			if p.ID == *productID {
				product = p
				break
			}
		}

		// product is not supposed to be nil by now
		if product == nil {
			return "", errors.New("failed to find product in package")
		}
	}

	if pkg.SoldAs == enum.SoldAsIndividualItem && product == nil {
		return "", errors.New("please select one of product from the package")
	}

	// begin TX
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return "", err
	}

	create := tx.InstallationJob.
		Create().
		SetType(installType).
		SetOwnerName(owner.Name).
		SetOwnerEmail(owner.Email).
		SetOwnerPhone(owner.Phone).
		SetOwnerAddress(owner.Address).
		SetNillableOwnerAddressLat(owner.Latitude).
		SetNillableOwnerAddressLng(owner.Longitude).
		SetNillableSpecialNote(owner.SpecialNote).
		SetNillableRequestingPartnerID(requestingPartnerID).
		SetSalesRepID(repUserID).
		SetCreatorID(creatorID).
		SetPkg(pkg.Name)

	// price & description
	var price float64
	var items []*ent.Product
	switch pkg.SoldAs {
	case enum.SoldAsPackage:
		price = pkg.Price
		items = pkg.Edges.Items
		if pkg.Type == enum.ProductHVAC {
			create.
				SetPkgDescription(pkg.Description).
				SetPkgFeatures(pkg.Features)
		}
	case enum.SoldAsIndividualItem:
		price = product.UnitPrice
		items = append(items, product)
	}

	// save installation-job info to DB
	j, err := create.SetPrice(price).Save(ctx)
	if err != nil {
		return "", err
	}

	// save job-items to DB
	bulk := make([]*ent.InstallationJobItemCreate, len(items))
	for idx, item := range pkg.Edges.Items {
		img := item.Edges.Image
		bulk[idx] = tx.InstallationJobItem.
			Create().
			SetJobID(j.ID).
			SetName(item.Name).
			SetDescription(item.Description).
			SetFeatures(item.Features).
			SetSpecialNote(item.SpecialNote).
			SetPrice(item.UnitPrice).
			SetImgKey(img.Key)
	}
	if err = tx.InstallationJobItem.CreateBulk(bulk...).Exec(ctx); err != nil {
		if txErr := tx.Rollback(); txErr != nil {
			log.Error(txErr)
		}

		return "", err
	}

	// commit TX
	if err = tx.Commit(); err != nil {
		return "", err
	}

	return j.ID, nil
}
