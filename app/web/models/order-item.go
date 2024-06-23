package models

import (
	"time"
)

type OrderItem struct {
	ID        string    `json:"id" gorm:"not null;uniqueIndex;primary_key"`
	OrderID   string    `json:"order_id" gorm:"not null;index"`
	Order     Order     `gorm:"foreignKey:OrderID" json:"order"`
	ProductID string    `json:"product_id" gorm:"not null;index"`
	Product   Product   `gorm:"foreignKey:ProductID" json:"product"`
	Quantity  int       `json:"quantity" gorm:"not null"`
	Price     float64   `json:"price" gorm:"not null"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

type OrderItemResponse struct {
	ID        string    `json:"id"`
	OrderID   string    `json:"order_id"`
	ProductID string    `json:"product_id"`
	Product   Product   `json:"product"`
	Quantity  int       `json:"quantity"`
	Price     float64   `json:"price"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type OrderItemCreate struct {
	OrderID   string    `json:"order_id"`
	ProductID string    `json:"product_id"`
	Quantity  int       `json:"quantity"`
	Price     float64   `json:"price"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type OrderItemUpdate struct {
	ID        string    `json:"id"`
	OrderID   string    `json:"order_id"`
	ProductID string    `json:"product_id"`
	Quantity  int       `json:"quantity"`
	Price     float64   `json:"price"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func ToOrderItemResponse(orderItem OrderItem) OrderItemResponse {
	return OrderItemResponse{
		ID:        orderItem.ID,
		OrderID:   orderItem.OrderID,
		ProductID: orderItem.ProductID,
		Product:   orderItem.Product,
		Quantity:  orderItem.Quantity,
		Price:     orderItem.Price,
		CreatedAt: orderItem.CreatedAt,
		UpdatedAt: orderItem.UpdatedAt,
	}
}

func ToOrderItemResponses(orderItems []OrderItem) []OrderItemResponse {
	var responses []OrderItemResponse
	for _, item := range orderItems {
		responses = append(responses, ToOrderItemResponse(item))
	}
	return responses
}
