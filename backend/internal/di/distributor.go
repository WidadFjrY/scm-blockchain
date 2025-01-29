package di

import (
	"scm-blockchain-ethereum/internal/controller"
	"scm-blockchain-ethereum/internal/repository"
	"scm-blockchain-ethereum/internal/service"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

func DistributorDI(db *gorm.DB, validator *validator.Validate) controller.DistributorController {
	repo := repository.NewDistributorRepository()
	serv := service.NewDistributorService(db, validator, repo)
	cntrl := controller.NewDistributorController(serv)

	return cntrl
}
