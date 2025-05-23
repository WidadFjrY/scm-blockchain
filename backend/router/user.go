package router

import (
	"scm-blockchain-ethereum/internal/controller"
	"scm-blockchain-ethereum/internal/middleware"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func UserRouter(router *gin.Engine, db *gorm.DB, contrl controller.UserController) {
	router.POST("/api/user/register", contrl.Register)
	router.POST("/api/user/validate", contrl.ValidateUser)
	router.POST("/api/user/login", contrl.Login)

	auth := router.Group("/")
	auth.Use(middleware.Auth(db))
	{
		auth.GET("/api/checkToken", contrl.CheckToken)
		auth.GET("/api/user", contrl.GetUserByIdToken)
		auth.GET("/api/user/:user_id", contrl.GetUserById)
		auth.GET("/api/users", contrl.GetAll)
		auth.GET("/api/user/eth_addr/:addr", contrl.GetUserByETHAddr)
		auth.GET("/api/user/count", contrl.GetCountUser)
		auth.POST("/api/user/tx", contrl.CreateUserTx)
		auth.GET("/api/user/tx/:block_number", contrl.GetUserTx)
		// 	auth.GET("/api/users/manager", contrl.GetUserByManager)
		// 	auth.PUT("/api/user", contrl.UpdateById)
		// 	auth.PUT("/api/user/password", contrl.UpdatePasswordById)
		auth.POST("/api/user/logout", contrl.Logout)
	}
}
