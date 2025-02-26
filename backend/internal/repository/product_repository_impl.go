package repository

import (
	"context"
	"scm-blockchain-ethereum/domain/model"
	"scm-blockchain-ethereum/pkg/exception"
	"scm-blockchain-ethereum/pkg/helper"
	"strings"

	"gorm.io/gorm"
)

type ProductRepositoryImpl struct{}

func NewProductRepository() ProductRepository {
	return &ProductRepositoryImpl{}
}

func (repo *ProductRepositoryImpl) ProductCreate(ctx context.Context, tx *gorm.DB, product model.NormalizedProduct) {
	helper.Err(tx.WithContext(ctx).Create([]interface{}{
		product.Product,
		product.ProductPrice,
		product.ProductStock,
	}).Error)
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
