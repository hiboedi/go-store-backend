package repositories

import (
	"context"

	"github.com/hiboedi/go-store-backend/app/helpers"
	"github.com/hiboedi/go-store-backend/app/web/models"
	"gorm.io/gorm"
)

type OrderRepository interface {
	CreateOrder(ctx context.Context, db *gorm.DB, order models.Order) (models.Order, error)
	UpdateOrder(ctx context.Context, db *gorm.DB, order models.Order) (models.Order, error)
	CreateOrderItem(ctx context.Context, db *gorm.DB, orderItem models.OrderItem) (models.OrderItem, error)
	FindAllOrder(ctx context.Context, db *gorm.DB) ([]models.Order, error)
}

type orderRepositoryImpl struct {
}

func NewOrderRepository() OrderRepository {
	return &orderRepositoryImpl{}
}

func (r *orderRepositoryImpl) CreateOrder(ctx context.Context, db *gorm.DB, order models.Order) (models.Order, error) {

	err := db.WithContext(ctx).Create(&order).Error
	helpers.PanicIfError(err)

	return order, nil
}

func (r *orderRepositoryImpl) UpdateOrder(ctx context.Context, db *gorm.DB, order models.Order) (models.Order, error) {

	err := db.WithContext(ctx).Model(&models.Order{}).Where("id = ?", order.ID).Updates(&order).Error
	helpers.PanicIfError(err)

	return order, nil
}

func (r *orderRepositoryImpl) CreateOrderItem(ctx context.Context, db *gorm.DB, orderItem models.OrderItem) (models.OrderItem, error) {

	err := db.WithContext(ctx).Create(&orderItem).Error
	helpers.PanicIfError(err)

	return orderItem, nil
}

func (r *orderRepositoryImpl) FindAllOrder(ctx context.Context, db *gorm.DB) ([]models.Order, error) {
	var Orders []models.Order

	err := db.WithContext(ctx).Model(&models.Order{}).Find(&Orders).Error
	helpers.PanicIfError(err)

	return Orders, nil
}
