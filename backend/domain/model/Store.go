package model

import (
	"time"

	"gorm.io/gorm"
)

type Store struct {
	ID        string    `gorm:"primaryKey;type:char(15);not null"`
	UserID    string    `gorm:"type:char(15);not null"`
	User      User      `gorm:"foreignKey:UserID;references:ID"`
	Product   []Product `gorm:"foreignKey:StoreID;references:ID"`
	Name      string    `gorm:"type:varchar(100);not null"`
	Address   string    `gorm:"type:text;not null"`
	Telp      string    `gorm:"type:varchar(20);not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}
