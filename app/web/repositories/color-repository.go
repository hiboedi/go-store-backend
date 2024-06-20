package repositories

import (
	"context"

	"github.com/google/uuid"
	"github.com/hiboedi/go-store-backend/app/helpers"
	"github.com/hiboedi/go-store-backend/app/web/models"
	"gorm.io/gorm"
)

type ColorRepositoryImpl struct {
}

type ColorRepository interface {
	CreateColor(ctx context.Context, db *gorm.DB, color models.Color) (models.Color, error)
	UpdateColor(ctx context.Context, db *gorm.DB, color models.Color) (models.Color, error)
	DeleteColor(ctx context.Context, db *gorm.DB, color models.Color) error
	GetColorById(ctx context.Context, db *gorm.DB, colorId string) (models.Color, error)
	FindAllColors(ctx context.Context, db *gorm.DB) ([]models.Color, error)
}

func NewColorRepository() ColorRepository {
	return &ColorRepositoryImpl{}
}

func (r *ColorRepositoryImpl) CreateColor(ctx context.Context, db *gorm.DB, color models.Color) (models.Color, error) {
	colorModel := models.Color{
		ID:      uuid.New().String(),
		Name:    color.Name,
		StoreID: color.StoreID,
		Value:   color.Value,
	}

	err := db.WithContext(ctx).Create(&colorModel).Error
	helpers.PanicIfError(err)

	return colorModel, nil
}

func (r *ColorRepositoryImpl) UpdateColor(ctx context.Context, db *gorm.DB, color models.Color) (models.Color, error) {

	colorModel := models.Color{
		Name:  color.Name,
		Value: color.Value,
	}
	err := db.WithContext(ctx).Model(&models.Color{}).Where("id = ?", color.ID).Updates(&colorModel).Error
	helpers.PanicIfError(err)

	return color, nil
}

func (r *ColorRepositoryImpl) DeleteColor(ctx context.Context, db *gorm.DB, color models.Color) error {
	err := db.WithContext(ctx).Model(&models.Color{}).Where("id = ?", color.ID).Delete(&color).Error
	helpers.PanicIfError(err)

	return nil
}

func (r *ColorRepositoryImpl) GetColorById(ctx context.Context, db *gorm.DB, colorId string) (models.Color, error) {
	var color models.Color
	err := db.WithContext(ctx).Model(&models.Color{}).Where("id = ?", colorId).Take(&color).Error
	helpers.PanicIfError(err)

	return color, nil
}

func (r *ColorRepositoryImpl) FindAllColors(ctx context.Context, db *gorm.DB) ([]models.Color, error) {
	var colors []models.Color
	err := db.WithContext(ctx).Model(&models.Color{}).Find(&colors).Error
	helpers.PanicIfError(err)

	return colors, nil
}
