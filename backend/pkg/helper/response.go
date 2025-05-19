package helper

import (
	"scm-blockchain-ethereum/domain/web"

	"github.com/gin-gonic/gin"
)

func Response(ctx *gin.Context, code int, status string, response interface{}) {
	ctx.JSON(code, web.SuccessResponse{
		Code:   code,
		Status: status,
		Data:   response,
	})
}
