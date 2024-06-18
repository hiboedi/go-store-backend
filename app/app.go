package app

import (
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/hiboedi/go-store-backend/app/database"
	"github.com/hiboedi/go-store-backend/app/helpers"
	"github.com/hiboedi/go-store-backend/app/middleware"
	"github.com/hiboedi/go-store-backend/app/web/controllers"
	"github.com/hiboedi/go-store-backend/app/web/repositories"
	"github.com/hiboedi/go-store-backend/app/web/router"
	"github.com/hiboedi/go-store-backend/app/web/services"
)

func Start() {
	db := database.InitializeDB()
	validate := validator.New()

	userRepository := repositories.NewUserRepository()
	userService := services.NewUserService(userRepository, db, validate)
	userController := controllers.NewUserController(userService)

	router := router.RouterInit(userController)
	database.DBMigrate()

	server := http.Server{
		Addr:    "localhost:8000",
		Handler: middleware.NewAuthMiddleware(router),
	}
	fmt.Println("starting on port :8000")

	err := server.ListenAndServe()
	helpers.PanicIfError(err)
}
