package repositories

import (
	"context"

	"github.com/hiboedi/go-store-backend/app/helpers"
	"github.com/hiboedi/go-store-backend/app/web/models"
	"gorm.io/gorm"
)

type StoreRepositoryImpl struct {
}

type StoreRepository interface {
	CreateStore(ctx context.Context, db *gorm.DB, store models.Store) (models.Store, error)
	UpdateStore(ctx context.Context, db *gorm.DB, store models.Store) (models.Store, error)
	DeleteStore(ctx context.Context, db *gorm.DB, store models.Store) error
	GetStoreById(ctx context.Context, db *gorm.DB, storeId string) (models.Store, error)
	FindAllStore(ctx context.Context, db *gorm.DB, userId string) ([]models.Store, error)
}

func NewStoreRepository() StoreRepository {
	return &StoreRepositoryImpl{}
}

func (r *StoreRepositoryImpl) CreateStore(ctx context.Context, db *gorm.DB, store models.Store) (models.Store, error) {

	err := db.WithContext(ctx).Create(&store).Error
	helpers.PanicIfError(err)

	return store, nil
}

func (r *StoreRepositoryImpl) UpdateStore(ctx context.Context, db *gorm.DB, store models.Store) (models.Store, error) {

	err := db.WithContext(ctx).Model(&models.Store{}).Where("id = ?", store.ID).Updates(&store).Error

	helpers.PanicIfError(err)

	return store, nil
}

func (r *StoreRepositoryImpl) GetStoreById(ctx context.Context, db *gorm.DB, storeId string) (models.Store, error) {
	var store models.Store
	err := db.WithContext(ctx).Model(&models.Store{}).Preload("Colors").Preload("Orders").Preload("Products", func(db *gorm.DB) *gorm.DB {
		return db.Preload("Images")
	}).Preload("Sizes").Preload("Categories").Preload("Billboards").Preload("User", func(db *gorm.DB) *gorm.DB {
		return db.Omit("password")
	}).Where("id = ?", storeId).Take(&store).Error
	helpers.PanicIfError(err)

	return store, nil
}

func (r *StoreRepositoryImpl) DeleteStore(ctx context.Context, db *gorm.DB, store models.Store) error {

	err := db.WithContext(ctx).Model(&models.Store{}).Where("id = ?", store.ID).Delete(&store).Error
	helpers.PanicIfError(err)

	return nil
}

func (r *StoreRepositoryImpl) FindAllStore(ctx context.Context, db *gorm.DB, userId string) ([]models.Store, error) {
	var stores []models.Store

	err := db.WithContext(ctx).Model(&models.Store{}).Where("user_id = ?", userId).Preload("Colors").Preload("Orders").Preload("Products", func(db *gorm.DB) *gorm.DB {
		return db.Preload("Images")
	}).Preload("Sizes").Preload("Categories").Preload("Billboards").Preload("User", func(db *gorm.DB) *gorm.DB {
		return db.Omit("password")
	}).Find(&stores).Error
	helpers.PanicIfError(err)

	return stores, nil
}
