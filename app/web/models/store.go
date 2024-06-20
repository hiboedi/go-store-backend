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
	ID         string                         `json:"id"`
	Name       string                         `json:"name"`
	UserID     string                         `json:"user_id"`
	User       User                           `json:"user"`
	Billboards []BillboardResponseHiddenStore `json:"billboards"`
	Categories []Category                     `json:"categories"`
	Sizes      []Size                         `json:"sizes"`
	Colors     []Color                        `json:"colors"`
	Products   []Product                      `json:"products"`
	Orders     []Order                        `json:"orders"`
	CreatedAt  time.Time                      `json:"created_at"`
	UpdatedAt  time.Time                      `json:"updated_at"`
}

type StoreCreate struct {
	ID         string      `json:"id" `
	Name       string      `json:"name" `
	UserID     string      `json:"user_id" `
	User       User        `json:"user" `
	Billboards []Billboard `json:"billboards"`
	Categories []Category  `json:"categories"`
	Sizes      []Size      `json:"sizes"`
	Colors     []Color     `json:"colors"`
	Products   []Product   `json:"products"`
	Orders     []Order     `json:"orders"`
	CreatedAt  time.Time   `json:"created_at" `
	UpdatedAt  time.Time   `json:"updated_at" `
}

type StoreUpdate struct {
	ID         string      `json:"id" `
	Name       string      `json:"name" `
	UserID     string      `json:"user_id" `
	User       User        `json:"user" `
	Billboards []Billboard `json:"billboards"`
	Categories []Category  `json:"categories"`
	Sizes      []Size      `json:"sizes"`
	Colors     []Color     `json:"colors"`
	Products   []Product   `json:"products"`
	Orders     []Order     `json:"orders"`
	CreatedAt  time.Time   `json:"created_at" `
	UpdatedAt  time.Time   `json:"updated_at" `
}

func ToStoreResponse(store Store) StoreResponse {
	var billboards []BillboardResponseHiddenStore
	for _, billboard := range store.Billboards {
		billboards = append(billboards, ToBillboardResponseHiddenStore(billboard))
	}

	return StoreResponse{
		ID:         store.ID,
		Name:       store.Name,
		UserID:     store.UserID,
		User:       store.User,
		Billboards: billboards,
		Categories: store.Categories,
		Sizes:      store.Sizes,
		Colors:     store.Colors,
		Products:   store.Products,
		Orders:     store.Orders,
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
