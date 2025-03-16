package repository

import (
	"context"
	"scm-blockchain-ethereum/domain/model"

	"gorm.io/gorm"
)

type UserRepository interface {
	Create(ctx context.Context, tx *gorm.DB, user model.User) model.User
	GetUserByEmail(ctx context.Context, tx *gorm.DB, email string) model.User
	IsUserExsit(ctx context.Context, tx *gorm.DB, userId string) bool
	GetUserById(ctx context.Context, tx *gorm.DB, userId string) model.User
	GetUserByManager(ctx context.Context, tx *gorm.DB) []model.User
	GetAll(ctx context.Context, tx *gorm.DB) []model.User
	UpdateById(ctx context.Context, tx *gorm.DB, user model.User)
	UpdatePassword(ctx context.Context, tx *gorm.DB, newPassword string, userId string)
	GetUserByETHAddr(ctx context.Context, tx *gorm.DB, addr string) model.User
}
