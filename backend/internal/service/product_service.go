package service

import (
	"context"
	"scm-blockchain-ethereum/domain/web"
)

type ProductService interface {
	ProductCreate(ctx context.Context, request web.ProductCreateRequest, filePath string) web.ProductCreateResponse
	BrandCreate(ctx context.Context, request web.BrandCreateRequest) web.BrandCreateResponse
	UnitCreate(ctx context.Context, request web.UnitCreateRequest) web.UnitCreateResponse
}
