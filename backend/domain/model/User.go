package model

import (
	"time"
)

type User struct {
	ID            string  `gorm:"primaryKey;type:char(15);not null"`
	DistributorID *string `gorm:"type:char(15)"`
	StoreID       *string `gorm:"type:char(15)"`
	Name          string  `gorm:"type:varchar(100);not null"`
	Email         string  `gorm:"type:varchar(100);not null;unique"`
	Telp          string  `gorm:"type:varchar(100);not null"`
	Password      string  `gorm:"type:varchar(100);not null"`
	Role          string  `gorm:"type:varchar(100);not null;default:Admin"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
}
