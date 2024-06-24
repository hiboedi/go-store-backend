package repositories

import (
	"context"

	"github.com/google/uuid"
	"github.com/hiboedi/go-store-backend/app/helpers"
	"github.com/hiboedi/go-store-backend/app/web/models"
	"gorm.io/gorm"
)

type SizeRepositoryImpl struct {
}

type SizeRepository interface {
	CreateSize(ctx context.Context, db *gorm.DB, size models.Size) (models.Size, error)
	UpdateSize(ctx context.Context, db *gorm.DB, size models.Size) (models.Size, error)
	DeleteSize(ctx context.Context, db *gorm.DB, size models.Size) error
	GetSizeById(ctx context.Context, db *gorm.DB, sizeId string) (models.Size, error)
	FindAllSizes(ctx context.Context, db *gorm.DB) ([]models.Size, error)
}

func NewSizeRepository() SizeRepository {
	return &SizeRepositoryImpl{}
}

func (r *SizeRepositoryImpl) CreateSize(ctx context.Context, db *gorm.DB, size models.Size) (models.Size, error) {
	sizeModel := models.Size{
		ID:      uuid.New().String(),
		StoreID: size.StoreID,
		Name:    size.Name,
		Value:   size.Value,
	}

	err := db.WithContext(ctx).Create(&sizeModel).Error
	helpers.PanicIfError(err)

	return sizeModel, nil
}

func (r *SizeRepositoryImpl) UpdateSize(ctx context.Context, db *gorm.DB, size models.Size) (models.Size, error) {
	sizeModel := models.Size{
		ID:        size.ID,
		StoreID:   size.StoreID,
		Name:      size.Name,
		Value:     size.Value,
		CreatedAt: size.CreatedAt,
		UpdatedAt: size.UpdatedAt,
	}

	err := db.WithContext(ctx).Model(&models.Size{}).Where("id = ?", size.ID).Updates(&sizeModel).Error
	helpers.PanicIfError(err)

	return sizeModel, nil
}

func (r *SizeRepositoryImpl) GetSizeById(ctx context.Context, db *gorm.DB, sizeId string) (models.Size, error) {
	var size models.Size
	err := db.WithContext(ctx).Model(&models.Size{}).Preload("Store", func(db *gorm.DB) *gorm.DB {
		return db.Preload("User", func(db *gorm.DB) *gorm.DB {
			return db.Omit("password", "email", "store")
		})
	}).Where("id = ?", sizeId).Take(&size).Error
	helpers.PanicIfError(err)

	return size, nil
}

func (r *SizeRepositoryImpl) DeleteSize(ctx context.Context, db *gorm.DB, size models.Size) error {
	err := db.WithContext(ctx).Model(&models.Size{}).Where("id = ?", size.ID).Delete(&size).Error
	helpers.PanicIfError(err)

	return nil
}

func (r *SizeRepositoryImpl) FindAllSizes(ctx context.Context, db *gorm.DB) ([]models.Size, error) {
	var sizes []models.Size

	err := db.WithContext(ctx).Model(&models.Size{}).Preload("Store", func(db *gorm.DB) *gorm.DB {
		return db.Preload("User", func(db *gorm.DB) *gorm.DB {
			return db.Omit("password", "email", "store")
		})
	}).Find(&sizes).Error
	helpers.PanicIfError(err)

	return sizes, nil
}
