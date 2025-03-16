package repository

import (
	"context"
	"errors"
	"fmt"
	"scm-blockchain-ethereum/domain/model"
	"scm-blockchain-ethereum/pkg/exception"
	"scm-blockchain-ethereum/pkg/helper"
	"strings"
	"time"

	"gorm.io/gorm"
)

type UserRepositoryImpl struct{}

func NewUserRepository() UserRepository {
	return &UserRepositoryImpl{}
}

func (repo *UserRepositoryImpl) Create(ctx context.Context, tx *gorm.DB, user model.User) model.User {
	result := tx.WithContext(ctx).Create(&user)
	if result.Error != nil {
		if strings.Contains(result.Error.Error(), "Error 1062") {
			panic(exception.NewBadRequestError("email already in use"))
		} else {
			panic(result.Error.Error())
		}
	}

	user.CreatedAt = time.Now()
	return user
}

func (Repo *UserRepositoryImpl) GetUserByEmail(ctx context.Context, tx *gorm.DB, email string) model.User {
	var user model.User

	tx.WithContext(ctx).Where("email = ?", email).First(&user)

	return user
}

func (Repo *UserRepositoryImpl) IsUserExsit(ctx context.Context, tx *gorm.DB, userId string) bool {
	var user model.User

	result := tx.WithContext(ctx).Where("id = ?", userId).First(&user)
	return errors.Is(result.Error, gorm.ErrRecordNotFound) // true if user not found
}

func (Repo *UserRepositoryImpl) GetUserById(ctx context.Context, tx *gorm.DB, userId string) model.User {
	var user model.User

	result := tx.WithContext(ctx).Where("id = ?", userId).First(&user)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		panic(exception.NewNotFoundError(fmt.Sprintf("user with id %s not found", userId)))
	}

	helper.Err(result.Error)
	return user
}

func (repo *UserRepositoryImpl) GetAll(ctx context.Context, tx *gorm.DB) []model.User {
	var user []model.User

	err := tx.WithContext(ctx).Find(&user).Error
	helper.Err(err)

	return user
}

func (repo *UserRepositoryImpl) GetUserByManager(ctx context.Context, tx *gorm.DB) []model.User {
	var user []model.User

	err := tx.WithContext(ctx).Where("role = 'Store Manager'").Find(&user).Error
	helper.Err(err)

	return user
}

func (repo *UserRepositoryImpl) UpdateById(ctx context.Context, tx *gorm.DB, user model.User) {
	err := tx.WithContext(ctx).Table("users").Where("id = ?", user.ID).Updates(map[string]interface{}{
		"name":       user.Name,
		"email":      user.Email,
		"telp":       user.Telp,
		"updated_at": time.Now(),
	}).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		panic(exception.NewNotFoundError(fmt.Sprintf("user with id %s not found", user.ID)))
	}

	helper.Err(err)
}

func (repo *UserRepositoryImpl) UpdatePassword(ctx context.Context, tx *gorm.DB, newPassword string, userId string) {
	err := tx.WithContext(ctx).Table("users").Where("id = ?", userId).Updates(map[string]interface{}{
		"password":   newPassword,
		"updated_at": time.Now(),
	}).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		panic(exception.NewNotFoundError(fmt.Sprintf("user with id %s not found", userId)))
	}

	helper.Err(err)
}

func (repo *UserRepositoryImpl) GetUserByETHAddr(ctx context.Context, tx *gorm.DB, addr string) model.User {
	var user model.User

	result := tx.WithContext(ctx).Where("eth_address = ?", addr).First(&user)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		panic(exception.NewNotFoundError(fmt.Sprintf("user with id %s not found", addr)))
	}

	helper.Err(result.Error)
	return user
}
