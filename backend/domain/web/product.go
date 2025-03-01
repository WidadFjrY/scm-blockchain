package web

import (
	"time"
)

type ProductResponse struct {
	ID          string    `json:"id"`
	ProductName string    `json:"product_name"`
	Stock       int       `json:"stock"`
	Unit        string    `json:"unit"`
	Brand       string    `json:"brand"`
	FilePath    string    `json:"filepath"`
	Price       float32   `json:"price"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
}

type BrandResponse struct {
	ID        string `json:"id"`
	BrandName string `json:"brand_name"`
}

type UnitResponse struct {
	ID       string `json:"id"`
	UnitName string `json:"unit_name"`
}

type ProductCreateRequest struct {
	ProductName string  `json:"product_name" validate:"required"`
	BrandId     string  `json:"brand_id" validate:"required"`
	UnitId      string  `json:"unit_id" validate:"required"`
	Description string  `json:"description" validate:"omitempty"`
	Price       float32 `json:"price" validate:"required"`
	Stock       int     `json:"stock" validate:"required"`
}

type ProductCreateResponse struct {
	ProductId string    `json:"product_id"`
	CreatedAt time.Time `json:"created_at"`
}

type BrandCreateRequest struct {
	BrandName string `json:"brand_name" validate:"required"`
}

type BrandCreateResponse struct {
	BrandName string    `json:"brand_name"`
	CreatedAt time.Time `json:"created_at"`
}

type UnitCreateRequest struct {
	UnitName string `json:"unit_name" validate:"required"`
}

type UnitCreateResponse struct {
	UnitName  string    `json:"unit_name"`
	CreatedAt time.Time `json:"created_at"`
}
