package di

import (
	"scm-blockchain-ethereum/internal/controller"
	"scm-blockchain-ethereum/internal/repository"
	"scm-blockchain-ethereum/internal/service"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

func ProductDI(db *gorm.DB, validator *validator.Validate) controller.ProductController {
	repo := repository.NewProductRepository()
	serv := service.NewProductService(db, validator, repo)
	cntrl := controller.NewProductController(serv)

	return cntrl
}
