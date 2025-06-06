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
	if err != nil {
		if strings.Contains(err.Error(), "Duplicate entry") {
			panic(exception.NewBadRequestError("Brand already exists!"))
		}
		panic(err)
	}
}

func (repo *ProductRepositoryImpl) UnitCreate(ctx context.Context, tx *gorm.DB, unit model.ProductUnit) {
	err := tx.WithContext(ctx).Create(&unit).Error

	if err != nil {
		if strings.Contains(err.Error(), "Duplicate entry") {
			panic(exception.NewBadRequestError("Unit already exists!"))
		}
		panic(err)
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

func (repo *ProductRepositoryImpl) DeleteProductById(ctx context.Context, tx *gorm.DB, productId string) {
	tables := []string{"product_prices", "product_stocks", "carts"}
	for _, table := range tables {
		helper.Err(tx.WithContext(ctx).Table(table).Where("product_id = ?", productId).Delete(nil).Error)
	}
	helper.Err(tx.WithContext(ctx).Table("products").Where("id = ?", productId).Delete(nil).Error)
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

func (repo *ProductRepositoryImpl) DeleteItemCart(ctx context.Context, tx *gorm.DB, productId string) {
	helper.Err(tx.WithContext(ctx).Where("product_id = ?", productId).Table("carts").Delete(nil).Error)
}

func (repo *ProductRepositoryImpl) GetCarts(ctx context.Context, tx *gorm.DB, userId string) []model.Cart {
	var carts []model.Cart
	helper.Err(tx.WithContext(ctx).Where("user_id = ?", userId).Preload("Product").Preload("Product.Prices", func(db *gorm.DB) *gorm.DB {
		return db.Order("updated_at DESC")
	}).Preload("Product.Stocks", func(db *gorm.DB) *gorm.DB {
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

func (repo *ProductRepositoryImpl) GetTotalProducts(ctx context.Context, tx *gorm.DB) int64 {
	var totalProducts int64

	helper.Err(tx.WithContext(ctx).Table("products").Count(&totalProducts).Error)
	return totalProducts
}

func (repo *ProductRepositoryImpl) UpdateProductStock(ctx context.Context, tx *gorm.DB, productStock model.ProductStock) {
	var latestStock model.ProductStock
	helper.Err(tx.WithContext(ctx).Where("product_id = ?", productStock.ProductID).Order("updated_at DESC").First(&latestStock).Error)

	if productStock.StockIn > 0 {
		newProductStock := model.ProductStock{
			ID:         productStock.ID,
			ProductID:  productStock.ProductID,
			StockIn:    productStock.StockIn,
			StockTotal: latestStock.StockTotal + productStock.StockIn,
			UpdatedAt:  time.Now(),
		}
		helper.Err(tx.WithContext(ctx).Create(&newProductStock).Error)
	} else if productStock.StockOut > 0 {
		if productStock.StockOut > latestStock.StockTotal {
			panic(exception.NewBadRequestError("stock out exceeds total stock"))
		}
		newProductStock := model.ProductStock{
			ID:         productStock.ID,
			ProductID:  productStock.ProductID,
			StockOut:   productStock.StockOut,
			StockTotal: latestStock.StockTotal - productStock.StockOut,
			UpdatedAt:  time.Now(),
		}
		helper.Err(tx.WithContext(ctx).Create(&newProductStock).Error)
	}
}
