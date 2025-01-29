package repository

import (
	"context"
	"scm-blockchain-ethereum/domain/model"
	"scm-blockchain-ethereum/pkg/helper"

	"gorm.io/gorm"
)

type DistributorRepositoryImpl struct{}

func NewDistributorRepository() DistributorRepository {
	return &DistributorRepositoryImpl{}
}

func (repo *DistributorRepositoryImpl) Create(ctx context.Context, tx *gorm.DB, distributor model.Distributor) model.Distributor {
	result := tx.WithContext(ctx).Create(&distributor)
	helper.Err(result.Error)

	return distributor
}

func (repo *DistributorRepositoryImpl) GetAll(ctx context.Context, tx *gorm.DB) []model.Distributor {
	var distributors []model.Distributor

	err := tx.WithContext(ctx).Find(&distributors).Error
	helper.Err(err)

	return distributors
}

func (repo *DistributorRepositoryImpl) GetByID(ctx context.Context, tx *gorm.DB, distributorId string) model.Distributor {
	var distributor model.Distributor

	err := tx.WithContext(ctx).Where("id = ?", distributorId).First(&distributor).Error
	helper.Err(err)

	return distributor
}
