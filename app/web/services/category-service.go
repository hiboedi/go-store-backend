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

type CategoryServiceImpl struct {
	CategoryRepository repositories.CategoryRepository
	DB                 *gorm.DB
	Validate           *validator.Validate
}

type CategoryService interface {
	Create(ctx context.Context, request models.CategoryCreate) models.CategoryResponse
	Update(ctx context.Context, request models.CategoryUpdate, categoryId string) models.CategoryResponse
	Delete(ctx context.Context, categoryId string)
	FindById(ctx context.Context, categoryId string) models.CategoryResponse
	FindAll(ctx context.Context) []models.CategoryResponse
}

func NewCategoryService(categoryRepo repositories.CategoryRepository, db *gorm.DB, validate *validator.Validate) CategoryService {
	return &CategoryServiceImpl{
		CategoryRepository: categoryRepo,
		DB:                 db,
		Validate:           validate,
	}
}

func (s *CategoryServiceImpl) Create(ctx context.Context, request models.CategoryCreate) models.CategoryResponse {
	err := s.Validate.Struct(request)
	helpers.PanicIfError(err)

	tx := s.DB.Begin()
	defer helpers.CommitOrRollback(tx)

	category := models.Category{
		StoreID:     request.StoreID,
		BillboardID: request.BillboardID,
		Name:        request.Name,
	}

	data, err := s.CategoryRepository.CreateCategory(ctx, tx, category)
	helpers.PanicIfError(err)

	return models.ToCategoryResponse(data)
}

func (s *CategoryServiceImpl) Update(ctx context.Context, request models.CategoryUpdate, categoryId string) models.CategoryResponse {
	err := s.Validate.Struct(request)
	helpers.PanicIfError(err)

	tx := s.DB.Begin()
	defer helpers.CommitOrRollback(tx)

	category, err := s.CategoryRepository.GetCategoryById(ctx, tx, categoryId)
	if err != nil {
		panic(exceptions.NewNotFoundError(err.Error()))
	}

	category.Name = request.Name
	category.BillboardID = request.BillboardID

	data, err := s.CategoryRepository.UpdateCategory(ctx, tx, category)
	helpers.PanicIfError(err)

	return models.ToCategoryResponse(data)
}

func (s *CategoryServiceImpl) Delete(ctx context.Context, categoryId string) {
	tx := s.DB.Begin()
	defer helpers.CommitOrRollback(tx)

	category, err := s.CategoryRepository.GetCategoryById(ctx, tx, categoryId)
	if err != nil {
		panic(exceptions.NewNotFoundError(err.Error()))
	}

	err = s.CategoryRepository.DeleteCategory(ctx, tx, category)
	helpers.PanicIfError(err)
}

func (s *CategoryServiceImpl) FindAll(ctx context.Context) []models.CategoryResponse {
	tx := s.DB.Begin()
	defer helpers.CommitOrRollback(tx)

	categories, err := s.CategoryRepository.FindAllCategories(ctx, tx)
	helpers.PanicIfError(err)

	return models.ToCategoryResponses(categories)
}

func (s *CategoryServiceImpl) FindById(ctx context.Context, categoryId string) models.CategoryResponse {
	tx := s.DB.Begin()
	defer helpers.CommitOrRollback(tx)

	category, err := s.CategoryRepository.GetCategoryById(ctx, tx, categoryId)
	if err != nil {
		panic(exceptions.NewNotFoundError(err.Error()))
	}

	return models.ToCategoryResponse(category)
}
