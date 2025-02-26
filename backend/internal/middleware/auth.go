package middleware

import (
	"net/http"
	"scm-blockchain-ethereum/domain/web"
	"scm-blockchain-ethereum/internal/repository"
	"scm-blockchain-ethereum/internal/service"
	"scm-blockchain-ethereum/pkg/helper"

	"strings"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

var adminOnlyPaths = []string{
	"/api/product/add",
	"/api/brand/add",
	"/api/unit/add",
}

func Auth(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		godotenv.Load()

		tokenRepo := repository.NewTokenRepository()
		tokenSrvc := service.NewTokenService(db, tokenRepo)

		authToken := ctx.GetHeader("Authorization")

		if authToken == "" {
			ctx.JSON(http.StatusUnauthorized, web.ErrorResponse{
				Code:    http.StatusUnauthorized,
				Status:  "Unauthorized",
				Message: "missing authorization header",
			})
			ctx.Abort()
			return
		}

		tokenParts := strings.Split(authToken, " ")
		if len(tokenParts) != 2 || strings.ToLower(tokenParts[0]) != "bearer" {
			ctx.JSON(http.StatusUnauthorized, web.ErrorResponse{
				Code:    http.StatusUnauthorized,
				Status:  "Unauthorized",
				Message: "invalid authorization header format",
			})
			ctx.Abort()
			return
		}

		token := tokenParts[1]
		isValid := tokenSrvc.IsValid(ctx.Request.Context(), token)

		claims, err := helper.ValidationToken(token)
		if err != nil || !isValid {
			ctx.JSON(http.StatusUnauthorized, web.ErrorResponse{
				Code:    http.StatusUnauthorized,
				Status:  "Unauthorized",
				Message: "invalid or expired token",
			})
			ctx.Abort()
			return
		}

		for _, path := range adminOnlyPaths {
			if strings.HasPrefix(ctx.Request.URL.Path, path) {
				if claims.Role != "Admin" {
					ctx.JSON(http.StatusForbidden, web.ErrorResponse{
						Code:    http.StatusForbidden,
						Status:  "Forbidden",
						Message: "You are not authorized",
					})
					ctx.Abort()
					return
				}
			}
		}

		ctx.Set("email", claims.Email)
		ctx.Set("user_id", claims.UserId)
		ctx.Set("role", claims.Role)

		ctx.Next()
	}
}
