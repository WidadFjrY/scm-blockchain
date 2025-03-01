package service

import (
	"context"
	"scm-blockchain-ethereum/domain/web"
)

type ProductService interface {
	ProductCreate(ctx context.Context, request web.ProductCreateRequest, filePath string, role string) web.ProductCreateResponse
	BrandCreate(ctx context.Context, request web.BrandCreateRequest, role string) web.BrandCreateResponse
	UnitCreate(ctx context.Context, request web.UnitCreateRequest, role string) web.UnitCreateResponse

	GetProducts(ctx context.Context) []web.ProductResponse
	GetProduct(ctx context.Context, productId string) web.ProductResponse

	GetBrands(ctx context.Context) []web.BrandResponse
	GetUnits(ctx context.Context) []web.UnitResponse
}
