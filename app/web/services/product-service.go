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

type ProductServiceImpl struct {
	ProductRepository repositories.ProductRepository
	DB                *gorm.DB
	Validate          *validator.Validate
}

type ProductService interface {
	Create(ctx context.Context, request models.ProductCreate) models.ProductResponseHiddenStore
	Update(ctx context.Context, request models.ProductUpdate, productId string) models.ProductResponseHiddenStore
	Delete(ctx context.Context, productId string)
	FindById(ctx context.Context, productId string) models.ProductResponse
	FindAll(ctx context.Context, storeId string, pageParam int64) models.ProductPagination
}

func NewProductService(productRepo repositories.ProductRepository, db *gorm.DB, validate *validator.Validate) ProductService {
	return &ProductServiceImpl{
		ProductRepository: productRepo,
		DB:                db,
		Validate:          validate,
	}
}

func (s *ProductServiceImpl) Create(ctx context.Context, request models.ProductCreate) models.ProductResponseHiddenStore {
	err := s.Validate.Struct(request)
	helpers.PanicIfError(err)

	tx := s.DB.Begin()
	defer helpers.CommitOrRollback(tx)

	product := models.Product{
		StoreID:    request.StoreID,
		CategoryID: request.CategoryID,
		Name:       request.Name,
		Stock:      request.Stock,
		Price:      request.Price,
		IsFeatured: request.IsFeatured,
		IsArchived: request.IsArchived,
		SizeID:     request.SizeID,
		ColorID:    request.ColorID,
		Images:     request.Images,
		OrderItems: request.OrderItems,
	}

	data, err := s.ProductRepository.CreateProduct(ctx, tx, product)
	helpers.PanicIfError(err)

	return models.ToProductResponseHiddenStore(data)
}

func (s *ProductServiceImpl) Update(ctx context.Context, request models.ProductUpdate, productId string) models.ProductResponseHiddenStore {
	err := s.Validate.Struct(request)
	helpers.PanicIfError(err)

	tx := s.DB.Begin()
	defer helpers.CommitOrRollback(tx)

	product, err := s.ProductRepository.GetProductById(ctx, tx, productId)
	if err != nil {
		panic(exceptions.NewNotFoundError(err.Error()))
	}

	product.CategoryID = request.CategoryID
	product.StoreID = request.StoreID
	product.Name = request.Name
	product.Price = request.Price
	product.IsFeatured = request.IsFeatured
	product.IsArchived = request.IsArchived
	product.SizeID = request.SizeID
	product.ColorID = request.ColorID
	product.Stock = request.Stock
	product.Images = request.Images
	product.OrderItems = request.OrderItems
	product.Images = request.Images

	data, err := s.ProductRepository.UpdateProduct(ctx, tx, product)
	helpers.PanicIfError(err)

	return models.ToProductResponseHiddenStore(data)
}

func (s *ProductServiceImpl) Delete(ctx context.Context, productId string) {
	tx := s.DB.Begin()
	defer helpers.CommitOrRollback(tx)

	product, err := s.ProductRepository.GetProductById(ctx, tx, productId)
	if err != nil {
		panic(exceptions.NewNotFoundError(err.Error()))
	}

	err = s.ProductRepository.DeleteProduct(ctx, tx, product)
	helpers.PanicIfError(err)
}

func (s *ProductServiceImpl) FindAll(ctx context.Context, storeId string, pageParams int64) models.ProductPagination {
	tx := s.DB.Begin()
	defer helpers.CommitOrRollback(tx)

	products, pagination, err := s.ProductRepository.FindAllProducts(ctx, tx, storeId, pageParams)
	helpers.PanicIfError(err)
	return models.ToProductResponses(products, pagination)
}

func (s *ProductServiceImpl) FindById(ctx context.Context, productId string) models.ProductResponse {
	tx := s.DB.Begin()
	defer helpers.CommitOrRollback(tx)

	product, err := s.ProductRepository.GetProductById(ctx, tx, productId)
	helpers.PanicIfError(err)
	return models.ToProductResponse(product)
}
