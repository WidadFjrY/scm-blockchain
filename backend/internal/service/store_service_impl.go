package service

import (
	"context"
	"math/rand"
	"scm-blockchain-ethereum/domain/model"
	"scm-blockchain-ethereum/domain/web"
	"scm-blockchain-ethereum/internal/repository"
	"scm-blockchain-ethereum/pkg/exception"
	"scm-blockchain-ethereum/pkg/helper"
	"time"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type StoreServiceImpl struct {
	DB        *gorm.DB
	Validate  *validator.Validate
	StoreRepo repository.StoreRepository
}

func NewStoreService(db *gorm.DB, validate *validator.Validate, storeRepo repository.StoreRepository) StoreService {
	return &StoreServiceImpl{DB: db, Validate: validate, StoreRepo: storeRepo}
}

func (serv *StoreServiceImpl) Create(ctx context.Context, role string, request web.StoreCreateRequest) {
	valErr := serv.Validate.Struct(&request)
	helper.ValError(valErr)

	if role != "Admin" {
		panic(exception.NewBadRequestError("have no authority"))
	}

	rand.NewSource(time.Now().Unix())
	err := serv.DB.Transaction(func(tx *gorm.DB) error {
		serv.StoreRepo.Create(ctx, tx, model.Store{
			ID:      helper.GenerateRandomString(15),
			UserID:  request.UserID,
			Name:    request.Name,
			Address: request.Address,
			Telp:    request.Telp,
		})
		return nil
	})
	helper.Err(err)
}

func (serv *StoreServiceImpl) GetAll(ctx context.Context) []web.StoreResponse {
	var stores []web.StoreResponse

	err := serv.DB.Transaction(func(tx *gorm.DB) error {
		for _, store := range serv.StoreRepo.GetAll(ctx, tx) {
			stores = append(stores, web.StoreResponse{
				ID:          store.ID,
				Name:        store.Name,
				ManagerName: store.User.Name,
				Address:     store.Address,
				Telp:        store.Telp,
			})
		}
		return nil
	})
	helper.Err(err)

	return stores
}

func (serv *StoreServiceImpl) GetById(ctx context.Context, stroreId string) web.StoreResponse {
	var store model.Store

	err := serv.DB.Transaction(func(tx *gorm.DB) error {
		store = serv.StoreRepo.GetById(ctx, tx, stroreId)
		return nil
	})
	helper.Err(err)

	return web.StoreResponse{
		ID:          store.ID,
		Name:        store.Name,
		ManagerName: store.User.Name,
		Address:     store.Address,
		Telp:        store.Telp,
		CreatedAt:   store.CreatedAt,
		UpdatedAt:   store.UpdatedAt,
	}
}
