package models

import (
	"time"
)

type OrderItem struct {
	ID        string    `json:"id" gorm:"not null;uniqueIndex;primary_key"`
	OrderID   string    `json:"order_id" gorm:"not null;index"`
	Order     Order     `gorm:"foreignKey:OrderID" json:"order"`
	ProductID string    `json:"product_id" gorm:"not null;index"`
	Product   Product   `gorm:"foreignKey:ProductID" json:"product"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}
