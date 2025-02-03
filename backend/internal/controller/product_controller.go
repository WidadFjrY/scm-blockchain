package controller

import "github.com/gin-gonic/gin"

type ProductController interface {
	Create(ctx *gin.Context)
	GetAll(ctx *gin.Context)
	GetById(ctx *gin.Context)
	FileUpload(ctx *gin.Context)
}
