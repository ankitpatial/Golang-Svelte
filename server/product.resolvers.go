package server

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.32

import (
	"context"
	"roofix/ent"
	"roofix/pkg/account"
	"roofix/pkg/document"
	"roofix/pkg/enum"
	"roofix/pkg/product"
	"roofix/server/generated"
	"roofix/server/model"
)

// SaveProductPackage is the resolver for the saveProductPackage field.
func (r *mutationResolver) SaveProductPackage(ctx context.Context, input model.ProductPackageInput) (bool, error) {
	uID := account.CtxUserID(ctx)
	err := product.SavePackage(ctx, uID, &input)
	return err == nil, err
}

// SaveProduct is the resolver for the saveProduct field.
func (r *mutationResolver) SaveProduct(ctx context.Context, input model.ProductInput) (bool, error) {
	uID := account.CtxUserID(ctx)
	err := product.Save(ctx, uID, &input)
	return err == nil, err
}

// Category is the resolver for the category field.
func (r *packageResolver) Category(ctx context.Context, obj *ent.ProductPackage) (enum.Product, error) {
	if obj == nil {
		return "", nil
	}

	return obj.Type, nil
}

// Image is the resolver for the image field.
func (r *productDetailResolver) Image(ctx context.Context, obj *ent.Product) (*document.Info, error) {
	if obj == nil || obj.Edges.Image == nil {
		return nil, nil
	}

	img := obj.Edges.Image
	return &document.Info{
		ID:      img.ID,
		Key:     img.Key,
		Folder:  img.Folder,
		Section: img.Section,
	}, nil
}

// Category is the resolver for the category field.
func (r *productInfoResolver) Category(ctx context.Context, obj *ent.Product) (enum.Product, error) {
	if obj == nil {
		return "", nil
	}

	return obj.Type, nil
}

// Price is the resolver for the price field.
func (r *productInfoResolver) Price(ctx context.Context, obj *ent.Product) (float64, error) {
	if obj == nil {
		return 0, nil
	}

	return obj.UnitPrice, nil
}

// Image is the resolver for the image field.
func (r *productInfoResolver) Image(ctx context.Context, obj *ent.Product) (*document.Info, error) {
	if obj == nil || obj.Edges.Image == nil {
		return nil, nil
	}

	img := obj.Edges.Image
	return &document.Info{
		ID:      img.ID,
		Folder:  img.Folder,
		Section: img.Section,
		Key:     img.Key,
	}, nil
}

// Category is the resolver for the category field.
func (r *productPackageResolver) Category(ctx context.Context, obj *ent.ProductPackage) (enum.Product, error) {
	if obj == nil {
		return "", nil
	}

	return obj.Type, nil
}

// ProductPackages is the resolver for the productPackages field.
func (r *queryResolver) ProductPackages(ctx context.Context, category *enum.Product, search *string, page model.PageInput) (*ent.ProductPackageConnection, error) {
	return product.Packages(ctx, category, search, page)
}

// Products is the resolver for the products field.
func (r *queryResolver) Products(ctx context.Context, category *enum.Product, search *string, page model.PageInput) (*ent.ProductConnection, error) {
	return product.List(ctx, category, search, page)
}

// SmartHomePackages is the resolver for the smartHomePackages field.
func (r *queryResolver) SmartHomePackages(ctx context.Context, page model.PageInput) (*ent.ProductPackageConnection, error) {
	cat := enum.ProductSmartHome
	return product.Packages(ctx, &cat, nil, page)
}

// HvacPackages is the resolver for the hvacPackages field.
func (r *queryResolver) HvacPackages(ctx context.Context, page model.PageInput) (*ent.ProductPackageConnection, error) {
	cat := enum.ProductHVAC
	return product.Packages(ctx, &cat, nil, page)
}

// Package returns generated.PackageResolver implementation.
func (r *Resolver) Package() generated.PackageResolver { return &packageResolver{r} }

// ProductDetail returns generated.ProductDetailResolver implementation.
func (r *Resolver) ProductDetail() generated.ProductDetailResolver { return &productDetailResolver{r} }

// ProductInfo returns generated.ProductInfoResolver implementation.
func (r *Resolver) ProductInfo() generated.ProductInfoResolver { return &productInfoResolver{r} }

// ProductPackage returns generated.ProductPackageResolver implementation.
func (r *Resolver) ProductPackage() generated.ProductPackageResolver {
	return &productPackageResolver{r}
}

type packageResolver struct{ *Resolver }
type productDetailResolver struct{ *Resolver }
type productInfoResolver struct{ *Resolver }
type productPackageResolver struct{ *Resolver }
