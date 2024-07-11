package models

import "time"

type Store struct {
	ID         string      `json:"id" gorm:"not null;uniqueIndex;primary_key"`
	Name       string      `json:"name" gorm:"not null;type:varchar(50)"`
	UserID     string      `json:"user_id" gorm:"not null;index"`
	User       User        `json:"user" gorm:"foreignKey:UserID"`
	Billboards []Billboard `json:"billboards"`
	Categories []Category  `gorm:"foreignKey:StoreID" json:"categories"`
	Sizes      []Size      `gorm:"foreignKey:StoreID" json:"sizes"`
	Colors     []Color     `gorm:"foreignKey:StoreID" json:"colors"`
	Products   []Product   `gorm:"foreignKey:StoreID" json:"products"`
	Orders     []Order     `gorm:"foreignKey:StoreID" json:"orders"`
	CreatedAt  time.Time   `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt  time.Time   `json:"updated_at" gorm:"autoUpdateTime"`
}

type StoreResponse struct {
	ID         string              `json:"id"`
	Name       string              `json:"name"`
	UserID     string              `json:"user_id"`
	User       UserResponse        `json:"user"`
	Billboards []BillboardResponse `json:"billboards"`
	Categories []CategoryResponse  `json:"categories"`
	Sizes      []SizeResponse      `json:"sizes"`
	Colors     []ColorResponse     `json:"colors"`
	Products   []ProductResponse   `json:"products"`
	Orders     []OrderResponse     `json:"orders"`
	CreatedAt  time.Time           `json:"created_at"`
	UpdatedAt  time.Time           `json:"updated_at"`
}

type StoreCreate struct {
	Name string `json:"name" `
}

type StoreUpdate struct {
	Name string `json:"name" `
}

func ToStoreResponse(store Store) StoreResponse {
	var billboards []BillboardResponse
	var categories []CategoryResponse
	var sizes []SizeResponse
	var colors []ColorResponse
	var products []ProductResponse
	var orders []OrderResponse
	user := ToUserReponse(store.User)

	for _, billboard := range store.Billboards {
		billboards = append(billboards, ToBillboardReponse(billboard))
	}
	for _, category := range store.Categories {
		categories = append(categories, ToCategoryResponse(category))
	}
	for _, size := range store.Sizes {
		sizes = append(sizes, ToSizeResponse(size))
	}
	for _, color := range store.Colors {
		colors = append(colors, ToColorResponse(color))
	}
	for _, product := range store.Products {
		products = append(products, ToProductResponse(product))
	}
	for _, order := range store.Orders {
		orders = append(orders, ToOrderResponse(order))
	}

	return StoreResponse{
		ID:         store.ID,
		Name:       store.Name,
		UserID:     store.UserID,
		User:       user,
		Billboards: billboards,
		Categories: categories,
		Sizes:      sizes,
		Colors:     colors,
		Products:   products,
		Orders:     orders,
		CreatedAt:  store.CreatedAt,
		UpdatedAt:  store.UpdatedAt,
	}
}

func ToStoreResponses(stores []Store) []StoreResponse {
	var responses []StoreResponse

	for _, store := range stores {
		responses = append(responses, ToStoreResponse(store))
	}
	return responses
}
