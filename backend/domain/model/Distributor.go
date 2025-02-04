package model

import (
	"time"

	"gorm.io/gorm"
)

type Distributor struct {
	ID        string    `gorm:"primaryKey;type:char(15);not null"`
	Name      string    `gorm:"type:char(100);not null"`
	Address   string    `gorm:"type:char(100); not null"`
	Telp      string    `gorm:"type:char(100); not null"`
	Products  []Product `gorm:"foreignKey:DistributorID"`
	User      User      `gorm:"foreignKey:DistributorID"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}
