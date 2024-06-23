package models

import (
	"time"
)

type Order struct {
	ID         string      `json:"id" gorm:"not null;uniqueIndex;primary_key"`
	StoreID    string      `json:"store_id" gorm:"not null;index"`
	Store      Store       `gorm:"-"` // Jangan sertakan field Store dalam JSON
	OrderItems []OrderItem `gorm:"foreignKey:OrderID" json:"order_items"`
	IsPaid     bool        `json:"is_paid"`
	Phone      string      `json:"phone"`
	Address    string      `json:"address"`
	CreatedAt  time.Time   `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt  time.Time   `json:"updated_at" gorm:"autoUpdateTime"`
}

type OrderResponse struct {
	ID         string      `json:"id"`
	StoreID    string      `json:"store_id"`
	Store      Store       `json:"store"`
	OrderItems []OrderItem `json:"order_items"`
	IsPaid     bool        `json:"is_paid"`
	Phone      string      `json:"phone"`
	Address    string      `json:"address"`
	CreatedAt  time.Time   `json:"created_at"`
	UpdatedAt  time.Time   `json:"updated_at"`
}

type OrderResponseHiddenStore struct {
	ID         string      `json:"id"`
	StoreID    string      `json:"store_id"`
	OrderItems []OrderItem `json:"order_items"`
	IsPaid     bool        `json:"is_paid"`
	Phone      string      `json:"phone"`
	Address    string      `json:"address"`
	CreatedAt  time.Time   `json:"created_at"`
	UpdatedAt  time.Time   `json:"updated_at"`
}

type OrderCreate struct {
	ID         string      `json:"id"`
	StoreID    string      `json:"store_id"`
	OrderItems []OrderItem `json:"order_items"`
	IsPaid     bool        `json:"is_paid"`
	Phone      string      `json:"phone"`
	Address    string      `json:"address"`
}

type OrderUpdate struct {
	ID         string      `json:"id"`
	StoreID    string      `json:"store_id"`
	OrderItems []OrderItem `json:"order_items"`
	IsPaid     bool        `json:"is_paid"`
	Phone      string      `json:"phone"`
	Address    string      `json:"address"`
}

func ToOrderResponse(order Order) OrderResponse {
	return OrderResponse{
		ID:         order.ID,
		StoreID:    order.StoreID,
		Store:      order.Store,
		OrderItems: order.OrderItems,
		IsPaid:     order.IsPaid,
		Phone:      order.Phone,
		Address:    order.Address,
		CreatedAt:  order.CreatedAt,
		UpdatedAt:  order.UpdatedAt,
	}
}

func ToOrderResponseHiddenStore(order Order) OrderResponseHiddenStore {
	return OrderResponseHiddenStore{
		ID:         order.ID,
		StoreID:    order.StoreID,
		OrderItems: order.OrderItems,
		IsPaid:     order.IsPaid,
		Phone:      order.Phone,
		Address:    order.Address,
		CreatedAt:  order.CreatedAt,
		UpdatedAt:  order.UpdatedAt,
	}
}

func ToOrderResponses(orders []Order) []OrderResponse {
	var responses []OrderResponse

	for _, order := range orders {
		responses = append(responses, ToOrderResponse(order))
	}
	return responses
}

func ToOrderResponsesHiddenStore(orders []Order) []OrderResponseHiddenStore {
	var responses []OrderResponseHiddenStore

	for _, order := range orders {
		responses = append(responses, ToOrderResponseHiddenStore(order))
	}
	return responses
}
