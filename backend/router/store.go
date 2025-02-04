package router

import (
	"scm-blockchain-ethereum/internal/controller"
	"scm-blockchain-ethereum/internal/middleware"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func StoreRouter(router *gin.Engine, db *gorm.DB, contrl controller.StoreController) {
	auth := router.Group("/")
	auth.Use(middleware.Auth(db))
	{
		auth.POST("/api/store", contrl.Create)
		auth.GET("/api/store", contrl.GetAll)
		auth.GET("/api/store/:store_id", contrl.GetById)
	}
}
