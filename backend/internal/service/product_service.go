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
	GetTotalProduct(ctx context.Context) web.GetTotalProductResponse

	DeleteProductById(ctx context.Context, productId string)

	GetBrands(ctx context.Context) []web.BrandResponse
	GetUnits(ctx context.Context) []web.UnitResponse

	AddToCart(ctx context.Context, request web.AddToCartRequest, userId string)
	GetCarts(ctx context.Context, userId string) []web.CartResponse
	UpdateCartQty(ctx context.Context, productId string, qty int)
	DeleteItemCart(ctx context.Context, productId string)
}
