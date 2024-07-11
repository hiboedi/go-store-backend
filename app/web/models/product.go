package models

import (
	"time"
)

type Product struct {
	ID                 string      `json:"id" gorm:"not null;uniqueIndex;primary_key"`
	StoreID            string      `json:"store_id" gorm:"not null;index"`
	Store              Store       `gorm:"foreignKey:StoreID;references:ID" json:"-"`
	CategoryID         string      `json:"category_id" gorm:"not null;index"`
	Category           Category    `gorm:"foreignKey:CategoryID;references:ID" json:"category"`
	Name               string      `json:"name"`
	Price              float64     `json:"price" `
	Stock              int64       `json:"stock"`
	PriceAfterDiscount float64     `json:"price_discount"`
	DiscountPercent    uint32      `json:"discount_percent"`
	IsFeatured         bool        `json:"is_featured"`
	IsArchived         bool        `json:"is_archived"`
	SizeID             string      `json:"size_id" gorm:"not null;index"`
	Size               Size        `gorm:"foreignKey:SizeID" json:"size"`
	ColorID            string      `json:"color_id" gorm:"not null;index"`
	Color              Color       `gorm:"foreignKey:ColorID" json:"color"`
	Images             []Image     `gorm:"foreignKey:ProductID" json:"images"`
	CartItems          []CartItem  `gorm:"foreignKey:ProductID" json:"cart_items"`
	OrderItems         []OrderItem `gorm:"foreignKey:ProductID" json:"order_items"`
	CreatedAt          time.Time   `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt          time.Time   `json:"updated_at" gorm:"autoUpdateTime"`
}

type ProductResponse struct {
	ID                 string     `json:"id"`
	StoreID            string     `json:"store_id"`
	CategoryID         string     `json:"category_id"`
	Name               string     `json:"name"`
	Stock              int64      `json:"stock"`
	Price              float64    `json:"price"`
	PriceAfterDiscount float64    `json:"price_discount"`
	DiscountPercent    uint32     `json:"discount_percent"`
	IsFeatured         bool       `json:"is_featured"`
	IsArchived         bool       `json:"is_archived"`
	SizeID             string     `json:"size_id"`
	ColorID            string     `json:"color_id"`
	Images             []Image    `json:"images"`
	CartItems          []CartItem `json:"cart_items"`
	CreatedAt          time.Time  `json:"created_at"`
	UpdatedAt          time.Time  `json:"updated_at"`
}

type ProductCreate struct {
	CategoryID      string  `json:"category_id"`
	Name            string  `json:"name"`
	Price           float64 `json:"price"`
	Stock           int64   `json:"stock"`
	DiscountPercent uint32  `json:"discount_percent"`
	IsFeatured      bool    `json:"is_featured"`
	IsArchived      bool    `json:"is_archived"`
	SizeID          string  `json:"size_id"`
	ColorID         string  `json:"color_id"`
	Images          []Image `json:"images"`
}

type ProductUpdate struct {
	CategoryID      string  `json:"category_id"`
	Name            string  `json:"name"`
	Price           float64 `json:"price"`
	Stock           int64   `json:"stock"`
	DiscountPercent uint32  `json:"discount_percent"`
	IsFeatured      bool    `json:"is_featured"`
	IsArchived      bool    `json:"is_archived"`
	SizeID          string  `json:"size_id"`
	ColorID         string  `json:"color_id"`
	Images          []Image `json:"images"`
}

func ToProductResponse(product Product) ProductResponse {
	return ProductResponse{
		ID:                 product.ID,
		StoreID:            product.StoreID,
		CategoryID:         product.CategoryID,
		Name:               product.Name,
		Price:              product.Price,
		PriceAfterDiscount: product.PriceAfterDiscount,
		DiscountPercent:    product.DiscountPercent,
		Stock:              product.Stock,
		IsFeatured:         product.IsFeatured,
		IsArchived:         product.IsArchived,
		SizeID:             product.SizeID,
		ColorID:            product.ColorID,
		Images:             product.Images,
		CartItems:          product.CartItems,
		CreatedAt:          product.CreatedAt,
		UpdatedAt:          product.UpdatedAt,
	}
}

func ToProductResponses(products []Product, pagination Pagination) ProductPagination {
	var prodctResponse []ProductResponse
	paginationResponse := PaginationResponse{
		Total:       pagination.Total,
		CurrentPage: pagination.CurrentPage,
		PerPage:     pagination.PerPage,
		LastPage:    pagination.LastPage,
	}
	var responses ProductPagination

	for _, product := range products {
		prodctResponse = append(prodctResponse, ToProductResponse(product))
	}

	responses.Pagination = paginationResponse
	responses.Product = prodctResponse
	return responses
}

type ProductPagination struct {
	Pagination PaginationResponse
	Product    []ProductResponse
}
