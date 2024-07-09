package services

import (
	"context"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/hiboedi/go-store-backend/app/auth"
	"github.com/hiboedi/go-store-backend/app/exceptions"
	"github.com/hiboedi/go-store-backend/app/helpers"
	"github.com/hiboedi/go-store-backend/app/web/models"
	"github.com/hiboedi/go-store-backend/app/web/repositories"
	"gorm.io/gorm"
)

type UserServiceimpl struct {
	UserRepo repositories.UserRepository
	CartRepo repositories.CartRepository
	DB       *gorm.DB
	Validate *validator.Validate
}

type UserService interface {
	Create(ctx context.Context, request models.UserCreate) models.UserResponse
	Login(ctx context.Context, requestLogin models.UserLogin) (models.UserLoginResponse, bool)
}

func NewUserService(userRepo repositories.UserRepository, db *gorm.DB, cartRepo repositories.CartRepository, validate *validator.Validate) UserService {
	return &UserServiceimpl{
		UserRepo: userRepo,
		CartRepo: cartRepo,
		DB:       db,
		Validate: validate,
	}
}

func (s *UserServiceimpl) Create(ctx context.Context, request models.UserCreate) models.UserResponse {
	err := s.Validate.Struct(request)
	helpers.PanicIfError(err)

	tx := s.DB.Begin()
	defer helpers.CommitOrRollback(tx)

	hashPassword, _ := helpers.MakePassword(request.Password)

	user := models.User{
		ID:       uuid.New().String(),
		Name:     request.Name,
		Email:    request.Email,
		Password: hashPassword,
		Phone:    request.Phone,
	}

	data, err := s.UserRepo.Create(ctx, tx, user)
	helpers.PanicIfError(err)

	cart := models.Cart{
		ID:     uuid.New().String(),
		UserID: user.ID,
	}

	_, err = s.CartRepo.CreateCart(ctx, tx, cart)
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
	token, _ := auth.CreateToken(user.ID)

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
