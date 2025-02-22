package controller

import (
	"fmt"
	"net/http"
	"scm-blockchain-ethereum/domain/web"
	"scm-blockchain-ethereum/internal/service"
	"scm-blockchain-ethereum/pkg/helper"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

type ProductControllerImpl struct {
	ProductServ service.ProductService
}

func NewProductContrller(productServ service.ProductService) ProductController {
	return &ProductControllerImpl{ProductServ: productServ}
}

func (contr *ProductControllerImpl) Create(ctx *gin.Context) {
	role, _ := ctx.Get("role")

	_ = godotenv.Load()

	priceInt, _ := strconv.Atoi(ctx.PostForm("price"))
	stockInt, _ := strconv.Atoi(ctx.PostForm("stock"))

	request := web.ProductCreateRequest{
		ProductName:   ctx.PostForm("product_name"),
		DistributorID: ctx.PostForm("distributor_id"),
		Brand:         ctx.PostForm("brand"),
		Price:         float32(priceInt),
		Unit:          ctx.PostForm("unit"),
		Stock:         stockInt,
		Status:        ctx.PostForm("status"),
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

	response := contr.ProductServ.Create(ctx.Request.Context(), request, filePath, role.(string))
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

func (contr *ProductControllerImpl) FileUpload(ctx *gin.Context) {

}
