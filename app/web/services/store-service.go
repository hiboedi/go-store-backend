package services

import (
	"context"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/hiboedi/go-store-backend/app/exceptions"
	"github.com/hiboedi/go-store-backend/app/helpers"
	"github.com/hiboedi/go-store-backend/app/web/models"
	"github.com/hiboedi/go-store-backend/app/web/repositories"
	"gorm.io/gorm"
)

type StoreService interface {
	Create(ctx context.Context, request models.StoreCreate, userID string) models.StoreResponse
	Update(ctx context.Context, request models.StoreUpdate, storeId string) models.StoreResponse
	Delete(ctx context.Context, storeId string)
	FindById(ctx context.Context, storeId string) models.StoreResponse
	FindAllByUserId(ctx context.Context, userId string) []models.StoreResponse
}

type StoreServiceImpl struct {
	StoreRepository repositories.StoreRepository
	UserRepository  repositories.UserRepository
	DB              *gorm.DB
	Validate        *validator.Validate
}

func NewStoreService(storeRepo repositories.StoreRepository, userRepo repositories.UserRepository, db *gorm.DB, validate *validator.Validate) StoreService {
	return &StoreServiceImpl{
		UserRepository:  userRepo,
		StoreRepository: storeRepo,
		DB:              db,
		Validate:        validate,
	}
}

func (s *StoreServiceImpl) Create(ctx context.Context, request models.StoreCreate, userID string) models.StoreResponse {
	err := s.Validate.Struct(request)
	helpers.PanicIfError(err)

	tx := s.DB.Begin()
	defer helpers.CommitOrRollback(tx)

	user, err := s.UserRepository.GetUserById(ctx, tx, userID)
	if err != nil {
		panic(exceptions.NewNotFoundError(err.Error()))
	}

	store := models.Store{
		ID:     uuid.New().String(),
		Name:   request.Name,
		UserID: userID,
		User:   user,
	}

	data, err := s.StoreRepository.CreateStore(ctx, tx, store)
	helpers.PanicIfError(err)

	return models.ToStoreResponse(data)
}

func (s *StoreServiceImpl) Update(ctx context.Context, request models.StoreUpdate, storeId string) models.StoreResponse {
	err := s.Validate.Struct(request)
	helpers.PanicIfError(err)

	tx := s.DB.Begin()
	defer helpers.CommitOrRollback(tx)

	store, err := s.StoreRepository.GetStoreById(ctx, tx, storeId)
	if err != nil {
		panic(exceptions.NewNotFoundError(err.Error()))
	}

	store.Name = request.Name

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

func (s *StoreServiceImpl) FindAllByUserId(ctx context.Context, userId string) []models.StoreResponse {
	tx := s.DB.Begin()
	defer helpers.CommitOrRollback(tx)

	stores, err := s.StoreRepository.FindAllStore(ctx, tx, userId)
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
