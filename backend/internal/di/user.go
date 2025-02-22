package di

import (
	"scm-blockchain-ethereum/internal/controller"
	"scm-blockchain-ethereum/internal/repository"
	"scm-blockchain-ethereum/internal/service"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

func UserDI(db *gorm.DB, validator *validator.Validate) controller.UserController {
	repo := repository.NewUserRepository()
	serv := service.NewUserService(db, validator, repo)
	cntrl := controller.NewUserController(serv)

	return cntrl
}
