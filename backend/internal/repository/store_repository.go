package repository

import (
	"context"
	"scm-blockchain-ethereum/domain/model"

	"gorm.io/gorm"
)

type StoreRepository interface {
	Create(ctx context.Context, tx *gorm.DB, store model.Store)
	GetAll(ctx context.Context, tx *gorm.DB) []model.Store
	GetById(ctx context.Context, tx *gorm.DB, storeId string) model.Store
}
