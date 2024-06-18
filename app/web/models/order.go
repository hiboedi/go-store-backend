package models

import (
	"time"
)

type Order struct {
	ID         string      `json:"id" gorm:"not null;uniqueIndex;primary_key"`
	StoreID    string      `json:"store_id" gorm:"not null;index"`
	Store      Store       `gorm:"foreignKey:StoreID" json:"store"`
	OrderItems []OrderItem `gorm:"foreignKey:OrderID" json:"order_items"`
	IsPaid     bool        `json:"is_paid"`
	Phone      string      `json:"phone"`
	Address    string      `json:"address"`
	CreatedAt  time.Time   `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt  time.Time   `json:"updated_at" gorm:"autoUpdateTime"`
}
