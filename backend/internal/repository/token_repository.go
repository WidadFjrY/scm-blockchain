package repository

import (
	"context"
	"scm-blockchain-ethereum/domain/model"

	"gorm.io/gorm"
)

type TokenRepository interface {
	Create(ctx context.Context, tx *gorm.DB, token model.TokenAuth)
	IsValid(ctx context.Context, tx *gorm.DB, token string) bool
	Update(ctx context.Context, tx *gorm.DB, token string)
}
