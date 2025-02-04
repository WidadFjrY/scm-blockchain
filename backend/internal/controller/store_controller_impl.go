package controller

import (
	"net/http"
	"scm-blockchain-ethereum/domain/web"
	"scm-blockchain-ethereum/internal/service"
	"scm-blockchain-ethereum/pkg/helper"
	"time"

	"github.com/gin-gonic/gin"
)

type StoreControllerImpl struct {
	StoreServ service.StoreService
}

func NewStoreController(storeServ service.StoreService) StoreController {
	return &StoreControllerImpl{StoreServ: storeServ}
}

func (controller *StoreControllerImpl) Create(ctx *gin.Context) {
	var request web.StoreCreateRequest
	err := ctx.ShouldBind(&request)
	helper.Err(err)

	role, _ := ctx.Get("role")

	controller.StoreServ.Create(ctx.Request.Context(), role.(string), request)
	helper.Response(ctx, http.StatusCreated, "Created", time.Now())
}

func (controller *StoreControllerImpl) GetAll(ctx *gin.Context) {
	response := controller.StoreServ.GetAll(ctx.Request.Context())
	helper.Response(ctx, http.StatusOK, "Ok", response)
}

func (controller *StoreControllerImpl) GetById(ctx *gin.Context) {
	storeId := ctx.Params.ByName("store_id")

	response := controller.StoreServ.GetById(ctx.Request.Context(), storeId)
	helper.Response(ctx, http.StatusOK, "Ok", response)
}
