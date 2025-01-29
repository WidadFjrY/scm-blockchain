package controller

import (
	"encoding/json"
	"net/http"
	"scm-blockchain-ethereum/domain/web"
	"scm-blockchain-ethereum/internal/service"
	"scm-blockchain-ethereum/pkg/helper"

	"github.com/gin-gonic/gin"
)

type DistributorControllerImpl struct {
	DistributorServ service.DistributorService
}

func NewDistributorController(distributorServ service.DistributorService) DistributorController {
	return &DistributorControllerImpl{DistributorServ: distributorServ}
}

func (contrl *DistributorControllerImpl) Create(ctx *gin.Context) {
	userId, _ := ctx.Get("user_id")

	var request web.DistributorCreateRequest

	decoder := json.NewDecoder(ctx.Request.Body)
	err := decoder.Decode(&request)
	helper.Err(err)

	response := contrl.DistributorServ.Create(ctx.Request.Context(), request, userId.(string))
	helper.Response(ctx, http.StatusCreated, "Created", response)
}

func (contrl *DistributorControllerImpl) GetAll(ctx *gin.Context) {
	response := contrl.DistributorServ.GetAll(ctx.Request.Context())
	helper.Response(ctx, http.StatusOK, "Ok", response)
}

func (contrl *DistributorControllerImpl) GetById(ctx *gin.Context) {
	distributorId := ctx.Params.ByName("distId")
	response := contrl.DistributorServ.GetByID(ctx.Request.Context(), distributorId)

	helper.Response(ctx, http.StatusOK, "Ok", response)
}
