package models

import (
	"time"
)

type Billboard struct {
	ID         string     `json:"id" gorm:"not null;uniqueIndex;primary_key"`
	StoreID    string     `json:"store_id" gorm:"not null;index"`
	Store      Store      `json:"-"`
	Label      string     `json:"label" gorm:"not null"`
	ImageURL   string     `json:"image_url" gorm:"not null"`
	Categories []Category `gorm:"foreignKey:BillboardID" json:"categories"`
	CreatedAt  time.Time  `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt  time.Time  `json:"updated_at" gorm:"autoUpdateTime"`
}

type BillboardResponse struct {
	ID         string     `json:"id"`
	StoreID    string     `json:"store_id"`
	Store      Store      `json:"store"`
	Label      string     `json:"label"`
	ImageURL   string     `json:"image_url"`
	Categories []Category `json:"categories"`
	CreatedAt  time.Time  `json:"created_at"`
	UpdatedAt  time.Time  `json:"updated_at"`
}

type BillboardResponseHiddenStore struct {
	ID         string     `json:"id"`
	StoreID    string     `json:"store_id"`
	Label      string     `json:"label"`
	ImageURL   string     `json:"image_url"`
	Categories []Category `json:"categories"`
	CreatedAt  time.Time  `json:"created_at"`
	UpdatedAt  time.Time  `json:"updated_at"`
}

type BillboardCreate struct {
	StoreID  string `json:"store_id"`
	Label    string `json:"label" validate:"required,min=4,max=50"`
	ImageURL string `json:"image_url" validate:"required"`
}

type BillboardUpdate struct {
	Label    string `json:"label" validate:"required,min=4,max=50"`
	ImageURL string `json:"image_url" validate:"required"`
}

func ToBillboardReponse(billboard Billboard) BillboardResponse {
	return BillboardResponse{
		ID:         billboard.ID,
		Label:      billboard.Label,
		StoreID:    billboard.StoreID,
		ImageURL:   billboard.ImageURL,
		Store:      billboard.Store,
		Categories: billboard.Categories,
		CreatedAt:  billboard.CreatedAt,
		UpdatedAt:  billboard.UpdatedAt,
	}
}

func ToBillboardResponseHiddenStore(billboard Billboard) BillboardResponseHiddenStore {
	return BillboardResponseHiddenStore{
		ID:         billboard.ID,
		StoreID:    billboard.StoreID,
		Label:      billboard.Label,
		ImageURL:   billboard.ImageURL,
		Categories: billboard.Categories,
		CreatedAt:  billboard.CreatedAt,
		UpdatedAt:  billboard.UpdatedAt,
	}
}
func ToBillboardResponsesHiddenStore(billboards []Billboard) []BillboardResponseHiddenStore {
	var responses []BillboardResponseHiddenStore

	for _, billboard := range billboards {
		responses = append(responses, ToBillboardResponseHiddenStore(billboard))
	}
	return responses
}

func ToBillboardResponses(billboards []Billboard) []BillboardResponse {
	var responses []BillboardResponse

	for _, billboard := range billboards {
		responses = append(responses, ToBillboardReponse(billboard))
	}
	return responses
}
