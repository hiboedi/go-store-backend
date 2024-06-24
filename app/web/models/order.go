package models

import (
	"time"
)

type Order struct {
	ID              string      `json:"id" gorm:"not null;uniqueIndex;primary_key"`
	StoreID         string      `json:"store_id" gorm:"not null;index"`
	Store           Store       `gorm:"-"` // Jangan sertakan field Store dalam JSON
	OrderItems      []OrderItem `gorm:"foreignKey:OrderID" json:"order_items"`
	IsPaid          bool        `json:"is_paid"`
	Phone           string      `json:"phone"`
	BaseTotalPrice  float64     `json:"base_total_price"`
	TaxAmount       float64     `json:"tax_amount"`
	TaxPercent      float64     `json:"tax_percent"`
	DiscountAmount  float64     `json:"discount_amount"`
	DiscountPercent float64     `json:"discount_percent"`
	GrandTotaPrice  float64     `json:"grand_total_price"`
	Address         string      `json:"address"`
	CreatedAt       time.Time   `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt       time.Time   `json:"updated_at" gorm:"autoUpdateTime"`
}

type OrderResponse struct {
	ID              string      `json:"id"`
	StoreID         string      `json:"store_id"`
	Store           Store       `json:"store"`
	OrderItems      []OrderItem `json:"order_items"`
	IsPaid          bool        `json:"is_paid"`
	Phone           string      `json:"phone"`
	Address         string      `json:"address"`
	BaseTotalPrice  float64     `json:"base_total_price"`
	TaxAmount       float64     `json:"tax_amount"`
	TaxPercent      float64     `json:"tax_percent"`
	DiscountAmount  float64     `json:"discount_amount"`
	DiscountPercent float64     `json:"discount_percent"`
	GrandTotaPrice  float64     `json:"grand_total_price"`
	CreatedAt       time.Time   `json:"created_at"`
	UpdatedAt       time.Time   `json:"updated_at"`
}

type OrderResponseHiddenStore struct {
	ID              string      `json:"id"`
	StoreID         string      `json:"store_id"`
	OrderItems      []OrderItem `json:"order_items"`
	IsPaid          bool        `json:"is_paid"`
	Phone           string      `json:"phone"`
	Address         string      `json:"address"`
	BaseTotalPrice  float64     `json:"base_total_price"`
	TaxAmount       float64     `json:"tax_amount"`
	TaxPercent      float64     `json:"tax_percent"`
	DiscountAmount  float64     `json:"discount_amount"`
	DiscountPercent float64     `json:"discount_percent"`
	GrandTotaPrice  float64     `json:"grand_total_price"`
	CreatedAt       time.Time   `json:"created_at"`
	UpdatedAt       time.Time   `json:"updated_at"`
}

type OrderCreate struct {
	ID              string      `json:"id"`
	StoreID         string      `json:"store_id"`
	OrderItems      []OrderItem `json:"order_items"`
	IsPaid          bool        `json:"is_paid"`
	Phone           string      `json:"phone" validate:"required"`
	Address         string      `json:"address" validate:"required"`
	BaseTotalPrice  float64     `json:"base_total_price"`
	TaxAmount       float64     `json:"tax_amount"`
	TaxPercent      float64     `json:"tax_percent"`
	DiscountAmount  float64     `json:"discount_amount"`
	DiscountPercent float64     `json:"discount_percent"`
	GrandTotaPrice  float64     `json:"grand_total_price"`
}

type OrderUpdate struct {
	IsPaid bool `json:"is_paid"`
}

func ToOrderResponse(order Order) OrderResponse {
	return OrderResponse{
		ID:              order.ID,
		StoreID:         order.StoreID,
		Store:           order.Store,
		OrderItems:      order.OrderItems,
		IsPaid:          order.IsPaid,
		Phone:           order.Phone,
		Address:         order.Address,
		BaseTotalPrice:  order.BaseTotalPrice,
		DiscountAmount:  order.DiscountAmount,
		DiscountPercent: order.DiscountPercent,
		TaxAmount:       order.TaxAmount,
		TaxPercent:      order.TaxPercent,
		CreatedAt:       order.CreatedAt,
		UpdatedAt:       order.UpdatedAt,
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
