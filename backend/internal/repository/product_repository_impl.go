package repository

import (
	"context"
	"errors"
	"scm-blockchain-ethereum/domain/model"
	"scm-blockchain-ethereum/pkg/exception"
	"scm-blockchain-ethereum/pkg/helper"
	"strings"
	"time"

	"gorm.io/gorm"
)

type ProductRepositoryImpl struct{}

func NewProductRepository() ProductRepository {
	return &ProductRepositoryImpl{}
}

func (repo *ProductRepositoryImpl) ProductCreate(ctx context.Context, tx *gorm.DB, product model.NormalizedProduct) {
	err := tx.WithContext(ctx).Model(&model.Product{}).Create(&product.Product).Error
	helper.Err(err)

	err = tx.WithContext(ctx).Model(&model.ProductStock{}).Create(&product.ProductStock).Error
	helper.Err(err)

	err = tx.WithContext(ctx).Model(&model.ProductPrice{}).Create(&product.ProductPrice).Error
	helper.Err(err)
}

func (repo *ProductRepositoryImpl) BrandCreate(ctx context.Context, tx *gorm.DB, Brand model.ProductBrand) {
	err := tx.WithContext(ctx).Create(&Brand).Error
	if strings.Contains(err.Error(), "Duplicate entry") {
		panic(exception.NewBadRequestError("Brand already exsits!"))
	}
}

func (repo *ProductRepositoryImpl) UnitCreate(ctx context.Context, tx *gorm.DB, Unit model.ProductUnit) {
	err := tx.WithContext(ctx).Create(&Unit).Error
	if strings.Contains(err.Error(), "Duplicate entry") {
		panic(exception.NewBadRequestError("Unit already exsits!"))
	}
}

func (repo *ProductRepositoryImpl) GetProducts(ctx context.Context, tx *gorm.DB) []model.Product {
	var products []model.Product

	helper.Err(tx.WithContext(ctx).
		Preload("Brand").
		Preload("Unit").
		Preload("Prices", func(db *gorm.DB) *gorm.DB {
			return db.Order("updated_at DESC")
		}).
		Preload("Stocks", func(db *gorm.DB) *gorm.DB {
			return db.Order("updated_at DESC")
		}).
		Find(&products).Error)

	return products
}

func (repo *ProductRepositoryImpl) GetProduct(ctx context.Context, tx *gorm.DB, productId string) model.Product {
	var product model.Product

	err := tx.WithContext(ctx).Where("id = ?", productId).Preload("Brand").Preload("Unit").Preload("Prices", func(db *gorm.DB) *gorm.DB {
		return db.Order("updated_at DESC").Limit(1)
	}).Preload("Stocks", func(db *gorm.DB) *gorm.DB {
		return db.Order("updated_at DESC").Limit(1)
	}).First(&product).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		panic(exception.NewNotFoundError("product not found"))
	}

	return product
}

func (repo *ProductRepositoryImpl) GetBrands(ctx context.Context, tx *gorm.DB) []model.ProductBrand {
	var brands []model.ProductBrand

	helper.Err(tx.WithContext(ctx).Find(&brands).Error)
	return brands
}

func (repo *ProductRepositoryImpl) GetUnits(ctx context.Context, tx *gorm.DB) []model.ProductUnit {
	var units []model.ProductUnit

	helper.Err(tx.WithContext(ctx).Find(&units).Error)
	return units
}

func (repo *ProductRepositoryImpl) AddToCart(ctx context.Context, tx *gorm.DB, cart model.Cart) {
	helper.Err(tx.WithContext(ctx).Create(&cart).Error)
}

func (repo *ProductRepositoryImpl) DeleteItemCart(ctx context.Context, tx *gorm.DB, cartId string) {
	helper.Err(tx.WithContext(ctx).Table("carts").Delete(nil).Where("id = ?", cartId).Error)
}

func (repo *ProductRepositoryImpl) GetCarts(ctx context.Context, tx *gorm.DB, userId string) []model.Cart {
	var carts []model.Cart
	helper.Err(tx.WithContext(ctx).Where("user_id = ?", userId).Preload("Product").Preload("Product.Prices", func(db *gorm.DB) *gorm.DB {
		return db.Order("updated_at DESC")
	}).Preload("Product.Brand").Preload("Product.Unit").Find(&carts).Error)
	return carts
}

func (repo *ProductRepositoryImpl) UpdateCartQty(ctx context.Context, tx *gorm.DB, cart model.Cart) {
	helper.Err(tx.WithContext(ctx).Where("product_id = ?", cart.ProductID).Table("carts").Updates(map[string]interface{}{
		"quantity":   cart.Quantity,
		"updated_at": time.Now(),
	}).Error)
}
