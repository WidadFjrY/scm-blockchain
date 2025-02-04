package repository

import (
	"context"
	"scm-blockchain-ethereum/domain/model"
	"scm-blockchain-ethereum/pkg/helper"

	"gorm.io/gorm"
)

type StoreRepositoryImpl struct{}

func NewStoreRepository() StoreRepository {
	return &StoreRepositoryImpl{}
}

func (repo *StoreRepositoryImpl) Create(ctx context.Context, tx *gorm.DB, store model.Store) {
	err := tx.WithContext(ctx).Create(&store).Error
	helper.Err(err)
}

func (repo *StoreRepositoryImpl) GetAll(ctx context.Context, tx *gorm.DB) []model.Store {
	var stores []model.Store

	err := tx.WithContext(ctx).Find(&stores).Error
	helper.Err(err)

	return stores
}

func (repo *StoreRepositoryImpl) GetById(ctx context.Context, tx *gorm.DB, storeId string) model.Store {
	var store model.Store

	err := tx.WithContext(ctx).Where("id = ?", storeId).Preload("User").First(&store).Error
	helper.Err(err)

	return store
}
