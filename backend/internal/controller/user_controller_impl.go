package controller

import (
	"encoding/json"
	"net/http"
	"scm-blockchain-ethereum/domain/web"
	"scm-blockchain-ethereum/internal/service"
	"scm-blockchain-ethereum/pkg/exception"
	"scm-blockchain-ethereum/pkg/helper"
	"strings"

	"github.com/gin-gonic/gin"
)

type UserControllerImpl struct {
	UserServ service.UserService
}

func NewUserController(userServ service.UserService) UserController {
	return &UserControllerImpl{UserServ: userServ}
}

func (contr UserControllerImpl) Register(ctx *gin.Context) {
	var request web.UserCreateRequest

	if ctx.Request.ContentLength == 0 {
		panic(exception.NewBadRequestError("request body is empty"))
	}

	decoder := json.NewDecoder(ctx.Request.Body)
	err := decoder.Decode(&request)
	helper.Err(err)

	response := contr.UserServ.Create(ctx.Request.Context(), request)
	helper.Response(ctx, http.StatusCreated, "Created", response)
}

func (contr *UserControllerImpl) Login(ctx *gin.Context) {
	var request web.UserLoginRequest

	if ctx.Request.ContentLength == 0 {
		panic(exception.NewBadRequestError("request body is empty"))
	}

	decoder := json.NewDecoder(ctx.Request.Body)
	err := decoder.Decode(&request)
	helper.Err(err)

	response := contr.UserServ.LoginUser(ctx.Request.Context(), request)
	helper.Response(ctx, http.StatusOK, "OK", response)
}

func (contr *UserControllerImpl) CheckToken(ctx *gin.Context) {
	helper.Response(ctx, http.StatusOK, "OK", "")
}

func (contr *UserControllerImpl) Logout(ctx *gin.Context) {
	token := strings.Split(ctx.GetHeader("Authorization"), " ")[1]
	userId, _ := ctx.Get("user_id")

	response := contr.UserServ.LogoutUser(ctx.Request.Context(), token, userId.(string))
	helper.Response(ctx, http.StatusOK, "Ok", response)
}

func (contr *UserControllerImpl) GetUserByIdToken(ctx *gin.Context) {
	userId, _ := ctx.Get("user_id")

	response := contr.UserServ.GetUserById(ctx.Request.Context(), userId.(string))
	helper.Response(ctx, http.StatusOK, "Ok", response)
}

func (contr *UserControllerImpl) GetAll(ctx *gin.Context) {
	response := contr.UserServ.GetAll(ctx.Request.Context())
	helper.Response(ctx, http.StatusOK, "Ok", response)
}

func (contr *UserControllerImpl) GetUserById(ctx *gin.Context) {
	userId := ctx.Params.ByName("user_id")

	response := contr.UserServ.GetUserById(ctx.Request.Context(), userId)
	helper.Response(ctx, http.StatusOK, "Ok", response)
}
