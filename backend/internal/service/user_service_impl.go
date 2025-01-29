package service

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"scm-blockchain-ethereum/domain/model"
	"scm-blockchain-ethereum/domain/web"
	"scm-blockchain-ethereum/internal/repository"
	"scm-blockchain-ethereum/pkg/exception"
	"scm-blockchain-ethereum/pkg/helper"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserServiceImpl struct {
	DB        *gorm.DB
	Validator *validator.Validate
	UserRepo  repository.UserRepository
}

func NewUserService(db *gorm.DB, validator *validator.Validate, userRepo repository.UserRepository) UserService {
	return &UserServiceImpl{DB: db, Validator: validator, UserRepo: userRepo}
}

func (serv *UserServiceImpl) Create(ctx context.Context, request web.UserCreateRequest) web.UserCreateResponse {
	valErr := serv.Validator.Struct(&request)
	helper.ValError(valErr)

	if request.Password != request.VerifyPassword {
		panic(exception.NewBadRequestError("password doesn't match"))
	}

	hasedPass, err := bcrypt.GenerateFromPassword([]byte(request.Password), 12)
	helper.Err(err)

	rand.NewSource(time.Now().UnixNano())

	var userId string
	isExsit := true

	for isExsit {
		userId = helper.GenerateRandomString(15)
		txErr := serv.DB.Transaction(func(tx *gorm.DB) error {
			isExsit = !serv.UserRepo.IsUserExsit(ctx, tx, userId)
			return nil
		})
		helper.Err(txErr)
	}

	userName := fmt.Sprintf("User_%s", helper.GenerateRandomString(9))

	user := model.User{
		ID:       userId,
		Email:    request.Email,
		Role:     request.Role,
		Password: string(hasedPass),
		Name:     userName,
	}

	var userResult model.User
	txErr := serv.DB.Transaction(func(tx *gorm.DB) error {
		userResult = serv.UserRepo.Create(ctx, tx, user)
		return nil
	})
	helper.Err(txErr)

	return web.UserCreateResponse{
		Email:     userResult.Email,
		Name:      userResult.Name,
		CreatedAt: userResult.CreatedAt,
	}

}

func (serv *UserServiceImpl) LoginUser(ctx context.Context, request web.UserLoginRequest) web.UserLoginResponse {
	godotenv.Load()

	tokenRepo := repository.NewTokenRepository()
	tokenServ := NewTokenService(serv.DB, tokenRepo)

	valErr := serv.Validator.Struct(&request)
	helper.ValError(valErr)

	var user model.User
	txErr := serv.DB.Transaction(func(tx *gorm.DB) error {
		user = serv.UserRepo.GetUserByEmail(ctx, tx, request.Email)
		return nil
	})
	helper.Err(txErr)

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password))
	if err != nil {
		panic(exception.NewUnauthorized("email or password wrong!"))
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, web.JwtClaims{
		Email:  user.Email,
		UserId: user.ID,
		Role:   user.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		},
	})

	tokenStr, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))
	if err != nil {
		panic(err.Error())
	}

	tokenServ.Create(ctx, tokenStr)

	return web.UserLoginResponse{
		Token: tokenStr,
	}
}

func (serv *UserServiceImpl) GetUserById(ctx context.Context, userId string) web.UserGetResponse {
	var user model.User
	txErr := serv.DB.Transaction(func(tx *gorm.DB) error {
		user = serv.UserRepo.GetUserById(ctx, tx, userId)
		return nil
	})
	helper.Err(txErr)

	return web.UserGetResponse{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Role:      user.Role,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

func (serv *UserServiceImpl) LogoutUser(ctx context.Context, token string, userId string) web.UserLogoutResponse {
	tokenRepo := repository.NewTokenRepository()
	tokenServ := NewTokenService(serv.DB, tokenRepo)

	tokenServ.Update(ctx, token)

	return web.UserLogoutResponse{
		Id:          userId,
		LoggedOutAt: time.Now(),
	}
}

func (serv *UserServiceImpl) GetAll(ctx context.Context) []web.UserGetResponse {
	var users []web.UserGetResponse

	txErr := serv.DB.Transaction(func(tx *gorm.DB) error {
		for _, user := range serv.UserRepo.GetAll(ctx, tx) {
			userResp := web.UserGetResponse{
				ID:        user.ID,
				Name:      user.Name,
				Email:     user.Email,
				Role:      user.Role,
				CreatedAt: user.CreatedAt,
				UpdatedAt: user.UpdatedAt,
			}

			users = append(users, userResp)
		}
		return nil
	})
	helper.Err(txErr)

	return users
}
