package repositories

import (
	"context"

	"github.com/google/uuid"
	"github.com/hiboedi/go-store-backend/app/helpers"
	"github.com/hiboedi/go-store-backend/app/web/models"
	"gorm.io/gorm"
)

type BillboardRepositoryImpl struct {
}

type BillboardRepository interface {
	CreateBillboard(ctx context.Context, db *gorm.DB, billboard models.Billboard) (models.Billboard, error)
	UpdateBillboard(ctx context.Context, db *gorm.DB, billboard models.Billboard) (models.Billboard, error)
	DeleteBillboard(ctx context.Context, db *gorm.DB, billboard models.Billboard) error
	GetBillboardById(ctx context.Context, db *gorm.DB, billboardId string) (models.Billboard, error)
	FindAllBillboards(ctx context.Context, db *gorm.DB) ([]models.Billboard, error)
}

func NewBillboardRepository() BillboardRepository {
	return &BillboardRepositoryImpl{}
}

func (r *BillboardRepositoryImpl) CreateBillboard(ctx context.Context, db *gorm.DB, billboard models.Billboard) (models.Billboard, error) {

	billboardModel := models.Billboard{
		ID:       uuid.New().String(),
		Label:    billboard.Label,
		StoreID:  billboard.StoreID,
		ImageURL: billboard.ImageURL,
	}

	err := db.WithContext(ctx).Create(&billboardModel).Error
	helpers.PanicIfError(err)

	return billboardModel, nil
}

func (r *BillboardRepositoryImpl) UpdateBillboard(ctx context.Context, db *gorm.DB, billboard models.Billboard) (models.Billboard, error) {

	billboardModel := models.Billboard{
		ID:        billboard.ID,
		Label:     billboard.Label,
		StoreID:   billboard.StoreID,
		ImageURL:  billboard.ImageURL,
		CreatedAt: billboard.CreatedAt,
		UpdatedAt: billboard.UpdatedAt,
	}

	err := db.WithContext(ctx).Model(&models.Billboard{}).Where("id = ?", billboard.ID).Updates(&billboardModel).Error
	helpers.PanicIfError(err)

	return billboardModel, nil
}

func (r *BillboardRepositoryImpl) GetBillboardById(ctx context.Context, db *gorm.DB, billboardId string) (models.Billboard, error) {
	var billboard models.Billboard
	err := db.WithContext(ctx).Model(&models.Billboard{}).Preload("Store", func(db *gorm.DB) *gorm.DB {
		return db.Preload("User", func(db *gorm.DB) *gorm.DB {
			return db.Omit("password", "email", "store")
		})
	}).Where("id = ?", billboardId).Take(&billboard).Error
	helpers.PanicIfError(err)

	return billboard, nil
}

func (r *BillboardRepositoryImpl) DeleteBillboard(ctx context.Context, db *gorm.DB, billboard models.Billboard) error {

	err := db.WithContext(ctx).Model(&models.Billboard{}).Where("id = ?", billboard.ID).Delete(&billboard).Error
	helpers.PanicIfError(err)

	return nil
}

func (r *BillboardRepositoryImpl) FindAllBillboards(ctx context.Context, db *gorm.DB) ([]models.Billboard, error) {
	var billboards []models.Billboard

	err := db.WithContext(ctx).Model(&models.Billboard{}).Preload("Store", func(db *gorm.DB) *gorm.DB {
		return db.Preload("User", func(db *gorm.DB) *gorm.DB {
			return db.Omit("password", "email", "store")
		})
	}).Find(&billboards).Error
	helpers.PanicIfError(err)

	return billboards, nil
}
