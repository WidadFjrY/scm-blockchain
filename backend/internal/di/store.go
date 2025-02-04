package di

import (
	"scm-blockchain-ethereum/internal/controller"
	"scm-blockchain-ethereum/internal/repository"
	"scm-blockchain-ethereum/internal/service"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

func StoreDI(db *gorm.DB, validator *validator.Validate) controller.StoreController {
	repo := repository.NewStoreRepository()
	serv := service.NewStoreService(db, validator, repo)
	cntrl := controller.NewStoreController(serv)

	return cntrl
}
