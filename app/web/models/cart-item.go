package models

import (
	"time"
)

type CartItem struct {
	ID              string    `json:"id" gorm:"not null;uniqueIndex;primary_key"`
	CartID          string    `json:"cart_id" gorm:"not null;index"`
	Cart            Cart      `gorm:"foreignKey:CartID" json:"cart"`
	ProductID       string    `json:"product_id" gorm:"not null;index"`
	Product         Product   `gorm:"foreignKey:ProductID" json:"product"`
	Quantity        uint32    `json:"quantity" gorm:"not null"`
	TaxAmount       float64   `json:"tax_amount"`
	TaxPercent      uint32    `json:"tax_percent"`
	DiscountAmount  float64   `json:"discount_amount"`
	DiscountPercent uint32    `json:"discount_percent"`
	BasePrice       float64   `json:"base_price"`
	TotalPrice      float64   `json:"total_price"`
	CreatedAt       time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt       time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

type CartItemResponse struct {
	ID              string    `json:"id"`
	CartID          string    `json:"cart_id"`
	Cart            Cart      `json:"cart"`
	ProductID       string    `json:"product_id"`
	Product         Product   `json:"product"`
	Quantity        uint32    `json:"quantity"`
	BasePrice       float64   `json:"base_price"`
	DiscountAmount  float64   `json:"discount_amount"`
	DiscountPercent uint32    `json:"discount_percent"`
	TaxAmount       float64   `json:"tax_amount"`
	TaxPercent      uint32    `json:"tax_percent"`
	TotaPrice       float64   `json:"grand_total_price"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

type CartItemCreate struct {
	ProductID string `json:"product_id"`
	Quantity  uint32 `json:"quantity"`
}

type CartItemUpdate struct {
	ProductID string `json:"product_id"`
	Quantity  uint32 `json:"quantity"`
}

func ToCartItemResponse(cartItem CartItem) CartItemResponse {
	return CartItemResponse{
		ID:              cartItem.ID,
		CartID:          cartItem.CartID,
		Cart:            cartItem.Cart,
		ProductID:       cartItem.ProductID,
		Product:         cartItem.Product,
		Quantity:        cartItem.Quantity,
		DiscountAmount:  cartItem.DiscountAmount,
		DiscountPercent: cartItem.DiscountPercent,
		TaxAmount:       cartItem.TaxAmount,
		TaxPercent:      cartItem.TaxPercent,
		BasePrice:       cartItem.BasePrice,
		TotaPrice:       cartItem.TotalPrice,
		CreatedAt:       cartItem.CreatedAt,
		UpdatedAt:       cartItem.UpdatedAt,
	}
}

func ToCartItemResponses(cartItems []CartItem) []CartItemResponse {
	var responses []CartItemResponse
	for _, item := range cartItems {
		responses = append(responses, ToCartItemResponse(item))
	}
	return responses
}
