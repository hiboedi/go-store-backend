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

type ColorServiceImpl struct {
	ColorRepository repositories.ColorRepository
	DB              *gorm.DB
	Validate        *validator.Validate
}

type ColorService interface {
	Create(ctx context.Context, request models.ColorCreate) models.ColorResponseHiddenStore
	Update(ctx context.Context, request models.ColorUpdate, colorId string) models.ColorResponseHiddenStore
	Delete(ctx context.Context, colorId string)
	FindById(ctx context.Context, colorId string) models.ColorResponse
	FindAll(ctx context.Context) []models.ColorResponse
}

func NewColorService(colorRepo repositories.ColorRepository, db *gorm.DB, validate *validator.Validate) ColorService {
	return &ColorServiceImpl{
		ColorRepository: colorRepo,
		DB:              db,
		Validate:        validate,
	}
}

func (s *ColorServiceImpl) Create(ctx context.Context, request models.ColorCreate) models.ColorResponseHiddenStore {
	err := s.Validate.Struct(request)
	helpers.PanicIfError(err)

	tx := s.DB.Begin()
	defer helpers.CommitOrRollback(tx)

	color := models.Color{
		Name:    request.Name,
		Value:   request.Value,
		StoreID: request.StoreID,
	}

	data, err := s.ColorRepository.CreateColor(ctx, tx, color)
	helpers.PanicIfError(err)

	return models.ToColorResponseHiddenStore(data)
}

func (s *ColorServiceImpl) Update(ctx context.Context, request models.ColorUpdate, colorId string) models.ColorResponseHiddenStore {
	err := s.Validate.Struct(request)
	helpers.PanicIfError(err)

	tx := s.DB.Begin()
	defer helpers.CommitOrRollback(tx)

	color, err := s.ColorRepository.GetColorById(ctx, tx, colorId)
	if err != nil {
		panic(exceptions.NewNotFoundError(err.Error()))
	}

	color.Name = request.Name
	color.Value = request.Value

	data, err := s.ColorRepository.UpdateColor(ctx, tx, color)
	helpers.PanicIfError(err)

	return models.ToColorResponseHiddenStore(data)
}

func (s *ColorServiceImpl) Delete(ctx context.Context, colorId string) {
	tx := s.DB.Begin()
	defer helpers.CommitOrRollback(tx)

	color, err := s.ColorRepository.GetColorById(ctx, tx, colorId)
	if err != nil {
		panic(exceptions.NewNotFoundError(err.Error()))
	}

	err = s.ColorRepository.DeleteColor(ctx, tx, color)
	helpers.PanicIfError(err)
}

func (s *ColorServiceImpl) FindAll(ctx context.Context) []models.ColorResponse {
	tx := s.DB.Begin()
	defer helpers.CommitOrRollback(tx)

	colors, err := s.ColorRepository.FindAllColors(ctx, tx)
	helpers.PanicIfError(err)

	return models.ToColorResponses(colors)
}

func (s *ColorServiceImpl) FindById(ctx context.Context, colorId string) models.ColorResponse {
	tx := s.DB.Begin()
	defer helpers.CommitOrRollback(tx)

	color, err := s.ColorRepository.GetColorById(ctx, tx, colorId)
	if err != nil {
		panic(exceptions.NewNotFoundError(err.Error()))
	}

	return models.ToColorResponse(color)
}
