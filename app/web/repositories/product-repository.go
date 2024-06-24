package repositories

import (
	"context"

	"github.com/google/uuid"
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
	FindAllProducts(ctx context.Context, db *gorm.DB) ([]models.Product, error)
}

func NewProductRepository() ProductRepository {
	return &ProductRepositoryImpl{}
}

func (r *ProductRepositoryImpl) CreateProduct(ctx context.Context, db *gorm.DB, product models.Product) (models.Product, error) {
	productModel := models.Product{
		ID:         uuid.New().String(),
		StoreID:    product.StoreID,
		CategoryID: product.CategoryID,
		Name:       product.Name,
		Price:      product.Price,
		Stock:      product.Stock,
		IsFeatured: product.IsFeatured,
		IsArchived: product.IsArchived,
		SizeID:     product.SizeID,
		ColorID:    product.ColorID,
		Images:     product.Images,
		OrderItems: product.OrderItems,
	}

	err := db.WithContext(ctx).Create(&productModel).Error
	helpers.PanicIfError(err)

	return productModel, nil
}

func (r *ProductRepositoryImpl) UpdateProduct(ctx context.Context, db *gorm.DB, product models.Product) (models.Product, error) {
	productModel := models.Product{
		ID:         product.ID,
		StoreID:    product.StoreID,
		CategoryID: product.CategoryID,
		Name:       product.Name,
		Price:      product.Price,
		Stock:      product.Stock,
		IsFeatured: product.IsFeatured,
		IsArchived: product.IsArchived,
		SizeID:     product.SizeID,
		ColorID:    product.ColorID,
		Images:     product.Images,
		OrderItems: product.OrderItems,
	}

	err := db.WithContext(ctx).Model(&models.Product{}).Where("id = ?", product.ID).Updates(&productModel).Error
	helpers.PanicIfError(err)

	return productModel, nil
}

func (r *ProductRepositoryImpl) DeleteProduct(ctx context.Context, db *gorm.DB, product models.Product) error {
	err := db.WithContext(ctx).Model(&models.Product{}).Where("id = ?", product.ID).Delete(&product).Error
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
		Preload("OrderItems").
		Where("id = ?", productId).
		Take(&product).
		Error
	helpers.PanicIfError(err)

	return product, nil
}

func (r *ProductRepositoryImpl) FindAllProducts(ctx context.Context, db *gorm.DB) ([]models.Product, error) {
	var products []models.Product
	err := db.WithContext(ctx).Model(&models.Product{}).
		Preload("Store").
		Preload("Category").
		Preload("Size").
		Preload("Color").
		Preload("Images").
		Preload("OrderItems").
		Find(&products).
		Error
	helpers.PanicIfError(err)

	return products, nil
}
