package models

import (
	"time"
)

type Color struct {
	ID        string    `json:"id" gorm:"not null;uniqueIndex;primary_key"`
	StoreID   string    `json:"store_id" gorm:"not null;index"`
	Store     Store     `gorm:"foreignKey:StoreID" json:"-"`
	Name      string    `json:"name"`
	Value     string    `json:"value"`
	Products  []Product `gorm:"foreignKey:ColorID" json:"products"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

type ColorResponse struct {
	ID        string    `json:"id"`
	StoreID   string    `json:"store_id"`
	Name      string    `json:"name"`
	Value     string    `json:"value"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type ColorResponseHiddenStore struct {
	ID        string    `json:"id"`
	StoreID   string    `json:"store_id"`
	Name      string    `json:"name"`
	Value     string    `json:"value"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type ColorCreate struct {
	ID        string    `json:"id"`
	StoreID   string    `json:"store_id"`
	Name      string    `json:"name"`
	Value     string    `json:"value"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type ColorUpdate struct {
	ID        string    `json:"id"`
	StoreID   string    `json:"store_id"`
	Name      string    `json:"name"`
	Value     string    `json:"value"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func ToColorResponse(color Color) ColorResponse {
	return ColorResponse{
		ID:        color.ID,
		StoreID:   color.StoreID,
		Name:      color.Name,
		Value:     color.Value,
		CreatedAt: color.CreatedAt,
		UpdatedAt: color.UpdatedAt,
	}
}

func ToColorResponseHiddenStore(color Color) ColorResponseHiddenStore {
	return ColorResponseHiddenStore{
		ID:        color.ID,
		StoreID:   color.StoreID,
		Name:      color.Name,
		Value:     color.Value,
		CreatedAt: color.CreatedAt,
		UpdatedAt: color.UpdatedAt,
	}
}

func ToColorResponses(colors []Color) []ColorResponse {
	var responses []ColorResponse

	for _, color := range colors {
		responses = append(responses, ToColorResponse(color))
	}

	return responses
}

func ToColorResponsesHiddenStore(colors []Color) []ColorResponseHiddenStore {
	var responses []ColorResponseHiddenStore

	for _, color := range colors {
		responses = append(responses, ToColorResponseHiddenStore(color))
	}

	return responses
}
