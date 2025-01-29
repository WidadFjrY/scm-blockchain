package repository

import (
	"context"
	"scm-blockchain-ethereum/domain/model"

	"gorm.io/gorm"
)

type DistributorRepository interface {
	Create(ctx context.Context, tx *gorm.DB, distributor model.Distributor) model.Distributor
	GetAll(ctx context.Context, tx *gorm.DB) []model.Distributor
	GetByID(ctx context.Context, tx *gorm.DB, distributorId string) model.Distributor
}
