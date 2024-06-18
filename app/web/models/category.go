package models

import "time"

type Category struct {
	ID          string    `json:"id" gorm:"not null;uniqueIndex;primary_key"`
	StoreID     string    `json:"store_id" gorm:"not null;index"`
	Store       Store     `gorm:"foreignKey:StoreID" json:"store"`
	BillboardID string    `json:"billboard_id" gorm:"not null;index"`
	Billboard   Billboard `gorm:"foreignKey:BillboardID" json:"billboard"`
	Products    []Product `gorm:"foreignKey:CategoryID" json:"products"`
	Name        string    `json:"name"`
	CreatedAt   time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}
