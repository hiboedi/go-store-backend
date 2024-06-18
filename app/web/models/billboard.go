package models

import "time"

type Billboard struct {
	ID         string     `json:"id" gorm:"not null;uniqueIndex;primary_key"`
	StoreID    string     `json:"store_id" gorm:"not null;index"`
	Store      Store      `json:"store" gorm:"foreignKey:StoreID"`
	Label      string     `json:"label" gorm:"not null"`
	ImageURL   string     `json:"image_url" gorm:"not null"`
	Categories []Category `gorm:"foreignKey:BillboardID" json:"categories"`
	CreatedAt  time.Time  `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt  time.Time  `json:"updated_at" gorm:"autoUpdateTime"`
}
