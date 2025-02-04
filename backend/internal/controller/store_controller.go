package controller

import "github.com/gin-gonic/gin"

type StoreController interface {
	Create(ctx *gin.Context)
	GetAll(ctx *gin.Context)
	GetById(ctx *gin.Context)
}
