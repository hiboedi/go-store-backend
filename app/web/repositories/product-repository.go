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
	FindAllProducts(ctx context.Context, db *gorm.DB, storeId string) ([]models.Product, error)
}

func NewProductRepository() ProductRepository {
	return &ProductRepositoryImpl{}
}

func (r *ProductRepositoryImpl) CreateProduct(ctx context.Context, db *gorm.DB, product models.Product) (models.Product, error) {
	productId := uuid.New().String()

	productModel := models.Product{
		ID:         productId,
		StoreID:    product.StoreID,
		CategoryID: product.CategoryID,
		Name:       product.Name,
		Price:      product.Price,
		Stock:      product.Stock,
		IsFeatured: product.IsFeatured,
		IsArchived: product.IsArchived,
		SizeID:     product.SizeID,
		ColorID:    product.ColorID,
	}

	err := db.WithContext(ctx).Save(&productModel).Error
	if err != nil {
		return models.Product{}, err
	}

	var images []models.Image
	for _, image := range product.Images {
		image.ID = uuid.New().String()
		image.ProductID = productId
		if err := db.WithContext(ctx).Create(&image).Error; err != nil {
			return models.Product{}, err
		}
		images = append(images, image)
	}

	productModel.Images = images

	err = db.WithContext(ctx).Save(&productModel).Error
	if err != nil {
		return models.Product{}, err
	}

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
		CreatedAt:  product.CreatedAt,
		UpdatedAt:  product.UpdatedAt,
	}

	err := db.WithContext(ctx).Model(&models.Product{}).Where("id = ?", product.ID).Updates(&productModel).Error
	if err != nil {
		return models.Product{}, err
	}

	var images []models.Image
	err = db.WithContext(ctx).Model(&models.Image{}).Where("product_id = ?", product.ID).Find(&images).Error
	helpers.PanicIfError(err)

	updatedImages := []models.Image{}
	for _, image := range images {
		for _, productImage := range product.Images {
			productImage.ID = image.ID
			productImage.ProductID = image.ProductID
			productImage.CreatedAt = image.CreatedAt
			productImage.UpdatedAt = image.UpdatedAt

			if err := db.WithContext(ctx).Model(&models.Image{}).Where("id = ?", image.ID).Updates(&productImage).Error; err != nil {
				return models.Product{}, err
			}
			updatedImages = append(updatedImages, productImage)

		}
	}

	productModel.Images = updatedImages

	err = db.WithContext(ctx).Model(&models.Product{}).Where("id = ?", product.ID).Updates(&productModel).Error
	if err != nil {
		return models.Product{}, err
	}

	return productModel, nil
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
		Preload("OrderItems").
		Where("id = ?", productId).
		Take(&product).
		Error
	helpers.PanicIfError(err)

	return product, nil
}

func (r *ProductRepositoryImpl) FindAllProducts(ctx context.Context, db *gorm.DB, storeId string) ([]models.Product, error) {
	var products []models.Product
	err := db.WithContext(ctx).Model(&models.Product{}).Where("store_id = ?", storeId).
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
