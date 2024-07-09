package repositories

import (
	"context"

	"github.com/hiboedi/go-store-backend/app/helpers"
	"github.com/hiboedi/go-store-backend/app/web/models"
	"gorm.io/gorm"
)

type CartRepository interface {
	GetCartByUserID(ctx context.Context, db *gorm.DB, userID string) (models.Cart, error)
	CreateCart(ctx context.Context, db *gorm.DB, cart models.Cart) (models.Cart, error)
	AddCartItem(ctx context.Context, db *gorm.DB, cartItem models.CartItem) (models.CartItem, error)
	UpdateCart(ctx context.Context, db *gorm.DB, cart models.Cart) (models.Cart, error)
	DeleteCart(ctx context.Context, db *gorm.DB, cart models.Cart) error
}

type cartRepositoryImpl struct {
}

func NewCartRepository() CartRepository {
	return &cartRepositoryImpl{}
}

func (r *cartRepositoryImpl) GetCartByUserID(ctx context.Context, db *gorm.DB, userID string) (models.Cart, error) {
	var cart models.Cart
	if err := db.WithContext(ctx).Where("user_id = ? AND order_id IS NULL", userID).First(&cart).Error; err != nil {
		helpers.PanicIfError(err)
	}
	return cart, nil
}

func (r *cartRepositoryImpl) CreateCart(ctx context.Context, db *gorm.DB, cart models.Cart) (models.Cart, error) {
	err := db.WithContext(ctx).Create(&cart).Error
	helpers.PanicIfError(err)

	return cart, nil
}

func (r *cartRepositoryImpl) AddCartItem(ctx context.Context, db *gorm.DB, cartItem models.CartItem) (models.CartItem, error) {
	err := db.WithContext(ctx).Create(&cartItem).Error
	helpers.PanicIfError(err)

	return cartItem, nil
}

func (r *cartRepositoryImpl) UpdateCart(ctx context.Context, db *gorm.DB, cart models.Cart) (models.Cart, error) {
	err := db.WithContext(ctx).Updates(&cart).Error
	helpers.PanicIfError(err)

	return cart, nil
}

func (r *cartRepositoryImpl) DeleteCart(ctx context.Context, db *gorm.DB, cart models.Cart) error {
	err := db.WithContext(ctx).Where(&models.Cart{}).Where("id = ?", cart.ID).Delete(&cart).Error
	helpers.PanicIfError(err)

	return nil
}
