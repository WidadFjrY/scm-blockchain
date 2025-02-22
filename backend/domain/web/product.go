package web

import "time"

type ProductCreateRequest struct {
	DistributorID string  `json:"distributor_id" validate:"required"`
	ProductName   string  `json:"product_name" validate:"required"`
	Brand         string  `json:"brand" validate:"required"`
	Price         float32 `json:"price" validate:"required"`
	Unit          string  `json:"unit" validate:"required"`
	Stock         int     `json:"stock" validate:"required"`
	Status        string  `json:"status" validate:"required"`
}

type ProductCreateResponse struct {
	ProductName string    `json:"product_name"`
	CreatedAt   time.Time `json:"created_at"`
}

type ProductResponse struct {
	ID             string    `json:"product_id"`
	DistrbutorName string    `json:"distributor_name"`
	ProductName    string    `json:"product_name"`
	Brand          string    `json:"brand"`
	Price          string    `json:"price"`
	Stock          int       `json:"stock"`
	Unit           string    `json:"unit"`
	Status         string    `json:"status"`
	UrlPricture    string    `json:"url_picture"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}
