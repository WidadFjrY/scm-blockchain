package service

import (
	"context"
	"scm-blockchain-ethereum/domain/web"
)

type UserService interface {
	Create(ctx context.Context, request web.UserCreateRequest) web.UserCreateResponse
	LoginUser(ctx context.Context, request web.UserLoginRequest) web.UserLoginResponse
	LogoutUser(ctx context.Context, token string, userId string) web.UserLogoutResponse
	GetUserById(ctx context.Context, userId string) web.UserGetResponse
	GetAll(ctx context.Context) []web.UserGetResponse
	GetUserByManager(ctx context.Context, role string) []web.UserGetResponse
}
