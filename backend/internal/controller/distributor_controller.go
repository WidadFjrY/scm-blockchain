package controller

import "github.com/gin-gonic/gin"

type DistributorController interface {
	Create(ctx *gin.Context)
	GetAll(ctx *gin.Context)
	GetById(ctx *gin.Context)
}
