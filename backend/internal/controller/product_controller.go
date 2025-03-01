package controller

import "github.com/gin-gonic/gin"

type ProductController interface {
	ProductCreate(ctx *gin.Context)
	BrandCreate(ctx *gin.Context)
	UnitCreate(ctx *gin.Context)
	GetProducts(ctx *gin.Context)
	GetProduct(ctx *gin.Context)
}
