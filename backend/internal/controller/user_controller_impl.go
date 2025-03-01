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
	role, _ := ctx.Get("role")
	helper.Response(ctx, http.StatusOK, "OK", role.(string))
}

func (contr *UserControllerImpl) GetUserByIdToken(ctx *gin.Context) {
	userId, _ := ctx.Get("user_id")

	response := contr.UserServ.GetUserById(ctx.Request.Context(), userId.(string))
	helper.Response(ctx, http.StatusOK, "Ok", response)
}

func (contr *UserControllerImpl) ValidateUser(ctx *gin.Context) {
	var request web.UserCreateRequest
	err := ctx.ShouldBind(&request)
	helper.Err(err)

	response := contr.UserServ.ValidateUser(ctx.Request.Context(), request)
	helper.Response(ctx, http.StatusOK, "Ok", response)
}

func (contr *UserControllerImpl) Logout(ctx *gin.Context) {
	token := strings.Split(ctx.GetHeader("Authorization"), " ")[1]
	userId, _ := ctx.Get("user_id")

	response := contr.UserServ.LogoutUser(ctx.Request.Context(), token, userId.(string))
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

/*
func (contr *UserControllerImpl) GetUserByManager(ctx *gin.Context) {
	role, _ := ctx.Get("role")
	response := contr.UserServ.GetUserByManager(ctx.Request.Context(), role.(string))
	helper.Response(ctx, http.StatusOK, "Ok", response)
}

func (contr *UserControllerImpl) UpdateById(ctx *gin.Context) {
	userId, _ := ctx.Get("user_id")

	var request web.UserUpdateRequest
	err := ctx.ShouldBind(&request)
	helper.Err(err)

	response := contr.UserServ.UpdateById(ctx.Request.Context(), request, userId.(string))
	helper.Response(ctx, http.StatusOK, "Ok", response)
}

func (contr *UserControllerImpl) UpdatePasswordById(ctx *gin.Context) {
	userId, _ := ctx.Get("user_id")

	var request web.UserUpdatePasswordRequest
	err := ctx.ShouldBind(&request)
	helper.Err(err)

	response := contr.UserServ.UpdatePassword(ctx.Request.Context(), request, userId.(string))
	helper.Response(ctx, http.StatusOK, "Ok", response)
}
*/
