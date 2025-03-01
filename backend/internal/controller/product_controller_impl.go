package controller

import (
	"net/http"
	"scm-blockchain-ethereum/domain/web"
	"scm-blockchain-ethereum/internal/service"
	"scm-blockchain-ethereum/pkg/helper"

	"github.com/gin-gonic/gin"
)

type ProductControllerImpl struct {
	Serv service.ProductService
}

func NewProductController(serv service.ProductService) ProductController {
	return &ProductControllerImpl{Serv: serv}
}

func (contr *ProductControllerImpl) ProductCreate(ctx *gin.Context) {
	var request web.ProductCreateRequest
	err := ctx.ShouldBind(&request)
	helper.Err(err)

	response := contr.Serv.ProductCreate(ctx.Request.Context(), request, "")
	helper.Response(ctx, http.StatusCreated, "Created", response)
}

func (contr *ProductControllerImpl) BrandCreate(ctx *gin.Context) {
	var request web.BrandCreateRequest
	err := ctx.ShouldBind(&request)
	helper.Err(err)

	response := contr.Serv.BrandCreate(ctx.Request.Context(), request)
	helper.Response(ctx, http.StatusCreated, "Created", response)
}

func (contr *ProductControllerImpl) UnitCreate(ctx *gin.Context) {
	var request web.UnitCreateRequest
	err := ctx.ShouldBind(&request)
	helper.Err(err)

	response := contr.Serv.UnitCreate(ctx.Request.Context(), request)
	helper.Response(ctx, http.StatusCreated, "Created", response)
}

func (contr *ProductControllerImpl) GetProducts(ctx *gin.Context) {
	response := contr.Serv.GetProducts(ctx.Request.Context())
	helper.Response(ctx, http.StatusOK, "Ok", response)
}

func (contr *ProductControllerImpl) GetProduct(ctx *gin.Context) {
	productId := ctx.Params.ByName("productId")

	response := contr.Serv.GetProduct(ctx.Request.Context(), productId)
	helper.Response(ctx, http.StatusOK, "Ok", response)
}
