package model

import "time"

type NormalizedProduct struct {
	Product      Product
	ProductPrice ProductPrice
	ProductStock ProductStock
}

type Product struct {
	ID          string `gorm:"primaryKey;type:char(15);not null"`
	ProductName string `gorm:"type:varchar(100);not null"`
	Description string `gorm:"type:text;default:null"`
	PicturePath string `gorm:"type:varchar(100);default:default.jpg"`
	BrandID     string `gorm:"not null"`
	UnitID      string `gorm:"not null"`
	CreatedAt   time.Time
	UpdatedAt   time.Time

	Brand ProductBrand `gorm:"foreignKey:BrandID;references:ID"`
	Unit  ProductUnit  `gorm:"foreignKey:UnitID;references:ID"`

	Prices []ProductPrice `gorm:"foreignKey:ProductID"`
	Stocks []ProductStock `gorm:"foreignKey:ProductID"`
}

type ProductBrand struct {
	ID   string `gorm:"primaryKey;type:char(15);not null"`
	Name string `gorm:"type:varchar(100);not null;unique"`
}

type ProductUnit struct {
	ID   string `gorm:"primaryKey;type:char(15);not null"`
	Name string `gorm:"type:varchar(50);not null;unique"`
}

type ProductPrice struct {
	ID        string    `gorm:"primaryKey;type:char(15);not null"`
	ProductID string    `gorm:"type:char(15);not null"`
	Price     float32   `gorm:"type:float;not null;default:0"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`

	Product Product `gorm:"foreignKey:ProductID;references:ID"`
}

type ProductStock struct {
	ID         string    `gorm:"primaryKey;type:char(15);not null"`
	ProductID  string    `gorm:"type:char(15);not null"`
	StockIn    int       `gorm:"type:integer;default:0"`
	StockOut   int       `gorm:"type:integer;default:0"`
	StockTotal int       `gorm:"type:integer;not null"`
	UpdatedAt  time.Time `gorm:"autoUpdateTime"`

	Product Product `gorm:"foreignKey:ProductID;references:ID"`
}
