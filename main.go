package main

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

func main() {
	db := database.InitializeDB()
	validate := validator.New()

	userRepository := repositories.NewUserRepository()
	storeRepository := repositories.NewStoreRepository()
	billboardRepository := repositories.NewBillboardRepository()
	categoryRepository := repositories.NewCategoryRepository()
	colorRepository := repositories.NewColorRepository()
	sizeRepository := repositories.NewSizeRepository()
	productRepository := repositories.NewProductRepository()
	imageRepository := repositories.NewImageRepository()
	orderRepository := repositories.NewOrderRepository()
	cartRepository := repositories.NewCartRepository()

	userService := services.NewUserService(userRepository, db, cartRepository, validate)
	storeService := services.NewStoreService(storeRepository, db, validate)
	billboardService := services.NewBillboardService(billboardRepository, db, validate)
	categoryService := services.NewCategoryService(categoryRepository, db, validate)
	colorService := services.NewColorService(colorRepository, db, validate)
	sizeService := services.NewSizeService(sizeRepository, db, validate)
	productService := services.NewProductService(productRepository, imageRepository, db, validate)
	orderService := services.NewOrderService(orderRepository, db)

	userController := controllers.NewUserController(userService)
	storeController := controllers.NewStoreController(storeService)
	billboardController := controllers.NewBillboardController(billboardService)
	categoryController := controllers.NewCategoryController(categoryService)
	colorController := controllers.NewColorController(colorService)
	sizeController := controllers.NewSizeController(sizeService)
	productController := controllers.NewProductController(productService)
	orderController := controllers.NewOrderController(orderService)

	router := router.RouterInit(
		userController,
		storeController,
		billboardController,
		categoryController,
		colorController,
		sizeController,
		productController,
		orderController,
	)

	database.DBMigrate()

	authRouter := middleware.AuthMiddleware(router)

	server := http.Server{
		Addr:    "localhost:9000",
		Handler: authRouter,
	}

	fmt.Println("Starting server on port :9000")
	err := server.ListenAndServe()
	helpers.PanicIfError(err)
}
