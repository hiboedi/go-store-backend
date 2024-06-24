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

type BillboardServiceImpl struct {
	BillboardRepository repositories.BillboardRepository
	DB                  *gorm.DB
	Validate            *validator.Validate
}

type BillboardService interface {
	Create(ctx context.Context, request models.BillboardCreate) models.BillboardResponseHiddenStore
	Update(ctx context.Context, request models.BillboardUpdate, billboardId string) models.BillboardResponseHiddenStore
	Delete(ctx context.Context, billboardId string)
	FindById(ctx context.Context, billboardId string) models.BillboardResponse
	FindAll(ctx context.Context) []models.BillboardResponse
}

func NewBillboardService(billboardRepo repositories.BillboardRepository, db *gorm.DB, validate *validator.Validate) BillboardService {
	return &BillboardServiceImpl{
		BillboardRepository: billboardRepo,
		DB:                  db,
		Validate:            validate,
	}
}

func (s *BillboardServiceImpl) Create(ctx context.Context, request models.BillboardCreate) models.BillboardResponseHiddenStore {
	err := s.Validate.Struct(request)
	helpers.PanicIfError(err)

	tx := s.DB.Begin()
	defer helpers.CommitOrRollback(tx)

	billboard := models.Billboard{
		Label:    request.Label,
		ImageURL: request.ImageURL,
		StoreID:  request.StoreID,
	}

	data, err := s.BillboardRepository.CreateBillboard(ctx, tx, billboard)
	helpers.PanicIfError(err)

	return models.ToBillboardResponseHiddenStore(data)
}

func (s *BillboardServiceImpl) Update(ctx context.Context, request models.BillboardUpdate, billboardId string) models.BillboardResponseHiddenStore {
	err := s.Validate.Struct(request)
	helpers.PanicIfError(err)

	tx := s.DB.Begin()
	defer helpers.CommitOrRollback(tx)

	billboard, err := s.BillboardRepository.GetBillboardById(ctx, tx, billboardId)
	if err != nil {
		panic(exceptions.NewNotFoundError(err.Error()))
	}

	billboard.Label = request.Label
	billboard.ImageURL = request.ImageURL

	data, err := s.BillboardRepository.UpdateBillboard(ctx, tx, billboard)
	helpers.PanicIfError(err)

	return models.ToBillboardResponseHiddenStore(data)
}

func (s *BillboardServiceImpl) Delete(ctx context.Context, billboardId string) {
	tx := s.DB.Begin()
	defer helpers.CommitOrRollback(tx)

	billboard, err := s.BillboardRepository.GetBillboardById(ctx, tx, billboardId)
	if err != nil {
		panic(exceptions.NewNotFoundError(err.Error()))
	}

	err = s.BillboardRepository.DeleteBillboard(ctx, tx, billboard)
	helpers.PanicIfError(err)
}

func (s *BillboardServiceImpl) FindAll(ctx context.Context) []models.BillboardResponse {
	tx := s.DB.Begin()
	defer helpers.CommitOrRollback(tx)

	billboards, err := s.BillboardRepository.FindAllBillboards(ctx, tx)
	helpers.PanicIfError(err)
	return models.ToBillboardResponses(billboards)
}

func (s *BillboardServiceImpl) FindById(ctx context.Context, billboardId string) models.BillboardResponse {
	tx := s.DB.Begin()
	defer helpers.CommitOrRollback(tx)

	billboard, err := s.BillboardRepository.GetBillboardById(ctx, tx, billboardId)
	helpers.PanicIfError(err)
	return models.ToBillboardReponse(billboard)
}
