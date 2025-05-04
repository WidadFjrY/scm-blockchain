package controller

import "github.com/gin-gonic/gin"

type ProductController interface {
	ProductCreate(ctx *gin.Context)
	BrandCreate(ctx *gin.Context)
	UnitCreate(ctx *gin.Context)

	GetProducts(ctx *gin.Context)
	GetProduct(ctx *gin.Context)
	GetTotalProduct(ctx *gin.Context)

	DeleteProductById(ctx *gin.Context)

	GetBrands(ctx *gin.Context)
	GetUnits(ctx *gin.Context)

	AddToCart(ctx *gin.Context)
	GetCarts(ctx *gin.Context)
	UpdateCartQty(ctx *gin.Context)
	DeleteItemCart(ctx *gin.Context)

	UpdateProductStock(ctx *gin.Context)
}
