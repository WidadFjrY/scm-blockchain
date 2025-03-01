package service

import (
	"context"
	"scm-blockchain-ethereum/domain/web"
)

type UserService interface {
	Create(ctx context.Context, request web.UserCreateRequest) web.UserCreateResponse
	ValidateUser(ctx context.Context, request web.UserCreateRequest) bool
	LoginUser(ctx context.Context, request web.UserLoginRequest) web.UserLoginResponse
	LogoutUser(ctx context.Context, token string, userId string) web.UserLogoutResponse
	GetUserById(ctx context.Context, userId string) web.UserGetResponse
	GetAll(ctx context.Context) []web.UserGetResponse
	// GetUserByManager(ctx context.Context, role string) []web.UserGetResponse
	// UpdateById(ctx context.Context, request web.UserUpdateRequest, userId string) time.Time
	// UpdatePassword(ctx context.Context, request web.UserUpdatePasswordRequest, userId string) time.Time
}
