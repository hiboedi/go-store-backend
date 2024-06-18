package models

import (
	"time"
)

type Product struct {
	ID         string      `json:"id" gorm:"not null;uniqueIndex;primary_key"`
	StoreID    string      `json:"store_id" gorm:"not null;index"`
	Store      Store       `gorm:"foreignKey:StoreID;references:ID" json:"store"`
	CategoryID string      `json:"category_id" gorm:"not null;index"`
	Category   Category    `gorm:"foreignKey:CategoryID;references:ID" json:"category"`
	Name       string      `json:"name"`
	Price      float64     `json:"price" gorm:"type:decimal(10,2)"`
	IsFeatured bool        `json:"is_featured"`
	IsArchived bool        `json:"is_archived"`
	SizeID     string      `json:"size_id" gorm:"not null;index"`
	Size       Size        `gorm:"foreignKey:SizeID" json:"size"`
	ColorID    string      `json:"color_id" gorm:"not null;index"`
	Color      Color       `gorm:"foreignKey:ColorID" json:"color"`
	Images     []Image     `gorm:"foreignKey:ProductID" json:"images"`
	OrderItems []OrderItem `gorm:"foreignKey:ProductID" json:"order_items"`
	CreatedAt  time.Time   `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt  time.Time   `json:"updated_at" gorm:"autoUpdateTime"`
}
