package repositories

import (
	"context"

	"github.com/hiboedi/go-store-backend/app/helpers"
	"github.com/hiboedi/go-store-backend/app/web/models"
	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
}

type UserRepository interface {
	Create(ctx context.Context, db *gorm.DB, user models.User) (models.User, error)
	GetUserByEmail(ctx context.Context, db *gorm.DB, email string) (models.User, error)
	GetUserById(ctx context.Context, db *gorm.DB, userId string) (models.User, error)
}

func NewUserRepository() UserRepository {
	return &UserRepositoryImpl{}
}

func (r *UserRepositoryImpl) Create(ctx context.Context, db *gorm.DB, user models.User) (models.User, error) {

	err := db.WithContext(ctx).Create(&user).Error
	helpers.PanicIfError(err)

	return user, nil
}

func (r *UserRepositoryImpl) GetUserByEmail(ctx context.Context, db *gorm.DB, email string) (models.User, error) {
	var user models.User
	err := db.WithContext(ctx).Model(&models.User{}).Where("email = ?", email).Take(&user).Error

	helpers.PanicIfError(err)
	return user, nil
}

func (r *UserRepositoryImpl) GetUserById(ctx context.Context, db *gorm.DB, userId string) (models.User, error) {
	var user models.User
	err := db.WithContext(ctx).Model(&models.User{}).Where("id = ?", userId).Take(&user).Error

	helpers.PanicIfError(err)
	return user, nil
}
