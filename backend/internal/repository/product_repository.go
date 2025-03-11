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

	GetBrands(ctx context.Context, tx *gorm.DB) []model.ProductBrand
	GetUnits(ctx context.Context, tx *gorm.DB) []model.ProductUnit

	AddToCart(ctx context.Context, tx *gorm.DB, cart model.Cart)
	GetCarts(ctx context.Context, tx *gorm.DB, userId string) []model.Cart
	UpdateCartQty(ctx context.Context, tx *gorm.DB, cart model.Cart)
	DeleteItemCart(ctx context.Context, tx *gorm.DB, cartId string)
}
