package repositories

import (
	"context"
	"math"

	"github.com/hiboedi/go-store-backend/app/helpers"
	"github.com/hiboedi/go-store-backend/app/web/models"
	"gorm.io/gorm"
)

type ProductRepositoryImpl struct {
}

type ProductRepository interface {
	CreateProduct(ctx context.Context, db *gorm.DB, product models.Product) (models.Product, error)
	UpdateProduct(ctx context.Context, db *gorm.DB, product models.Product) (models.Product, error)
	DeleteProduct(ctx context.Context, db *gorm.DB, product models.Product) error
	GetProductById(ctx context.Context, db *gorm.DB, productId string) (models.Product, error)
	FindAllProducts(ctx context.Context, db *gorm.DB, storeId string, pageParam int64) ([]models.Product, models.Pagination, error)
}

func NewProductRepository() ProductRepository {
	return &ProductRepositoryImpl{}
}

func (r *ProductRepositoryImpl) CreateProduct(ctx context.Context, db *gorm.DB, product models.Product) (models.Product, error) {

	err := db.WithContext(ctx).Save(&product).Error
	if err != nil {
		return models.Product{}, err
	}

	return product, nil
}

func (r *ProductRepositoryImpl) UpdateProduct(ctx context.Context, db *gorm.DB, product models.Product) (models.Product, error) {

	err := db.WithContext(ctx).Model(&models.Product{}).Where("id = ?", product.ID).Updates(&product).Error
	if err != nil {
		return models.Product{}, err
	}

	return product, nil
}

func (r *ProductRepositoryImpl) DeleteProduct(ctx context.Context, db *gorm.DB, product models.Product) error {
	var images []models.Image
	err := db.WithContext(ctx).Model(&models.Image{}).Where("product_id = ?", product.ID).Find(&images).Error
	helpers.PanicIfError(err)

	err = db.WithContext(ctx).Model(&models.Image{}).Where("product_id = ?", product.ID).Delete(&images).Error
	helpers.PanicIfError(err)

	err = db.WithContext(ctx).Model(&models.Product{}).Where("id = ?", product.ID).Delete(&product).Error

	helpers.PanicIfError(err)

	return nil
}

func (r *ProductRepositoryImpl) GetProductById(ctx context.Context, db *gorm.DB, productId string) (models.Product, error) {
	var product models.Product
	err := db.WithContext(ctx).Model(&models.Product{}).
		Preload("Store").
		Preload("Category").
		Preload("Size").
		Preload("Color").
		Preload("Images").
		Preload("CartItems").
		Where("id = ?", productId).
		Take(&product).
		Error
	helpers.PanicIfError(err)

	return product, nil
}

func (r *ProductRepositoryImpl) FindAllProducts(ctx context.Context, db *gorm.DB, storeId string, pageParam int64) ([]models.Product, models.Pagination, error) {
	var page int64 = 1
	if pageParam != 0 {
		page = pageParam
	}

	var pagination models.Pagination
	var products []models.Product
	offset, limit := r.Paging(db, page, storeId, &pagination)

	err := db.WithContext(ctx).Model(&models.Product{}).
		Where("store_id = ?", storeId).
		Preload("Store").
		Preload("Category").
		Preload("Size").
		Preload("Color").
		Preload("Images").
		Preload("CartItems").
		Offset(int(offset)).
		Limit(int(limit)).
		Find(&products).
		Error

	if err != nil {
		return nil, pagination, err
	}

	return products, pagination, nil
}

func (r *ProductRepositoryImpl) Paging(db *gorm.DB, page int64, storeId string, pagination *models.Pagination) (int64, int64) {
	var total int64
	var limit int64 = 5
	var offset int64

	db.Model(&models.Product{}).Where("store_id = ?", storeId).Count(&total)
	if page == 1 {
		offset = 0
	} else {
		offset = (page - 1) * limit
	}

	pagination.Total = uint64(total)
	pagination.PerPage = uint32(limit)
	pagination.CurrentPage = uint32(page)
	pagination.LastPage = uint32(math.Ceil(float64(total) / float64(limit)))

	return offset, limit
}
