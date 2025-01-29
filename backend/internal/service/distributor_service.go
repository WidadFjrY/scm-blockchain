package service

import (
	"context"
	"scm-blockchain-ethereum/domain/web"
)

type DistributorService interface {
	Create(ctx context.Context, request web.DistributorCreateRequest, userId string) web.DistributorCreateResponse
	GetAll(ctx context.Context) []web.DistributorResponse
	GetByID(ctx context.Context, distributorId string) web.DistributorResponse
}
