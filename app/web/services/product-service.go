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
	ImageService      ImageService
	DB                *gorm.DB
	Validate          *validator.Validate
}

type ProductService interface {
	Create(ctx context.Context, request models.ProductCreate) models.ProductResponse
	Update(ctx context.Context, request models.ProductUpdate) models.ProductResponse
	Delete(ctx context.Context, productId string)
	FindById(ctx context.Context, productId string) models.ProductResponse
	FindAll(ctx context.Context) []models.ProductResponse
}

func NewProductService(productRepo repositories.ProductRepository, imageService ImageService, db *gorm.DB, validate *validator.Validate) ProductService {
	return &ProductServiceImpl{
		ProductRepository: productRepo,
		ImageService:      imageService,
		DB:                db,
		Validate:          validate,
	}
}

func (s *ProductServiceImpl) Create(ctx context.Context, request models.ProductCreate) models.ProductResponse {
	err := s.Validate.Struct(request)
	helpers.PanicIfError(err)

	tx := s.DB.Begin()
	defer helpers.CommitOrRollback(tx)

	product := models.Product{
		StoreID:    request.StoreID,
		CategoryID: request.CategoryID,
		Name:       request.Name,
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

	return models.ToProductResponse(data)
}

func (s *ProductServiceImpl) Update(ctx context.Context, request models.ProductUpdate) models.ProductResponse {
	err := s.Validate.Struct(request)
	helpers.PanicIfError(err)

	tx := s.DB.Begin()
	defer helpers.CommitOrRollback(tx)

	product, err := s.ProductRepository.GetProductById(ctx, tx, request.ID)
	if err != nil {
		panic(exceptions.NewNotFoundError(err.Error()))
	}

	product.StoreID = request.StoreID
	product.CategoryID = request.CategoryID
	product.Name = request.Name
	product.Price = request.Price
	product.IsFeatured = request.IsFeatured
	product.IsArchived = request.IsArchived
	product.SizeID = request.SizeID
	product.ColorID = request.ColorID
	product.Images = request.Images
	product.OrderItems = request.OrderItems

	data, err := s.ProductRepository.UpdateProduct(ctx, tx, product)
	helpers.PanicIfError(err)

	return models.ToProductResponse(data)
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

func (s *ProductServiceImpl) FindAll(ctx context.Context) []models.ProductResponse {
	tx := s.DB.Begin()
	defer helpers.CommitOrRollback(tx)

	products, err := s.ProductRepository.FindAllProducts(ctx, tx)
	helpers.PanicIfError(err)
	return models.ToProductResponses(products)
}

func (s *ProductServiceImpl) FindById(ctx context.Context, productId string) models.ProductResponse {
	tx := s.DB.Begin()
	defer helpers.CommitOrRollback(tx)

	product, err := s.ProductRepository.GetProductById(ctx, tx, productId)
	helpers.PanicIfError(err)
	return models.ToProductResponse(product)
}
