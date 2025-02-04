package service

import (
	"context"
	"scm-blockchain-ethereum/domain/web"
)

type StoreService interface {
	Create(ctx context.Context, role string, request web.StoreCreateRequest)
	GetAll(ctx context.Context) []web.StoreResponse
	GetById(ctx context.Context, stroreId string) web.StoreResponse
}
