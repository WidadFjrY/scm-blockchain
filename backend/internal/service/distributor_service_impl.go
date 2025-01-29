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

type DistributorServiceImpl struct {
	DB              *gorm.DB
	Validator       *validator.Validate
	DistributorRepo repository.DistributorRepository
}

func NewDistributorService(db *gorm.DB, validator *validator.Validate, distributorRepo repository.DistributorRepository) DistributorService {
	return &DistributorServiceImpl{DB: db, Validator: validator, DistributorRepo: distributorRepo}
}

func (serv *DistributorServiceImpl) Create(ctx context.Context, request web.DistributorCreateRequest, userId string) web.DistributorCreateResponse {
	valErr := serv.Validator.Struct(&request)
	helper.ValError(valErr)

	userRepo := repository.NewUserRepository()
	userServ := NewUserService(serv.DB, serv.Validator, userRepo)

	var distributor model.Distributor

	rand.NewSource(time.Now().UnixNano())
	txErr := serv.DB.Transaction(func(tx *gorm.DB) error {
		user := userServ.GetUserById(ctx, userId)
		if user.Role != "Admin" {
			panic(exception.NewBadRequestError("only admin can add distributor"))
		}

		distributor = serv.DistributorRepo.Create(ctx, tx, model.Distributor{
			ID:      helper.GenerateRandomString(15),
			Name:    request.Name,
			Address: request.Address,
		})
		return nil
	})
	helper.Err(txErr)

	return web.DistributorCreateResponse{
		Name:      distributor.Name,
		Address:   distributor.Address,
		CreatedAt: time.Now(),
	}
}

func (serv *DistributorServiceImpl) GetAll(ctx context.Context) []web.DistributorResponse {
	var distributors []web.DistributorResponse

	errTx := serv.DB.Transaction(func(tx *gorm.DB) error {
		for _, distributor := range serv.DistributorRepo.GetAll(ctx, tx) {
			distributorResp := web.DistributorResponse{
				ID:        distributor.ID,
				Name:      distributor.Name,
				Address:   distributor.Address,
				CreatedAt: distributor.CreatedAt,
				UpdatedAt: distributor.UpdatedAt,
			}
			distributors = append(distributors, distributorResp)
		}
		return nil
	})

	helper.Err(errTx)

	return distributors
}

func (serv *DistributorServiceImpl) GetByID(ctx context.Context, distributorId string) web.DistributorResponse {
	var distributor model.Distributor

	errTx := serv.DB.Transaction(func(tx *gorm.DB) error {
		distributor = serv.DistributorRepo.GetByID(ctx, tx, distributorId)
		return nil
	})

	helper.Err(errTx)

	return web.DistributorResponse{
		ID:        distributor.ID,
		Name:      distributor.Name,
		Address:   distributor.Address,
		CreatedAt: distributor.CreatedAt,
		UpdatedAt: distributor.UpdatedAt,
	}
}
