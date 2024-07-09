package services

import (
	"context"

	"github.com/hiboedi/go-store-backend/app/helpers"
	"github.com/hiboedi/go-store-backend/app/web/models"
	"github.com/hiboedi/go-store-backend/app/web/repositories"
	"gorm.io/gorm"
)

type OrderService interface {
	FindAllOrder(ctx context.Context) ([]models.OrderResponse, error)
}

type OrderRepositoryImpl struct {
	OrderRepository repositories.OrderRepository
	DB              *gorm.DB
}

func NewOrderService(orderRepo repositories.OrderRepository, db *gorm.DB) OrderService {
	return &OrderRepositoryImpl{
		OrderRepository: orderRepo,
		DB:              db,
	}
}

func (s *OrderRepositoryImpl) FindAllOrder(ctx context.Context) ([]models.OrderResponse, error) {
	tx := s.DB.Begin()
	defer helpers.CommitOrRollback(tx)

	data, err := s.OrderRepository.FindAllOrder(ctx, tx)
	helpers.PanicIfError(err)

	return models.ToOrderResponses(data), nil
}
