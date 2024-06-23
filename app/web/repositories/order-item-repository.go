package repositories

import (
	"context"

	"github.com/google/uuid"
	"github.com/hiboedi/go-store-backend/app/helpers"
	"github.com/hiboedi/go-store-backend/app/web/models"
	"gorm.io/gorm"
)

type OrderItemRepositoryImpl struct {
}

type OrderItemRepository interface {
	CreateOrderItem(ctx context.Context, db *gorm.DB, orderItem models.OrderItem) (models.OrderItem, error)
	UpdateOrderItem(ctx context.Context, db *gorm.DB, orderItem models.OrderItem) (models.OrderItem, error)
	DeleteOrderItem(ctx context.Context, db *gorm.DB, orderItem models.OrderItem) error
	GetOrderItemById(ctx context.Context, db *gorm.DB, orderItemId string) (models.OrderItem, error)
	FindOrderItemsByOrderId(ctx context.Context, db *gorm.DB, orderId string) ([]models.OrderItem, error)
}

func NewOrderItemRepository() OrderItemRepository {
	return &OrderItemRepositoryImpl{}
}

func (r *OrderItemRepositoryImpl) CreateOrderItem(ctx context.Context, db *gorm.DB, orderItem models.OrderItem) (models.OrderItem, error) {
	orderItemModel := models.OrderItem{
		ID:        uuid.New().String(),
		OrderID:   orderItem.OrderID,
		ProductID: orderItem.ProductID,
		Quantity:  orderItem.Quantity,
		Price:     orderItem.Price,
	}

	err := db.WithContext(ctx).Create(&orderItemModel).Error
	helpers.PanicIfError(err)

	return orderItemModel, nil
}

func (r *OrderItemRepositoryImpl) UpdateOrderItem(ctx context.Context, db *gorm.DB, orderItem models.OrderItem) (models.OrderItem, error) {
	orderItemModel := models.OrderItem{
		OrderID:   orderItem.OrderID,
		ProductID: orderItem.ProductID,
		Quantity:  orderItem.Quantity,
		Price:     orderItem.Price,
	}

	err := db.WithContext(ctx).Model(&models.OrderItem{}).Where("id = ?", orderItem.ID).Updates(&orderItemModel).Error
	helpers.PanicIfError(err)

	return orderItemModel, nil
}

func (r *OrderItemRepositoryImpl) DeleteOrderItem(ctx context.Context, db *gorm.DB, orderItem models.OrderItem) error {
	err := db.WithContext(ctx).Model(&models.OrderItem{}).Where("id = ?", orderItem.ID).Delete(&orderItem).Error
	helpers.PanicIfError(err)

	return nil
}

func (r *OrderItemRepositoryImpl) GetOrderItemById(ctx context.Context, db *gorm.DB, orderItemId string) (models.OrderItem, error) {
	var orderItem models.OrderItem
	err := db.WithContext(ctx).Model(&models.OrderItem{}).Where("id = ?", orderItemId).Take(&orderItem).Error
	helpers.PanicIfError(err)

	return orderItem, nil
}

func (r *OrderItemRepositoryImpl) FindOrderItemsByOrderId(ctx context.Context, db *gorm.DB, orderId string) ([]models.OrderItem, error) {
	var orderItems []models.OrderItem
	err := db.WithContext(ctx).Model(&models.OrderItem{}).Where("order_id = ?", orderId).Find(&orderItems).Error
	helpers.PanicIfError(err)

	return orderItems, nil
}
