package web

import "github.com/golang-jwt/jwt/v5"

type JwtClaims struct {
	UserId string
	Email  string
	Role   string
	jwt.RegisteredClaims
}
