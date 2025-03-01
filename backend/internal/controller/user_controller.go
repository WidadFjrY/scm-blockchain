package controller

import "github.com/gin-gonic/gin"

type UserController interface {
	Register(ctx *gin.Context)
	Login(ctx *gin.Context)
	CheckToken(ctx *gin.Context)
	ValidateUser(ctx *gin.Context)
	Logout(ctx *gin.Context)
	GetUserByIdToken(ctx *gin.Context)
	GetUserById(ctx *gin.Context)
	GetAll(ctx *gin.Context)
	// GetUserByManager(ctx *gin.Context)
	// UpdateById(ctx *gin.Context)
	// UpdatePasswordById(ctx *gin.Context)
}
