package models

import (
	"time"

	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	ID           string      `json:"id" gorm:"not null;uniqueIndex;primary_key"`
	StoreID      string      `json:"store_id" gorm:"not null;index"`
	Store        Store       `gorm:"foreignKey:StoreID" json:"store"`
	UserID       string      `json:"user_id" gorm:"not null"`
	OrderItems   []OrderItem `json:"order_items" gorm:"foreignKey:OrderID"`
	IsPaid       bool        `json:"is_paid"`
	CustomerName string      `json:"customer_name"`
	Phone        string      `json:"phone"`
	TotalPrice   float64     `json:"total_price"`
	Address      string      `json:"address"`
	CreatedAt    time.Time   `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt    time.Time   `json:"updated_at" gorm:"autoUpdateTime"`
}

type OrderResponse struct {
	ID           string      `json:"id"`
	StoreID      string      `json:"store_id"`
	Store        Store       `json:"store"`
	UserID       string      `json:"user_id"`
	OrderItems   []OrderItem `json:"order_items"`
	IsPaid       bool        `json:"is_paid"`
	CustomerName string      `json:"customer_name"`
	Phone        string      `json:"phone"`
	Address      string      `json:"address"`
	TotalPrice   float64     `json:"base_total_price"`
	CreatedAt    time.Time   `json:"created_at"`
	UpdatedAt    time.Time   `json:"updated_at"`
}

type OrderCreate struct {
	ID         string      `json:"id"`
	StoreID    string      `json:"store_id"`
	Store      Store       `json:"store"`
	OrderItems []OrderItem `json:"order_items"`
	UserID     string      `json:"user_id"`
	IsPaid     bool        `json:"is_paid"`
	Phone      string      `json:"phone"`
	Address    string      `json:"address"`
	TotalPrice float64     `json:"base_total_price"`
	CreatedAt  time.Time   `json:"created_at"`
	UpdatedAt  time.Time   `json:"updated_at"`
}

type OrderUpdate struct {
	ID         string      `json:"id"`
	StoreID    string      `json:"store_id"`
	Store      Store       `json:"store"`
	UserID     string      `json:"user_id"`
	OrderItems []OrderItem `json:"order_items"`
	IsPaid     bool        `json:"is_paid"`
	Phone      string      `json:"phone"`
	Address    string      `json:"address"`
	TotalPrice float64     `json:"base_total_price"`
	CreatedAt  time.Time   `json:"created_at"`
	UpdatedAt  time.Time   `json:"updated_at"`
}

func ToOrderResponse(order Order) OrderResponse {
	return OrderResponse{
		ID:           order.ID,
		StoreID:      order.StoreID,
		Store:        order.Store,
		UserID:       order.UserID,
		OrderItems:   order.OrderItems,
		IsPaid:       order.IsPaid,
		CustomerName: order.CustomerName,
		Phone:        order.Phone,
		Address:      order.Address,
		TotalPrice:   order.TotalPrice,
		CreatedAt:    order.CreatedAt,
		UpdatedAt:    order.UpdatedAt,
	}
}

func ToOrderResponses(orders []Order) []OrderResponse {
	var responses []OrderResponse

	for _, order := range orders {
		responses = append(responses, ToOrderResponse(order))
	}
	return responses
}
