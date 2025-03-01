package service

import (
	"context"
	"fmt"
	"scm-blockchain-ethereum/domain/model"
	"scm-blockchain-ethereum/domain/web"
	"scm-blockchain-ethereum/internal/repository"
	"scm-blockchain-ethereum/pkg/helper"
	"time"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type ProductServiceImpl struct {
	DB        *gorm.DB
	Validator *validator.Validate
	Repo      repository.ProductRepository
}

func NewProductService(db *gorm.DB, validator *validator.Validate, repo repository.ProductRepository) ProductService {
	return &ProductServiceImpl{DB: db, Validator: validator, Repo: repo}
}

func (serv *ProductServiceImpl) ProductCreate(ctx context.Context, request web.ProductCreateRequest, filePath string) web.ProductCreateResponse {
	valErr := serv.Validator.Struct(&request)
	helper.ValError(valErr)

	productId := fmt.Sprintf("PRD%s", helper.GenerateRandomString(9))
	priceId := fmt.Sprintf("PRC%s", helper.GenerateRandomString(9))
	stockId := fmt.Sprintf("STK%s", helper.GenerateRandomString(9))

	helper.Err(serv.DB.Transaction(func(tx *gorm.DB) error {
		serv.Repo.ProductCreate(ctx, tx, model.NormalizedProduct{
			Product: model.Product{
				ID:          productId,
				ProductName: request.ProductName,
				BrandID:     request.BrandId,
				UnitID:      request.UnitId,
				Description: request.Description,
				PicturePath: filePath,
				CreatedAt:   time.Now(),
				UpdatedAt:   time.Now(),
			},
			ProductPrice: model.ProductPrice{
				ID:        priceId,
				ProductID: productId,
				Price:     request.Price,
				UpdatedAt: time.Now(),
			},
			ProductStock: model.ProductStock{
				ID:         stockId,
				ProductID:  productId,
				StockIn:    request.Stock,
				StockTotal: request.Stock,
				UpdatedAt:  time.Now(),
			},
		})
		return nil
	}))

	return web.ProductCreateResponse{
		ProductName: request.ProductName,
		CreatedAt:   time.Now(),
	}
}

func (serv *ProductServiceImpl) BrandCreate(ctx context.Context, request web.BrandCreateRequest) web.BrandCreateResponse {
	valErr := serv.Validator.Struct(&request)
	helper.ValError(valErr)

	brandId := fmt.Sprintf("BRN%s", helper.GenerateRandomString(9))

	helper.Err(serv.DB.Transaction(func(tx *gorm.DB) error {
		serv.Repo.BrandCreate(ctx, tx, model.ProductBrand{
			ID:   brandId,
			Name: request.BrandName,
		})
		return nil
	}))

	return web.BrandCreateResponse{
		BrandName: request.BrandName,
		CreatedAt: time.Now(),
	}
}

func (serv *ProductServiceImpl) UnitCreate(ctx context.Context, request web.UnitCreateRequest) web.UnitCreateResponse {
	valErr := serv.Validator.Struct(&request)
	helper.ValError(valErr)

	unitId := fmt.Sprintf("UNT%s", helper.GenerateRandomString(9))

	helper.Err(serv.DB.Transaction(func(tx *gorm.DB) error {
		serv.Repo.UnitCreate(ctx, tx, model.ProductUnit{
			ID:   unitId,
			Name: request.UnitName,
		})
		return nil
	}))

	return web.UnitCreateResponse{
		UnitName:  request.UnitName,
		CreatedAt: time.Now(),
	}
}

func (serv *ProductServiceImpl) GetProducts(ctx context.Context) []web.ProductResponse {
	var products []web.ProductResponse

	helper.Err(serv.DB.Transaction(func(tx *gorm.DB) error {
		for _, product := range serv.Repo.GetProducts(ctx, tx) {
			products = append(products, web.ProductResponse{
				ID:          product.ID,
				ProductName: product.ProductName,
				Stock:       product.Stocks[0].StockTotal,
				Unit:        product.Unit.Name,
				Brand:       product.Brand.Name,
				FilePath:    product.PicturePath,
				Price:       product.Prices[0].Price,
				Description: product.Description,
				CreatedAt:   product.CreatedAt,
			})
		}
		return nil
	}))

	return products
}

func (serv *ProductServiceImpl) GetProduct(ctx context.Context, productId string) web.ProductResponse {
	var product web.ProductResponse

	helper.Err(serv.DB.Transaction(func(tx *gorm.DB) error {
		productModel := serv.Repo.GetProduct(ctx, tx, productId)
		product = web.ProductResponse{
			ID:          productModel.ID,
			ProductName: productModel.ProductName,
			Stock:       productModel.Stocks[0].StockTotal,
			Unit:        productModel.Unit.Name,
			Brand:       productModel.Brand.Name,
			FilePath:    productModel.PicturePath,
			Price:       productModel.Prices[0].Price,
			Description: productModel.Description,
			CreatedAt:   productModel.CreatedAt,
		}

		return nil
	}))

	return product
}
