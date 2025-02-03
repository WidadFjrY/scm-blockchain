package repository

import (
	"context"
	"errors"
	"fmt"
	"scm-blockchain-ethereum/domain/model"
	"scm-blockchain-ethereum/pkg/exception"
	"scm-blockchain-ethereum/pkg/helper"
	"time"

	"gorm.io/gorm"
)

type ProductRepositoryImpl struct{}

func NewProductRepository() ProductRepository {
	return &ProductRepositoryImpl{}
}

func (repo *ProductRepositoryImpl) Create(ctx context.Context, tx *gorm.DB, product model.Product) model.Product {
	err := tx.WithContext(ctx).Create(&product).Error
	helper.Err(err)

	product.CreatedAt = time.Now()
	return product
}

func (repo *ProductRepositoryImpl) GetAll(ctx context.Context, tx *gorm.DB) []model.Product {
	var products []model.Product

	err := tx.WithContext(ctx).Find(&products).Error
	helper.Err(err)

	return products
}

func (repo *ProductRepositoryImpl) GetById(ctx context.Context, tx *gorm.DB, productId string) model.Product {
	var product model.Product

	err := tx.WithContext(ctx).Where("id = ?", productId).Preload("Distributor").First(&product).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		panic(exception.NewNotFoundError(fmt.Sprintf("product with id %s not found", productId)))
	}
	helper.Err(err)

	return product
}
