package router

import (
	"scm-blockchain-ethereum/internal/controller"
	"scm-blockchain-ethereum/internal/middleware"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ProductRouter(router *gin.Engine, db *gorm.DB, contrl controller.ProductController) {
	auth := router.Group("/")
	auth.Use(middleware.Auth(db))
	{
		auth.POST("/api/product", contrl.Create)
		auth.GET("/api/product", contrl.GetAll)
		auth.GET("/api/product/:product_id", contrl.GetById)
	}
}
