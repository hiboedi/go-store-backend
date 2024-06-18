package models

import "time"

// type Store struct {
// 	ID     string `json:"id" gorm:"not null;uniqueIndex;primary_key"`
// 	Name   string `json:"name"`
// 	UserID string `json:"user_id"`
// 	User   User   `json:"user" gorm:"foreignKey:UserID"`

// 	CreatedAt time.Time `json:"created_at"`
// 	UpdatedAt time.Time `json:"updated_at"`
// }

type Store struct {
	ID         string      `json:"id" gorm:"not null;uniqueIndex;primary_key"`
	Name       string      `json:"name" gorm:"not null;type:varchar(50)"`
	UserID     string      `json:"user_id" gorm:"not null;index"`
	User       User        `json:"user" gorm:"foreignKey:UserID"`
	Billboards []Billboard `json:"billboards"`
	Categories []Category  `gorm:"foreignKey:StoreID" json:"categories"`
	Sizes      []Size      `gorm:"foreignKey:StoreID" json:"sizes"`
	Colors     []Color     `gorm:"foreignKey:StoreID" json:"colors"`
	Products   []Product   `gorm:"foreignKey:StoreID" json:"products"`
	Orders     []Order     `gorm:"foreignKey:StoreID" json:"orders"`
	CreatedAt  time.Time   `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt  time.Time   `json:"updated_at" gorm:"autoUpdateTime"`
}
