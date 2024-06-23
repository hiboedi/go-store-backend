package services

import (
	"context"

	"github.com/go-playground/validator/v10"
	"github.com/hiboedi/go-store-backend/app/exceptions"
	"github.com/hiboedi/go-store-backend/app/helpers"
	"github.com/hiboedi/go-store-backend/app/web/models"
	"github.com/hiboedi/go-store-backend/app/web/repositories"
	"gorm.io/gorm"
)

type OrderItemServiceImpl struct {
	OrderItemRepository repositories.OrderItemRepository
	DB                  *gorm.DB
	Validate            *validator.Validate
}

type OrderItemService interface {
	Create(ctx context.Context, request models.OrderItemCreate) models.OrderItemResponse
	Update(ctx context.Context, request models.OrderItemUpdate) models.OrderItemResponse
	Delete(ctx context.Context, orderItemId string)
	FindByID(ctx context.Context, orderItemId string) models.OrderItemResponse
	FindAllByOrderID(ctx context.Context, orderId string) []models.OrderItemResponse
}

func NewOrderItemService(orderItemRepo repositories.OrderItemRepository, db *gorm.DB, validate *validator.Validate) OrderItemService {
	return &OrderItemServiceImpl{
		OrderItemRepository: orderItemRepo,
		DB:                  db,
		Validate:            validate,
	}
}

func (s *OrderItemServiceImpl) Create(ctx context.Context, request models.OrderItemCreate) models.OrderItemResponse {
	err := s.Validate.Struct(request)
	helpers.PanicIfError(err)

	tx := s.DB.Begin()
	defer helpers.CommitOrRollback(tx)

	orderItem := models.OrderItem{
		OrderID:   request.OrderID,
		ProductID: request.ProductID,
		Quantity:  request.Quantity,
		Price:     request.Price,
	}

	data, err := s.OrderItemRepository.CreateOrderItem(ctx, tx, orderItem)
	helpers.PanicIfError(err)

	return models.ToOrderItemResponse(data)
}

func (s *OrderItemServiceImpl) Update(ctx context.Context, request models.OrderItemUpdate) models.OrderItemResponse {
	err := s.Validate.Struct(request)
	helpers.PanicIfError(err)

	tx := s.DB.Begin()
	defer helpers.CommitOrRollback(tx)

	orderItem, err := s.OrderItemRepository.GetOrderItemById(ctx, tx, request.ID)
	if err != nil {
		panic(exceptions.NewNotFoundError(err.Error()))
	}

	orderItem.OrderID = request.OrderID
	orderItem.ProductID = request.ProductID
	orderItem.Quantity = request.Quantity
	orderItem.Price = request.Price

	data, err := s.OrderItemRepository.UpdateOrderItem(ctx, tx, orderItem)
	helpers.PanicIfError(err)

	return models.ToOrderItemResponse(data)
}

func (s *OrderItemServiceImpl) Delete(ctx context.Context, orderItemId string) {
	tx := s.DB.Begin()
	defer helpers.CommitOrRollback(tx)

	orderItem, err := s.OrderItemRepository.GetOrderItemById(ctx, tx, orderItemId)
	if err != nil {
		panic(exceptions.NewNotFoundError(err.Error()))
	}

	err = s.OrderItemRepository.DeleteOrderItem(ctx, tx, orderItem)
	helpers.PanicIfError(err)
}

func (s *OrderItemServiceImpl) FindByID(ctx context.Context, orderItemId string) models.OrderItemResponse {
	tx := s.DB.Begin()
	defer helpers.CommitOrRollback(tx)

	orderItem, err := s.OrderItemRepository.GetOrderItemById(ctx, tx, orderItemId)
	helpers.PanicIfError(err)

	return models.ToOrderItemResponse(orderItem)
}

func (s *OrderItemServiceImpl) FindAllByOrderID(ctx context.Context, orderId string) []models.OrderItemResponse {
	tx := s.DB.Begin()
	defer helpers.CommitOrRollback(tx)

	orderItems, err := s.OrderItemRepository.FindOrderItemsByOrderId(ctx, tx, orderId)
	helpers.PanicIfError(err)

	return models.ToOrderItemResponses(orderItems)
}
