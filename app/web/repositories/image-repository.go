package repositories

import (
	"context"

	"github.com/google/uuid"
	"github.com/hiboedi/go-store-backend/app/helpers"
	"github.com/hiboedi/go-store-backend/app/web/models"
	"gorm.io/gorm"
)

type ImageRepositoryImpl struct {
}

type ImageRepository interface {
	CreateImage(ctx context.Context, db *gorm.DB, image models.Image) (models.Image, error)
	UpdateImage(ctx context.Context, db *gorm.DB, image models.Image) (models.Image, error)
	DeleteImage(ctx context.Context, db *gorm.DB, image models.Image) error
	GetImageById(ctx context.Context, db *gorm.DB, imageId string) (models.Image, error)
	FindAllImages(ctx context.Context, db *gorm.DB) ([]models.Image, error)
}

func NewImageRepository() ImageRepository {
	return &ImageRepositoryImpl{}
}

func (r *ImageRepositoryImpl) CreateImage(ctx context.Context, db *gorm.DB, image models.Image) (models.Image, error) {
	imageModel := models.Image{
		ID:        uuid.New().String(),
		ProductID: image.ProductID,
		URL:       image.URL,
	}

	err := db.WithContext(ctx).Create(&imageModel).Error
	helpers.PanicIfError(err)

	return imageModel, nil
}

func (r *ImageRepositoryImpl) UpdateImage(ctx context.Context, db *gorm.DB, image models.Image) (models.Image, error) {
	imageModel := models.Image{
		URL: image.URL,
	}

	err := db.WithContext(ctx).Model(&models.Image{}).Where("id = ?", image.ID).Updates(&imageModel).Error
	helpers.PanicIfError(err)

	return imageModel, nil
}

func (r *ImageRepositoryImpl) GetImageById(ctx context.Context, db *gorm.DB, imageId string) (models.Image, error) {
	var image models.Image
	err := db.WithContext(ctx).Model(&models.Image{}).Preload("Product").Where("id = ?", imageId).Take(&image).Error
	helpers.PanicIfError(err)

	return image, nil
}

func (r *ImageRepositoryImpl) DeleteImage(ctx context.Context, db *gorm.DB, image models.Image) error {
	err := db.WithContext(ctx).Model(&models.Image{}).Where("id = ?", image.ID).Delete(&image).Error
	helpers.PanicIfError(err)

	return nil
}

func (r *ImageRepositoryImpl) FindAllImages(ctx context.Context, db *gorm.DB) ([]models.Image, error) {
	var images []models.Image

	err := db.WithContext(ctx).Model(&models.Image{}).Preload("Product").Find(&images).Error
	helpers.PanicIfError(err)

	return images, nil
}
