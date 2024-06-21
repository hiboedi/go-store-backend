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

type ImageServiceImpl struct {
	ImageRepository repositories.ImageRepository
	DB              *gorm.DB
	Validate        *validator.Validate
}

type ImageService interface {
	Create(ctx context.Context, request models.ImageCreate) models.ImageResponseHiddenProduct
	Update(ctx context.Context, request models.ImageUpdate) models.ImageResponseHiddenProduct
	Delete(ctx context.Context, imageId string)
	FindById(ctx context.Context, imageId string) models.ImageResponse
	FindAll(ctx context.Context) []models.ImageResponse
}

func NewImageService(imageRepo repositories.ImageRepository, db *gorm.DB, validate *validator.Validate) ImageService {
	return &ImageServiceImpl{
		ImageRepository: imageRepo,
		DB:              db,
		Validate:        validate,
	}
}

func (s *ImageServiceImpl) Create(ctx context.Context, request models.ImageCreate) models.ImageResponseHiddenProduct {
	err := s.Validate.Struct(request)
	helpers.PanicIfError(err)

	tx := s.DB.Begin()
	defer helpers.CommitOrRollback(tx)

	image := models.Image{
		ProductID: request.ProductID,
		URL:       request.URL,
	}

	data, err := s.ImageRepository.CreateImage(ctx, tx, image)
	helpers.PanicIfError(err)

	return models.ToImageResponseHiddenProduct(data)
}

func (s *ImageServiceImpl) Update(ctx context.Context, request models.ImageUpdate) models.ImageResponseHiddenProduct {
	err := s.Validate.Struct(request)
	helpers.PanicIfError(err)

	tx := s.DB.Begin()
	defer helpers.CommitOrRollback(tx)

	image, err := s.ImageRepository.GetImageById(ctx, tx, request.ID)
	if err != nil {
		panic(exceptions.NewNotFoundError(err.Error()))
	}

	image.URL = request.URL

	data, err := s.ImageRepository.UpdateImage(ctx, tx, image)
	helpers.PanicIfError(err)

	return models.ToImageResponseHiddenProduct(data)
}

func (s *ImageServiceImpl) Delete(ctx context.Context, imageId string) {
	tx := s.DB.Begin()
	defer helpers.CommitOrRollback(tx)

	image, err := s.ImageRepository.GetImageById(ctx, tx, imageId)
	if err != nil {
		panic(exceptions.NewNotFoundError(err.Error()))
	}

	err = s.ImageRepository.DeleteImage(ctx, tx, image)
	helpers.PanicIfError(err)
}

func (s *ImageServiceImpl) FindAll(ctx context.Context) []models.ImageResponse {
	tx := s.DB.Begin()
	defer helpers.CommitOrRollback(tx)

	images, err := s.ImageRepository.FindAllImages(ctx, tx)
	helpers.PanicIfError(err)

	return models.ToImageResponses(images)
}

func (s *ImageServiceImpl) FindById(ctx context.Context, imageId string) models.ImageResponse {
	tx := s.DB.Begin()
	defer helpers.CommitOrRollback(tx)

	image, err := s.ImageRepository.GetImageById(ctx, tx, imageId)
	helpers.PanicIfError(err)

	return models.ToImageResponse(image)
}
