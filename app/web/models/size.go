package models

import (
	"time"
)

type Size struct {
	ID        string    `json:"id" gorm:"not null;uniqueIndex;primary_key"`
	StoreID   string    `json:"store_id" gorm:"not null;index"`
	Store     Store     `gorm:"foreignKey:StoreID" json:"-"`
	Name      string    `json:"name"`
	Value     string    `json:"value"`
	Products  []Product `gorm:"foreignKey:SizeID" json:"products"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

type SizeResponse struct {
	ID        string    `json:"id"`
	StoreID   string    `json:"store_id"`
	Name      string    `json:"name"`
	Value     string    `json:"value"`
	Products  []Product `json:"products"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type SizeCreate struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type SizeUpdate struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

func ToSizeResponse(size Size) SizeResponse {
	return SizeResponse{
		ID:        size.ID,
		StoreID:   size.StoreID,
		Name:      size.Name,
		Value:     size.Value,
		Products:  size.Products,
		CreatedAt: size.CreatedAt,
		UpdatedAt: size.UpdatedAt,
	}
}

func ToSizeResponses(sizes []Size) []SizeResponse {
	var responses []SizeResponse

	for _, size := range sizes {
		responses = append(responses, ToSizeResponse(size))
	}
	return responses
}
