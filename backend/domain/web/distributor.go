package web

import "time"

type DistributorCreateRequest struct {
	Name    string `json:"name" validate:"required,min=1"`
	Address string `json:"address" validate:"required,min=1"`
}

type DistributorCreateResponse struct {
	Name      string    `json:"name"`
	Address   string    `json:"address"`
	CreatedAt time.Time `json:"created_at"`
}

type DistributorResponse struct {
	ID             string    `json:"id"`
	Name           string    `json:"name"`
	PersonInCharge string    `json:"person_in_charge"`
	Address        string    `json:"address"`
	Telp           string    `json:"telp"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}
