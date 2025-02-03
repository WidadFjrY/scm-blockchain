package service

import (
	"context"
	"scm-blockchain-ethereum/domain/web"
)

type ProductService interface {
	Create(ctx context.Context, request web.ProductCreateRequest, filePath string, role string) web.ProductCreateResponse
	GetAll(ctx context.Context) []web.ProductResponse
	GetById(ctx context.Context, productId string) web.ProductResponse
}
