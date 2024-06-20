package repositories

import (
	"context"

	"github.com/google/uuid"
	"github.com/hiboedi/go-store-backend/app/helpers"
	"github.com/hiboedi/go-store-backend/app/web/models"
	"gorm.io/gorm"
)

type CategoryRepositoryImpl struct {
}

type CategoryRepository interface {
	CreateCategory(ctx context.Context, db *gorm.DB, category models.Category) (models.Category, error)
	UpdateCategory(ctx context.Context, db *gorm.DB, category models.Category) (models.Category, error)
	DeleteCategory(ctx context.Context, db *gorm.DB, category models.Category) error
	GetCategoryById(ctx context.Context, db *gorm.DB, categoryId string) (models.Category, error)
	FindAllCategories(ctx context.Context, db *gorm.DB) ([]models.Category, error)
}

func NewCategoryRepository() CategoryRepository {
	return &CategoryRepositoryImpl{}
}

func (r *CategoryRepositoryImpl) CreateCategory(ctx context.Context, db *gorm.DB, category models.Category) (models.Category, error) {
	categoryModel := models.Category{
		ID:          uuid.New().String(),
		StoreID:     category.StoreID,
		BillboardID: category.BillboardID,
		Name:        category.Name,
	}

	err := db.WithContext(ctx).Create(&categoryModel).Error
	helpers.PanicIfError(err)

	return categoryModel, nil
}

func (r *CategoryRepositoryImpl) UpdateCategory(ctx context.Context, db *gorm.DB, category models.Category) (models.Category, error) {
	categoryModel := models.Category{
		BillboardID: category.BillboardID,
		Name:        category.Name,
	}

	err := db.WithContext(ctx).Model(&models.Category{}).Where("id = ?", category.ID).Updates(&categoryModel).Error
	helpers.PanicIfError(err)

	return categoryModel, nil
}

func (r *CategoryRepositoryImpl) DeleteCategory(ctx context.Context, db *gorm.DB, category models.Category) error {
	err := db.WithContext(ctx).Model(&models.Category{}).Where("id = ?", category.ID).Delete(&category).Error
	helpers.PanicIfError(err)

	return nil
}

func (r *CategoryRepositoryImpl) GetCategoryById(ctx context.Context, db *gorm.DB, categoryId string) (models.Category, error) {
	var category models.Category

	err := db.WithContext(ctx).Model(&models.Category{}).
		Preload("Billboard").
		Where("id = ?", categoryId).
		Take(&category).Error
	helpers.PanicIfError(err)

	return category, nil
}

func (r *CategoryRepositoryImpl) FindAllCategories(ctx context.Context, db *gorm.DB) ([]models.Category, error) {
	var categories []models.Category

	err := db.WithContext(ctx).Model(&models.Category{}).
		Preload("Billboard").
		Find(&categories).Error
	helpers.PanicIfError(err)

	return categories, nil
}
