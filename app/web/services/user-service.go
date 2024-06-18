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

type UserServiceimpl struct {
	UserRepo repositories.UserRepository
	DB       *gorm.DB
	Validate *validator.Validate
}

type UserService interface {
	Create(ctx context.Context, request models.UserCreate) models.UserResponse
	Login(ctx context.Context, requestLogin models.UserLogin) (models.UserLoginResponse, bool)
}

func NewUserService(userRepo repositories.UserRepository, db *gorm.DB, validate *validator.Validate) UserService {
	return &UserServiceimpl{
		UserRepo: userRepo,
		DB:       db,
		Validate: validate,
	}
}

func (s *UserServiceimpl) Create(ctx context.Context, request models.UserCreate) models.UserResponse {
	err := s.Validate.Struct(request)
	helpers.PanicIfError(err)

	tx := s.DB.Begin()
	defer helpers.CommitOrRollback(tx)

	user := models.User{
		Name:     request.Name,
		Email:    request.Email,
		Password: request.Password,
	}

	data, err := s.UserRepo.Create(ctx, tx, user)
	helpers.PanicIfError(err)

	return models.ToUserReponse(data)
}

func (s *UserServiceimpl) Login(ctx context.Context, requestLogin models.UserLogin) (models.UserLoginResponse, bool) {
	err := s.Validate.Struct(requestLogin)
	helpers.PanicIfError(err)

	tx := s.DB.Begin()
	defer helpers.CommitOrRollback(tx)

	user, err := s.UserRepo.GetUserByEmail(ctx, tx, requestLogin.Email)
	if err != nil {
		panic(exceptions.NewNotFoundError(err.Error()))
	}
	passwordSync := helpers.ComparePassword(requestLogin.Password, user.Password)
	token, _ := helpers.CreateToken(user.ID)

	if !passwordSync {
		return models.UserLoginResponse{}, false
	} else {

		userLoginResponse := models.UserLoginResponse{
			ID:    user.ID,
			Name:  user.Name,
			Email: user.Email,
			Token: token,
		}
		return userLoginResponse, true
	}
}
