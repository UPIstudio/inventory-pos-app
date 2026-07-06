package models

import (
	"time"
)

type Product struct {
	ID         uint       `json:"id" gorm:"primarykey"`
	SKU        string     `json:"sku" gorm:"unique;not null"`
	Name       string     `json:"name" gorm:"not null"`
	Price      float64    `json:"price" gorm:"type:decimal(12,2);not null"`
	CostPrice  float64    `json:"costprice" gorm:"type:decimal(12,2);not null"`
	Stock      int        `json:"stock" gorm:"default:0"`
	ExpiryDate *time.Time `json:"expiry_date"`
	CreatedAt  time.Time  `json:"created_at"`
	UpdatedAt  time.Time  `json:"updated_at"`
}
