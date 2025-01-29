package router

import (
	"scm-blockchain-ethereum/internal/controller"
	"scm-blockchain-ethereum/internal/middleware"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func DistributorRouter(router *gin.Engine, db *gorm.DB, contrl controller.DistributorController) {
	auth := router.Group("/")
	auth.Use(middleware.Auth(db))
	{
		auth.POST("/api/distributor", contrl.Create)
		auth.GET("/api/distributor", contrl.GetAll)
		auth.GET("/api/distributor/:distId", contrl.GetById)
	}
}
