package controller

import (
	"fmt"
	"net/http"
	"scm-blockchain-ethereum/domain/web"
	"scm-blockchain-ethereum/internal/service"
	"scm-blockchain-ethereum/pkg/exception"
	"scm-blockchain-ethereum/pkg/helper"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

type ProductControllerImpl struct {
	Serv service.ProductService
}

func NewProductController(serv service.ProductService) ProductController {
	return &ProductControllerImpl{Serv: serv}
}

func (contr *ProductControllerImpl) ProductCreate(ctx *gin.Context) {
	role, _ := ctx.Get("role")

	_ = godotenv.Load()

	priceInt, _ := strconv.Atoi(ctx.PostForm("price"))
	stockInt, _ := strconv.Atoi(ctx.PostForm("stock"))

	request := web.ProductCreateRequest{
		ProductName: ctx.PostForm("product_name"),
		BrandId:     ctx.PostForm("brand_id"),
		Price:       float32(priceInt),
		UnitId:      ctx.PostForm("unit_id"),
		Stock:       stockInt,
		Description: ctx.PostForm("description"),
	}

	file, err := ctx.FormFile("image")
	var filePath string

	if err == nil {
		allowedExtensions := map[string]bool{
			"jpg":  true,
			"png":  true,
			"jpeg": true,
		}

		parts := strings.Split(file.Filename, ".")
		if len(parts) < 2 {
			helper.Response(ctx, http.StatusBadRequest, "Invalid file format", nil)
			return
		}
		extension := parts[len(parts)-1]

		if !allowedExtensions[extension] {
			helper.Response(ctx, http.StatusBadRequest, "File must be jpg, jpeg, or png", nil)
			return
		}

		filePath = fmt.Sprintf("public/product_img/%s.%s", helper.GenerateRandomString(15), extension)

		if err := ctx.SaveUploadedFile(file, filePath); err != nil {
			helper.Response(ctx, http.StatusInternalServerError, "Failed to save file", nil)
			return
		}
	}

	response := contr.Serv.ProductCreate(ctx.Request.Context(), request, filePath, role.(string))
	helper.Response(ctx, http.StatusCreated, "Created", response)
}

func (contr *ProductControllerImpl) BrandCreate(ctx *gin.Context) {
	role, _ := ctx.Get("role")

	var request web.BrandCreateRequest
	err := ctx.ShouldBind(&request)
	helper.Err(err)

	response := contr.Serv.BrandCreate(ctx.Request.Context(), request, role.(string))
	helper.Response(ctx, http.StatusCreated, "Created", response)
}

func (contr *ProductControllerImpl) UnitCreate(ctx *gin.Context) {
	role, _ := ctx.Get("role")

	var request web.UnitCreateRequest
	err := ctx.ShouldBind(&request)
	helper.Err(err)

	response := contr.Serv.UnitCreate(ctx.Request.Context(), request, role.(string))
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

func (contr *ProductControllerImpl) GetTotalProduct(ctx *gin.Context) {
	response := contr.Serv.GetTotalProduct(ctx.Request.Context())
	helper.Response(ctx, http.StatusOK, "Ok", response)
}

func (contr *ProductControllerImpl) DeleteProductById(ctx *gin.Context) {
	productId := ctx.Params.ByName("productId")

	contr.Serv.DeleteProductById(ctx.Request.Context(), productId)
	helper.Response(ctx, http.StatusOK, "Ok", "")
}

func (contr *ProductControllerImpl) GetBrands(ctx *gin.Context) {
	response := contr.Serv.GetBrands(ctx.Request.Context())
	helper.Response(ctx, http.StatusOK, "Ok", response)
}

func (contr *ProductControllerImpl) GetUnits(ctx *gin.Context) {
	response := contr.Serv.GetUnits(ctx.Request.Context())
	helper.Response(ctx, http.StatusOK, "Ok", response)
}

func (contr *ProductControllerImpl) AddToCart(ctx *gin.Context) {
	userId, _ := ctx.Get("user_id")

	var request web.AddToCartRequest
	helper.Err(ctx.ShouldBind(&request))

	contr.Serv.AddToCart(ctx.Request.Context(), request, userId.(string))
	helper.Response(ctx, http.StatusCreated, "Created", time.Now())
}

func (contr *ProductControllerImpl) GetCarts(ctx *gin.Context) {
	userId, _ := ctx.Get("user_id")
	response := contr.Serv.GetCarts(ctx.Request.Context(), userId.(string))
	helper.Response(ctx, http.StatusOK, "Ok", response)
}

func (contr *ProductControllerImpl) UpdateCartQty(ctx *gin.Context) {
	productId := ctx.Query("product_id")
	qty := ctx.Query("qty")

	qtyInt, err := strconv.Atoi(qty)
	helper.Err(err)

	contr.Serv.UpdateCartQty(ctx.Request.Context(), productId, qtyInt)
	helper.Response(ctx, http.StatusOK, "Ok", "")
}

func (contr *ProductControllerImpl) DeleteItemCart(ctx *gin.Context) {
	productId := ctx.Params.ByName("product_id")
	contr.Serv.DeleteItemCart(ctx.Request.Context(), productId)
	helper.Response(ctx, http.StatusOK, "Ok", "")
}

func (contr *ProductControllerImpl) UpdateProductStock(ctx *gin.Context) {
	var request web.UpdateProductStockRequest
	role, _ := ctx.Get("role")
	err := ctx.ShouldBind(&request)
	if err != nil {
		panic(exception.NewBadRequestError("request body required"))
	}

	contr.Serv.UpdateProductStock(ctx.Request.Context(), role.(string), request)
	helper.Response(ctx, http.StatusOK, "Ok", "")
}
