package models

import "time"

type Cart struct {
	ID        string     `json:"id" gorm:"not null;uniqueIndex;primary_key"`
	UserID    string     `json:"user_id" gorm:"not null;index"`
	User      User       `gorm:"foreignKey:UserID;references:ID" json:"user"`
	OrderID   string     `json:"order_id" gorm:"index"`
	CartItems []CartItem `gorm:"foreignKey:CartID" json:"cart_items"`
	CreatedAt time.Time  `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time  `json:"updated_at" gorm:"autoUpdateTime"`
}

type CartResponse struct {
	ID        string     `json:"id" `
	UserID    string     `json:"user_id"`
	User      User       `json:"user"`
	CartItems []CartItem `json:"cart_items"`
	CreatedAt time.Time  `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time  `json:"updated_at" gorm:"autoUpdateTime"`
}

type CartCreate struct {
	ID        string     `json:"id" `
	UserID    string     `json:"user_id"`
	User      User       `json:"user"`
	CartItems []CartItem `json:"cart_items"`
}

type CartUpdate struct {
	ID        string     `json:"id" `
	UserID    string     `json:"user_id"`
	User      User       `json:"user"`
	CartItems []CartItem `json:"cart_items"`
}

func ToCartResponse(cart Cart) CartResponse {
	return CartResponse{
		ID:        cart.ID,
		UserID:    cart.UserID,
		CartItems: cart.CartItems,
		CreatedAt: cart.CreatedAt,
		UpdatedAt: cart.UpdatedAt,
	}
}
