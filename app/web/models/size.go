package models

import (
	"time"
)

type Size struct {
	ID        string    `json:"id" gorm:"not null;uniqueIndex;primary_key"`
	StoreID   string    `json:"store_id" gorm:"not null;index"`
	Store     Store     `gorm:"foreignKey:StoreID" json:"store"`
	Name      string    `json:"name"`
	Value     string    `json:"value"`
	Products  []Product `gorm:"foreignKey:SizeID" json:"products"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}
