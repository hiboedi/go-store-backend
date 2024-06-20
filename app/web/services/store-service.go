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

type StoreService interface {
	Create(ctx context.Context, request models.StoreCreate) models.StoreResponse
	Update(ctx context.Context, request models.StoreUpdate) models.StoreResponse
	Delete(ctx context.Context, storeId string)
	FindById(ctx context.Context, storeId string) models.StoreResponse
	FindAll(ctx context.Context) []models.StoreResponse
}

type StoreServiceImpl struct {
	StoreRepository repositories.StoreRepository
	DB              *gorm.DB
	Validate        *validator.Validate
}

func NewStoreService(storeRepo repositories.StoreRepository, db *gorm.DB, validate *validator.Validate) StoreService {
	return &StoreServiceImpl{
		StoreRepository: storeRepo,
		DB:              db,
		Validate:        validate,
	}
}

func (s *StoreServiceImpl) Create(ctx context.Context, request models.StoreCreate) models.StoreResponse {
	err := s.Validate.Struct(request)
	helpers.PanicIfError(err)

	tx := s.DB.Begin()
	defer helpers.CommitOrRollback(tx)

	store := models.Store{
		Name:   request.Name,
		UserID: request.UserID,
	}

	data, err := s.StoreRepository.CreateStore(ctx, tx, store)
	helpers.PanicIfError(err)

	return models.ToStoreResponse(data)
}

func (s *StoreServiceImpl) Update(ctx context.Context, request models.StoreUpdate) models.StoreResponse {
	err := s.Validate.Struct(request)
	helpers.PanicIfError(err)

	tx := s.DB.Begin()
	defer helpers.CommitOrRollback(tx)

	store, err := s.StoreRepository.GetStoreById(ctx, tx, request.ID)
	if err != nil {
		panic(exceptions.NewNotFoundError(err.Error()))
	}

	store.Name = request.Name
	// store.Billboards = request.Billboards
	// store.Categories = request.Categories
	// store.Sizes = request.Sizes
	// store.Colors = request.Colors
	// store.Products = request.Products
	// store.Orders = request.Orders

	data, err := s.StoreRepository.UpdateStore(ctx, tx, store)
	helpers.PanicIfError(err)

	return models.ToStoreResponse(data)
}

func (s *StoreServiceImpl) Delete(ctx context.Context, storeId string) {
	tx := s.DB.Begin()
	defer helpers.CommitOrRollback(tx)

	store, err := s.StoreRepository.GetStoreById(ctx, tx, storeId)
	if err != nil {
		panic(exceptions.NewNotFoundError(err.Error()))
	}

	err = s.StoreRepository.DeleteStore(ctx, tx, store)
	helpers.PanicIfError(err)
}

func (s *StoreServiceImpl) FindAll(ctx context.Context) []models.StoreResponse {
	tx := s.DB.Begin()
	defer helpers.CommitOrRollback(tx)

	stores, err := s.StoreRepository.FindAllStore(ctx, tx)
	helpers.PanicIfError(err)
	return models.ToStoreResponses(stores)
}

func (s *StoreServiceImpl) FindById(ctx context.Context, storeId string) models.StoreResponse {
	tx := s.DB.Begin()
	defer helpers.CommitOrRollback(tx)

	store, err := s.StoreRepository.GetStoreById(ctx, tx, storeId)
	helpers.PanicIfError(err)
	return models.ToStoreResponse(store)
}
