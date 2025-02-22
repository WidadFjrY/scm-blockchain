package server

import (
	"scm-blockchain-ethereum/internal/app"
	"scm-blockchain-ethereum/internal/di"
	"scm-blockchain-ethereum/internal/middleware"
	"scm-blockchain-ethereum/pkg/helper"
	"scm-blockchain-ethereum/router"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func Run() {
	logFilePaths := []string{"pkg/log/gorm.log", "pkg/log/gin.log"}
	helper.RefreshLog(logFilePaths)

	db := app.NewDB()
	validator := validator.New()

	gin := gin.Default()
	gin.Use(helper.NewFileGinLog(), middleware.CORSMiddleware(), middleware.ErrorHandling())

	userCntrl := di.UserDI(db, validator)
	ProductCntrl := di.ProductDI(db, validator)

	router.UserRouter(gin, db, userCntrl)
	router.ProductRouter(gin, db, ProductCntrl)

	gin.Static("/api/public", "./public")

	err := gin.Run("localhost:8080")
	helper.Err(err)
}
