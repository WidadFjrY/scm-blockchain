package web

import (
	"time"
)

type StoreCreateRequest struct {
	UserID  string `json:"user_id" validate:"required"`
	Name    string `json:"name" validate:"required"`
	Address string `json:"address" validate:"required"`
	Telp    string `json:"telp" validate:"required"`
}

type StoreResponse struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	ManagerName string    `json:"manager_name"`
	Address     string    `json:"address"`
	Telp        string    `json:"telp"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
