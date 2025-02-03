package service

import (
	"context"
	"math/rand"
	"scm-blockchain-ethereum/domain/model"
	"scm-blockchain-ethereum/domain/web"
	"scm-blockchain-ethereum/internal/repository"
	"scm-blockchain-ethereum/pkg/exception"
	"scm-blockchain-ethereum/pkg/helper"
	"time"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type ProductServiceImpl struct {
	DB          *gorm.DB
	Validator   *validator.Validate
	ProductRepo repository.ProductRepository
}

func NewProductService(db *gorm.DB, validator *validator.Validate, productRepo repository.ProductRepository) ProductService {
	return &ProductServiceImpl{DB: db, Validator: validator, ProductRepo: productRepo}
}

func (serv *ProductServiceImpl) Create(ctx context.Context, request web.ProductCreateRequest, filePath string, role string) web.ProductCreateResponse {
	valErr := serv.Validator.Struct(&request)
	helper.ValError(valErr)

	if role != "Admin" && role != "Distributor" {
		panic(exception.NewBadRequestError("have no authority"))
	}

	var product model.Product

	rand.NewSource(time.Now().Unix())
	err := serv.DB.Transaction(func(tx *gorm.DB) error {
		product = serv.ProductRepo.Create(ctx, tx, model.Product{
			ID:            helper.GenerateRandomString(15),
			DistributorID: request.DistributorID,
			ProductName:   request.ProductName,
			Brand:         request.Brand,
			Stock:         request.Stock,
			Price:         request.Price,
			Unit:          request.Unit,
			Status:        request.Status,
			UrlPricture:   filePath,
		})

		return nil
	})
	helper.Err(err)

	return web.ProductCreateResponse{
		ProductName: product.ProductName,
		CreatedAt:   product.CreatedAt,
	}
}

func (serv *ProductServiceImpl) GetAll(ctx context.Context) []web.ProductResponse {
	var products []web.ProductResponse

	err := serv.DB.Transaction(func(tx *gorm.DB) error {
		for _, productModel := range serv.ProductRepo.GetAll(ctx, tx) {
			product := web.ProductResponse{
				ID:          productModel.ID,
				ProductName: productModel.ProductName,
				Brand:       productModel.Brand,
				Price:       helper.FormatRupiah(productModel.Price),
				Stock:       productModel.Stock,
				Unit:        productModel.Unit,
				Status:      productModel.Status,
				UrlPricture: productModel.UrlPricture,
				CreatedAt:   productModel.CreatedAt,
				UpdatedAt:   productModel.UpdatedAt,
			}
			products = append(products, product)
		}
		return nil
	})
	helper.Err(err)
	return products
}

func (serv *ProductServiceImpl) GetById(ctx context.Context, productId string) web.ProductResponse {
	var product web.ProductResponse

	err := serv.DB.Transaction(func(tx *gorm.DB) error {
		productModel := serv.ProductRepo.GetById(ctx, tx, productId)
		product = web.ProductResponse{
			ID:             productModel.ID,
			DistrbutorName: productModel.Distributor.Name,
			ProductName:    productModel.ProductName,
			Brand:          productModel.Brand,
			Price:          helper.FormatRupiah(productModel.Price),
			Stock:          productModel.Stock,
			Unit:           productModel.Unit,
			Status:         productModel.Status,
			UrlPricture:    productModel.UrlPricture,
			CreatedAt:      productModel.CreatedAt,
			UpdatedAt:      productModel.UpdatedAt,
		}
		return nil
	})
	helper.Err(err)
	return product
}
