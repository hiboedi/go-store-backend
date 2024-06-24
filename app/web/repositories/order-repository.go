package repositories

import (
	"context"

	"github.com/google/uuid"
	"github.com/hiboedi/go-store-backend/app/helpers"
	"github.com/hiboedi/go-store-backend/app/web/models"
	"gorm.io/gorm"
)

type OrderRepositoryImpl struct {
}

type OrderRepository interface {
	CreateOrder(ctx context.Context, db *gorm.DB, order models.Order) (models.Order, error)
	UpdateOrder(ctx context.Context, db *gorm.DB, order models.Order) (models.Order, error)
	DeleteOrder(ctx context.Context, db *gorm.DB, order models.Order) error
	GetOrderById(ctx context.Context, db *gorm.DB, orderId string) (models.Order, error)
	FindAllOrders(ctx context.Context, db *gorm.DB) ([]models.Order, error)
}

func NewOrderRepository() OrderRepository {
	return &OrderRepositoryImpl{}
}

func (r *OrderRepositoryImpl) CreateOrder(ctx context.Context, db *gorm.DB, order models.Order) (models.Order, error) {
	discountAmount := helpers.GetDiscountAmount(order.BaseTotalPrice)
	taxAmount := helpers.GetTaxAmount(order.BaseTotalPrice)

	orderModel := models.Order{
		ID:              uuid.New().String(),
		StoreID:         order.StoreID,
		OrderItems:      order.OrderItems,
		IsPaid:          order.IsPaid,
		Phone:           order.Phone,
		Address:         order.Address,
		BaseTotalPrice:  order.BaseTotalPrice,
		DiscountAmount:  discountAmount,
		DiscountPercent: float64(helpers.DiscountPercent),
		TaxAmount:       taxAmount,
		TaxPercent:      float64(helpers.TaxPercent),
		GrandTotaPrice:  order.BaseTotalPrice - taxAmount - discountAmount,
	}

	err := db.WithContext(ctx).Create(&orderModel).Error
	helpers.PanicIfError(err)

	return orderModel, nil
}

func (r *OrderRepositoryImpl) UpdateOrder(ctx context.Context, db *gorm.DB, order models.Order) (models.Order, error) {

	orderModel := models.Order{
		ID:              order.ID,
		StoreID:         order.StoreID,
		OrderItems:      order.OrderItems,
		IsPaid:          order.IsPaid,
		Phone:           order.Phone,
		Address:         order.Address,
		BaseTotalPrice:  order.BaseTotalPrice,
		DiscountAmount:  order.DiscountAmount,
		DiscountPercent: order.DiscountPercent,
		TaxPercent:      order.TaxPercent,
		TaxAmount:       order.TaxAmount,
		GrandTotaPrice:  order.GrandTotaPrice,
		CreatedAt:       order.CreatedAt,
		UpdatedAt:       order.UpdatedAt,
	}

	err := db.WithContext(ctx).Model(&models.Order{}).Where("id = ?", order.ID).Updates(&orderModel).Error
	helpers.PanicIfError(err)

	return orderModel, nil
}

func (r *OrderRepositoryImpl) GetOrderById(ctx context.Context, db *gorm.DB, orderId string) (models.Order, error) {
	var order models.Order
	err := db.WithContext(ctx).Model(&models.Order{}).Preload("OrderItems").Where("id = ?", orderId).Take(&order).Error
	helpers.PanicIfError(err)

	return order, nil
}

func (r *OrderRepositoryImpl) DeleteOrder(ctx context.Context, db *gorm.DB, order models.Order) error {
	err := db.WithContext(ctx).Model(&models.Order{}).Where("id = ?", order.ID).Delete(&order).Error
	helpers.PanicIfError(err)

	return nil
}

func (r *OrderRepositoryImpl) FindAllOrders(ctx context.Context, db *gorm.DB) ([]models.Order, error) {
	var orders []models.Order

	err := db.WithContext(ctx).Model(&models.Order{}).Preload("OrderItems").Find(&orders).Error
	helpers.PanicIfError(err)

	return orders, nil
}
