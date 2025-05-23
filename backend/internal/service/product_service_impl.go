package service

import (
	"context"
	"fmt"
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
	DB        *gorm.DB
	Validator *validator.Validate
	Repo      repository.ProductRepository
}

func NewProductService(db *gorm.DB, validator *validator.Validate, repo repository.ProductRepository) ProductService {
	return &ProductServiceImpl{DB: db, Validator: validator, Repo: repo}
}

func (serv *ProductServiceImpl) ProductCreate(ctx context.Context, request web.ProductCreateRequest, filePath string, role string) web.ProductCreateResponse {
	if role != "Warehouse_Staff" {
		panic(exception.NewUnauthorized("no authority"))
	}
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
		ProductId: productId,
		CreatedAt: time.Now(),
	}
}

func (serv *ProductServiceImpl) BrandCreate(ctx context.Context, request web.BrandCreateRequest, role string) web.BrandCreateResponse {
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
		ID:        brandId,
		BrandName: request.BrandName,
	}
}

func (serv *ProductServiceImpl) UnitCreate(ctx context.Context, request web.UnitCreateRequest, role string) web.UnitCreateResponse {
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
		ID:       unitId,
		UnitName: request.UnitName,
	}
}

func (serv *ProductServiceImpl) GetProducts(ctx context.Context) []web.ProductResponse {
	var products []web.ProductResponse

	helper.Err(serv.DB.Transaction(func(tx *gorm.DB) error {
		for _, product := range serv.Repo.GetProducts(ctx, tx) {
			var stockTotal int
			if len(product.Stocks) > 0 {
				stockTotal = product.Stocks[0].StockTotal
			} else {
				stockTotal = 0
			}

			var price float32
			if len(product.Prices) > 0 {
				price = product.Prices[0].Price
			} else {
				price = 0
			}

			products = append(products, web.ProductResponse{
				ID:          product.ID,
				ProductName: product.ProductName,
				Stock:       stockTotal,
				Unit:        product.Unit.Name,
				Brand:       product.Brand.Name,
				FilePath:    product.PicturePath,
				Price:       price,
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

func (serv *ProductServiceImpl) GetTotalProduct(ctx context.Context) web.GetTotalProductResponse {
	var totalProduct web.GetTotalProductResponse

	err := serv.DB.Transaction(func(tx *gorm.DB) error {
		totalProduct.TotalProduct = serv.Repo.GetTotalProducts(ctx, tx)
		return nil
	})

	helper.Err(err)
	return totalProduct
}

func (serv *ProductServiceImpl) DeleteProductById(ctx context.Context, productId string) {
	helper.Err(serv.DB.Transaction(func(tx *gorm.DB) error {
		serv.Repo.DeleteProductById(ctx, tx, productId)
		return nil
	}))
}

func (serv *ProductServiceImpl) GetBrands(ctx context.Context) []web.BrandResponse {
	var brands []web.BrandResponse

	helper.Err(serv.DB.Transaction(func(tx *gorm.DB) error {
		for _, brand := range serv.Repo.GetBrands(ctx, tx) {
			brands = append(brands, web.BrandResponse{
				ID:        brand.ID,
				BrandName: brand.Name,
			})
		}
		return nil
	}))

	return brands
}

func (serv *ProductServiceImpl) GetUnits(ctx context.Context) []web.UnitResponse {
	var units []web.UnitResponse

	helper.Err(serv.DB.Transaction(func(tx *gorm.DB) error {
		for _, unit := range serv.Repo.GetUnits(ctx, tx) {
			units = append(units, web.UnitResponse{
				ID:       unit.ID,
				UnitName: unit.Name,
			})
		}
		return nil
	}))

	return units
}

func (serv *ProductServiceImpl) AddToCart(ctx context.Context, request web.AddToCartRequest, userId string) {
	helper.ValError(serv.Validator.Struct(&request))

	helper.Err(serv.DB.Transaction(func(tx *gorm.DB) error {
		serv.Repo.AddToCart(ctx, tx, model.Cart{
			UserID:    userId,
			ProductID: request.ProductId,
			Quantity:  request.Quantity,
		})
		return nil
	}))
}

func (serv *ProductServiceImpl) GetCarts(ctx context.Context, userId string) []web.CartResponse {
	var Carts []web.CartResponse

	helper.Err(serv.DB.Transaction(func(tx *gorm.DB) error {
		for _, cart := range serv.Repo.GetCarts(ctx, tx, userId) {
			Carts = append(Carts, web.CartResponse{
				UserID:    cart.UserID,
				ProductId: cart.ProductID,
				Quantity:  cart.Quantity,
				Product: web.ProductResponse{
					ID:          cart.Product.ID,
					ProductName: cart.Product.ProductName,
					Price:       cart.Product.Prices[0].Price,
					Stock:       cart.Product.Stocks[0].StockTotal,
					Brand:       cart.Product.Brand.Name,
					Unit:        cart.Product.Unit.Name,
					FilePath:    cart.Product.PicturePath,
				},
			})
		}
		return nil
	}))

	return Carts
}

func (serv *ProductServiceImpl) UpdateCartQty(ctx context.Context, productId string, qty int) {
	helper.Err(serv.DB.Transaction(func(tx *gorm.DB) error {
		serv.Repo.UpdateCartQty(ctx, tx, model.Cart{
			ProductID: productId,
			Quantity:  qty,
		})
		return nil
	}))
}

func (serv *ProductServiceImpl) DeleteItemCart(ctx context.Context, productId string) {
	helper.Err(serv.DB.Transaction(func(tx *gorm.DB) error {
		serv.Repo.DeleteItemCart(ctx, tx, productId)
		return nil
	}))
}

func (serv *ProductServiceImpl) UpdateProductStock(ctx context.Context, role string, request web.UpdateProductStockRequest) {
	if request.StockIn > 0 {
		if role != "Warehouse_Staff" {
			panic(exception.NewUnauthorized("no authority"))
		}
	}

	helper.ValError(serv.Validator.Struct(&request))
	err := serv.DB.Transaction(func(tx *gorm.DB) error {
		serv.Repo.UpdateProductStock(ctx, tx, model.ProductStock{
			ID:        helper.GenerateRandomString(15),
			ProductID: request.ProductID,
			StockIn:   request.StockIn,
			StockOut:  request.StockOut,
		})
		return nil
	})

	helper.Err(err)
}
