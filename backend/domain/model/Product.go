package model

import (
	"time"
)

type Product struct {
	ID            string      `gorm:"primaryKey;type:char(15);not null"`
	DistributorID string      `gorm:"type:char(15);not null"`
	StoreID       *string     `gorm:"type:char(15)"`
	ProductName   string      `gorm:"type:varchar(100);not null"`
	Store         *Store      `gorm:"foreignKey:StoreID;references:ID"`
	Distributor   Distributor `gorm:"foreignKey:DistributorID;references:ID"`
	Brand         string      `gorm:"type:varchar(100);not null"`
	Price         float32     `gorm:"type:float;not null;default:0"`
	Stock         int         `gorm:"type:integer;not null;default:0"`
	Unit          string      `gorm:"type:varchar(100);not null;"`
	UrlPricture   string      `gorm:"type:varchar(100);not null;default:default.jpg"`
	Status        string      `gorm:"type:varchar(100);not null;"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
}
