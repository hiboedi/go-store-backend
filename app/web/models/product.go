package models

import (
	"time"
)

type Product struct {
	ID         string      `json:"id" gorm:"not null;uniqueIndex;primary_key"`
	StoreID    string      `json:"store_id" gorm:"not null;index"`
	Store      Store       `gorm:"foreignKey:StoreID;references:ID" json:"-"`
	CategoryID string      `json:"category_id" gorm:"not null;index"`
	Category   Category    `gorm:"foreignKey:CategoryID;references:ID" json:"category"`
	Name       string      `json:"name"`
	Price      float64     `json:"price" `
	Stock      int64       `json:"stock"`
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

type ProductResponse struct {
	ID         string      `json:"id"`
	StoreID    string      `json:"store_id"`
	Store      Store       `json:"store"`
	CategoryID string      `json:"category_id"`
	Category   Category    `json:"category"`
	Name       string      `json:"name"`
	Stock      int64       `json:"stock"`
	Price      float64     `json:"price"`
	IsFeatured bool        `json:"is_featured"`
	IsArchived bool        `json:"is_archived"`
	SizeID     string      `json:"size_id"`
	Size       Size        `json:"size"`
	ColorID    string      `json:"color_id"`
	Color      Color       `json:"color"`
	Images     []Image     `json:"images"`
	OrderItems []OrderItem `json:"order_items"`
	CreatedAt  time.Time   `json:"created_at"`
	UpdatedAt  time.Time   `json:"updated_at"`
}

type ProductResponseHiddenStore struct {
	ID         string      `json:"id"`
	StoreID    string      `json:"store_id"`
	CategoryID string      `json:"category_id"`
	Category   Category    `json:"category"`
	Name       string      `json:"name"`
	Stock      int64       `json:"stock"`
	Price      float64     `json:"price"`
	IsFeatured bool        `json:"is_featured"`
	IsArchived bool        `json:"is_archived"`
	SizeID     string      `json:"size_id"`
	Size       Size        `json:"size"`
	ColorID    string      `json:"color_id"`
	Color      Color       `json:"color"`
	Images     []Image     `json:"images"`
	OrderItems []OrderItem `json:"order_items"`
	CreatedAt  time.Time   `json:"created_at"`
	UpdatedAt  time.Time   `json:"updated_at"`
}

type ProductCreate struct {
	ID         string      `json:"id"`
	StoreID    string      `json:"store_id"`
	Store      Store       `json:"store"`
	CategoryID string      `json:"category_id"`
	Category   Category    `json:"category"`
	Name       string      `json:"name"`
	Price      float64     `json:"price"`
	Stock      int64       `json:"stock"`
	IsFeatured bool        `json:"is_featured"`
	IsArchived bool        `json:"is_archived"`
	SizeID     string      `json:"size_id"`
	Size       Size        `json:"size"`
	ColorID    string      `json:"color_id"`
	Color      Color       `json:"color"`
	Images     []Image     `json:"images"`
	OrderItems []OrderItem `json:"order_items"`
	CreatedAt  time.Time   `json:"created_at"`
	UpdatedAt  time.Time   `json:"updated_at"`
}

type ProductUpdate struct {
	StoreID    string      `json:"store_id"`
	Store      Store       `json:"store"`
	CategoryID string      `json:"category_id"`
	Category   Category    `json:"category"`
	Name       string      `json:"name"`
	Price      float64     `json:"price"`
	Stock      int64       `json:"stock"`
	IsFeatured bool        `json:"is_featured"`
	IsArchived bool        `json:"is_archived"`
	SizeID     string      `json:"size_id"`
	Size       Size        `json:"size"`
	ColorID    string      `json:"color_id"`
	Color      Color       `json:"color"`
	Images     []Image     `json:"images"`
	OrderItems []OrderItem `json:"order_items"`
	CreatedAt  time.Time   `json:"created_at"`
	UpdatedAt  time.Time   `json:"updated_at"`
}

func ToProductResponse(product Product) ProductResponse {
	return ProductResponse{
		ID:         product.ID,
		StoreID:    product.StoreID,
		Store:      product.Store,
		CategoryID: product.CategoryID,
		Category:   product.Category,
		Name:       product.Name,
		Price:      product.Price,
		Stock:      product.Stock,
		IsFeatured: product.IsFeatured,
		IsArchived: product.IsArchived,
		SizeID:     product.SizeID,
		Size:       product.Size,
		ColorID:    product.ColorID,
		Color:      product.Color,
		Images:     product.Images,
		OrderItems: product.OrderItems,
		CreatedAt:  product.CreatedAt,
		UpdatedAt:  product.UpdatedAt,
	}
}

func ToProductResponseHiddenStore(product Product) ProductResponseHiddenStore {
	return ProductResponseHiddenStore{
		ID:         product.ID,
		StoreID:    product.StoreID,
		CategoryID: product.CategoryID,
		Category:   product.Category,
		Name:       product.Name,
		Price:      product.Price,
		IsFeatured: product.IsFeatured,
		IsArchived: product.IsArchived,
		SizeID:     product.SizeID,
		Size:       product.Size,
		Stock:      product.Stock,
		ColorID:    product.ColorID,
		Color:      product.Color,
		Images:     product.Images,
		OrderItems: product.OrderItems,
		CreatedAt:  product.CreatedAt,
		UpdatedAt:  product.UpdatedAt,
	}
}

func ToProductResponsesHiddenStore(products []Product) []ProductResponseHiddenStore {
	var responses []ProductResponseHiddenStore

	for _, product := range products {
		responses = append(responses, ToProductResponseHiddenStore(product))
	}
	return responses
}

func ToProductResponses(products []Product) []ProductResponse {
	var responses []ProductResponse

	for _, product := range products {
		responses = append(responses, ToProductResponse(product))
	}
	return responses
}
