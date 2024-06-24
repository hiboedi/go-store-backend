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

type SizeServiceImpl struct {
	SizeRepository repositories.SizeRepository
	DB             *gorm.DB
	Validate       *validator.Validate
}

type SizeService interface {
	Create(ctx context.Context, request models.SizeCreate) models.SizeResponseHiddenStore
	Update(ctx context.Context, request models.SizeUpdate, sizeId string) models.SizeResponseHiddenStore
	Delete(ctx context.Context, sizeId string)
	FindById(ctx context.Context, sizeId string) models.SizeResponse
	FindAll(ctx context.Context) []models.SizeResponse
}

func NewSizeService(sizeRepo repositories.SizeRepository, db *gorm.DB, validate *validator.Validate) SizeService {
	return &SizeServiceImpl{
		SizeRepository: sizeRepo,
		DB:             db,
		Validate:       validate,
	}
}

func (s *SizeServiceImpl) Create(ctx context.Context, request models.SizeCreate) models.SizeResponseHiddenStore {
	err := s.Validate.Struct(request)
	helpers.PanicIfError(err)

	tx := s.DB.Begin()
	defer helpers.CommitOrRollback(tx)

	size := models.Size{
		StoreID: request.StoreID,
		Name:    request.Name,
		Value:   request.Value,
	}

	data, err := s.SizeRepository.CreateSize(ctx, tx, size)
	helpers.PanicIfError(err)

	return models.ToSizeResponseHiddenStore(data)
}

func (s *SizeServiceImpl) Update(ctx context.Context, request models.SizeUpdate, sizeId string) models.SizeResponseHiddenStore {
	err := s.Validate.Struct(request)
	helpers.PanicIfError(err)

	tx := s.DB.Begin()
	defer helpers.CommitOrRollback(tx)

	size, err := s.SizeRepository.GetSizeById(ctx, tx, sizeId)
	if err != nil {
		panic(exceptions.NewNotFoundError(err.Error()))
	}

	size.Name = request.Name
	size.Value = request.Value

	data, err := s.SizeRepository.UpdateSize(ctx, tx, size)
	helpers.PanicIfError(err)

	return models.ToSizeResponseHiddenStore(data)
}

func (s *SizeServiceImpl) Delete(ctx context.Context, sizeId string) {
	tx := s.DB.Begin()
	defer helpers.CommitOrRollback(tx)

	size, err := s.SizeRepository.GetSizeById(ctx, tx, sizeId)
	if err != nil {
		panic(exceptions.NewNotFoundError(err.Error()))
	}

	err = s.SizeRepository.DeleteSize(ctx, tx, size)
	helpers.PanicIfError(err)
}

func (s *SizeServiceImpl) FindAll(ctx context.Context) []models.SizeResponse {
	tx := s.DB.Begin()
	defer helpers.CommitOrRollback(tx)

	sizes, err := s.SizeRepository.FindAllSizes(ctx, tx)
	helpers.PanicIfError(err)

	return models.ToSizeResponses(sizes)
}

func (s *SizeServiceImpl) FindById(ctx context.Context, sizeId string) models.SizeResponse {
	tx := s.DB.Begin()
	defer helpers.CommitOrRollback(tx)

	size, err := s.SizeRepository.GetSizeById(ctx, tx, sizeId)
	helpers.PanicIfError(err)

	return models.ToSizeResponse(size)
}
