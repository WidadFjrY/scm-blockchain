package repository

import (
	"context"
	"scm-blockchain-ethereum/domain/model"

	"gorm.io/gorm"
)

type ProductRepository interface {
	Create(ctx context.Context, tx *gorm.DB, product model.Product) model.Product
	GetAll(ctx context.Context, tx *gorm.DB) []model.Product
	GetById(ctx context.Context, tx *gorm.DB, productId string) model.Product
}
