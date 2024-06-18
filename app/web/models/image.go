package models

import (
	"time"
)

type Image struct {
	ID        string    `json:"id" gorm:"not null;uniqueIndex;primary_key"`
	ProductID string    `json:"product_id" gorm:"not null;index"`
	Product   Product   `gorm:"foreignKey:ProductID;constraint:OnDelete:CASCADE" json:"product"`
	URL       string    `json:"url"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}
