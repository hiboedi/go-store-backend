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

type OrderServiceImpl struct {
	OrderRepository  repositories.OrderRepository
	OrderItemService OrderItemService
	DB               *gorm.DB
	Validate         *validator.Validate
}

type OrderService interface {
	Create(ctx context.Context, request models.OrderCreate) models.OrderResponseHiddenStore
	Update(ctx context.Context, request models.OrderUpdate) models.OrderResponseHiddenStore
	Delete(ctx context.Context, orderId string)
	FindById(ctx context.Context, orderId string) models.OrderResponse
	FindAll(ctx context.Context) []models.OrderResponse
}

func NewOrderService(orderRepo repositories.OrderRepository, orderItemService OrderItemService, db *gorm.DB, validate *validator.Validate) OrderService {
	return &OrderServiceImpl{
		OrderRepository:  orderRepo,
		OrderItemService: orderItemService,
		DB:               db,
		Validate:         validate,
	}
}

func (s *OrderServiceImpl) Create(ctx context.Context, request models.OrderCreate) models.OrderResponseHiddenStore {
	err := s.Validate.Struct(request)
	helpers.PanicIfError(err)

	tx := s.DB.Begin()
	defer helpers.CommitOrRollback(tx)

	order := models.Order{
		StoreID:    request.StoreID,
		OrderItems: request.OrderItems,
		IsPaid:     request.IsPaid,
		Phone:      request.Phone,
		Address:    request.Address,
	}

	data, err := s.OrderRepository.CreateOrder(ctx, tx, order)
	helpers.PanicIfError(err)

	return models.ToOrderResponseHiddenStore(data)
}

func (s *OrderServiceImpl) Update(ctx context.Context, request models.OrderUpdate) models.OrderResponseHiddenStore {
	err := s.Validate.Struct(request)
	helpers.PanicIfError(err)

	tx := s.DB.Begin()
	defer helpers.CommitOrRollback(tx)

	order, err := s.OrderRepository.GetOrderById(ctx, tx, request.ID)
	if err != nil {
		panic(exceptions.NewNotFoundError(err.Error()))
	}

	order.OrderItems = request.OrderItems
	order.IsPaid = request.IsPaid
	order.Phone = request.Phone
	order.Address = request.Address

	data, err := s.OrderRepository.UpdateOrder(ctx, tx, order)
	helpers.PanicIfError(err)

	return models.ToOrderResponseHiddenStore(data)
}

func (s *OrderServiceImpl) Delete(ctx context.Context, orderId string) {
	tx := s.DB.Begin()
	defer helpers.CommitOrRollback(tx)

	order, err := s.OrderRepository.GetOrderById(ctx, tx, orderId)
	if err != nil {
		panic(exceptions.NewNotFoundError(err.Error()))
	}

	err = s.OrderRepository.DeleteOrder(ctx, tx, order)
	helpers.PanicIfError(err)
}

func (s *OrderServiceImpl) FindAll(ctx context.Context) []models.OrderResponse {
	tx := s.DB.Begin()
	defer helpers.CommitOrRollback(tx)

	orders, err := s.OrderRepository.FindAllOrders(ctx, tx)
	helpers.PanicIfError(err)
	return models.ToOrderResponses(orders)
}

func (s *OrderServiceImpl) FindById(ctx context.Context, orderId string) models.OrderResponse {
	tx := s.DB.Begin()
	defer helpers.CommitOrRollback(tx)

	order, err := s.OrderRepository.GetOrderById(ctx, tx, orderId)
	helpers.PanicIfError(err)
	return models.ToOrderResponse(order)
}
