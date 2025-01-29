package controller

import (
	"encoding/json"
	"net/http"
	"scm-blockchain-ethereum/domain/web"
	"scm-blockchain-ethereum/internal/service"
	"scm-blockchain-ethereum/pkg/helper"

	"github.com/gin-gonic/gin"
)

type ProductControllerImpl struct {
	ProductServ service.ProductService
}

func NewProductContrller(productServ service.ProductService) ProductController {
	return &ProductControllerImpl{ProductServ: productServ}
}

func (contr *ProductControllerImpl) Create(ctx *gin.Context) {
	var request web.ProductCreateRequest
	role, _ := ctx.Get("role")

	decoder := json.NewDecoder(ctx.Request.Body)
	err := decoder.Decode(&request)
	helper.Err(err)

	response := contr.ProductServ.Create(ctx.Request.Context(), request, role.(string))
	helper.Response(ctx, http.StatusCreated, "Created", response)
}

func (contr *ProductControllerImpl) GetAll(ctx *gin.Context) {
	response := contr.ProductServ.GetAll(ctx)
	helper.Response(ctx, http.StatusOK, "Ok", response)
}

func (contr *ProductControllerImpl) GetById(ctx *gin.Context) {
	productId := ctx.Params.ByName("product_id")

	response := contr.ProductServ.GetById(ctx.Request.Context(), productId)
	helper.Response(ctx, http.StatusOK, "Ok", response)
}
