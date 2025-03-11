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
		auth.POST("/api/product/add", contrl.ProductCreate)
		auth.GET("/api/product/:productId", contrl.GetProduct)
		auth.GET("/api/products", contrl.GetProducts)

		auth.POST("/api/brand/add", contrl.BrandCreate)
		auth.GET("/api/brands", contrl.GetBrands)

		auth.POST("/api/unit/add", contrl.UnitCreate)
		auth.GET("/api/units", contrl.GetUnits)

		auth.POST("/api/addToCart", contrl.AddToCart)
		auth.PUT("/api/cart", contrl.UpdateCartQty)
		auth.GET("/api/cart", contrl.GetCarts)
	}
}
