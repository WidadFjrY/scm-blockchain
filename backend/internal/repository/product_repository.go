package repository

import (
	"context"
	"scm-blockchain-ethereum/domain/model"

	"gorm.io/gorm"
)

type ProductRepository interface {
	ProductCreate(ctx context.Context, tx *gorm.DB, product model.NormalizedProduct)
	BrandCreate(ctx context.Context, tx *gorm.DB, Brand model.ProductBrand)
	UnitCreate(ctx context.Context, tx *gorm.DB, Unit model.ProductUnit)

	GetProducts(ctx context.Context, tx *gorm.DB) []model.Product
	GetProduct(ctx context.Context, tx *gorm.DB, productId string) model.Product
}
