package models

import "time"

type Category struct {
	ID          string    `json:"id" gorm:"not null;uniqueIndex;primary_key"`
	StoreID     string    `json:"store_id" gorm:"not null;index"`
	Store       Store     `json:"-"`
	BillboardID string    `json:"billboard_id" gorm:"not null;index"`
	Billboard   Billboard `gorm:"foreignKey:BillboardID" json:"billboard"`
	Products    []Product `gorm:"foreignKey:CategoryID" json:"products"`
	Name        string    `json:"name"`
	CreatedAt   time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

type CategoryResponse struct {
	ID          string    `json:"id"`
	StoreID     string    `json:"store_id"`
	BillboardID string    `json:"billboard_id"`
	Billboard   Billboard `json:"billboard"`
	Name        string    `json:"name"`
	Products    []Product `json:"products"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type CategoryResponseHiddenStore struct {
	ID          string    `json:"id"`
	StoreID     string    `json:"store_id"`
	BillboardID string    `json:"billboard_id"`
	Name        string    `json:"name"`
	Products    []Product `json:"products"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type CategoryCreate struct {
	ID          string    `json:"id"`
	StoreID     string    `json:"store_id"`
	BillboardID string    `json:"billboard_id"`
	Name        string    `json:"name"`
	Products    []Product `json:"products"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type CategoryUpdate struct {
	ID          string    `json:"id"`
	StoreID     string    `json:"store_id"`
	BillboardID string    `json:"billboard_id"`
	Name        string    `json:"name"`
	Products    []Product `json:"products"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func ToCategoryResponse(category Category) CategoryResponse {
	return CategoryResponse{
		ID:          category.ID,
		StoreID:     category.StoreID,
		BillboardID: category.BillboardID,
		Billboard:   category.Billboard,
		Name:        category.Name,
		Products:    category.Products,
		CreatedAt:   category.CreatedAt,
		UpdatedAt:   category.UpdatedAt,
	}
}

func ToCategoryResponseHiddenStore(category Category) CategoryResponseHiddenStore {
	return CategoryResponseHiddenStore{
		ID:          category.ID,
		StoreID:     category.StoreID,
		BillboardID: category.BillboardID,
		Name:        category.Name,
		Products:    category.Products,
		CreatedAt:   category.CreatedAt,
		UpdatedAt:   category.UpdatedAt,
	}
}

func ToCategoryResponsesHiddenStore(categories []Category) []CategoryResponseHiddenStore {
	var responses []CategoryResponseHiddenStore
	for _, category := range categories {
		responses = append(responses, ToCategoryResponseHiddenStore(category))
	}
	return responses
}

func ToCategoryResponses(categories []Category) []CategoryResponse {
	var responses []CategoryResponse
	for _, category := range categories {
		responses = append(responses, ToCategoryResponse(category))
	}
	return responses
}
