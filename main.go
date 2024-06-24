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

	// Initialize repositories
	userRepository := repositories.NewUserRepository()
	storeRepository := repositories.NewStoreRepository()
	billboardRepository := repositories.NewBillboardRepository()
	categoryRepository := repositories.NewCategoryRepository()
	colorRepository := repositories.NewColorRepository()
	sizeRepository := repositories.NewSizeRepository()
	productRepository := repositories.NewProductRepository()
	imageRepository := repositories.NewImageRepository()
	orderRepository := repositories.NewOrderRepository()
	orderItemRepository := repositories.NewOrderItemRepository()

	// Initialize services
	userService := services.NewUserService(userRepository, db, validate)
	storeService := services.NewStoreService(storeRepository, db, validate)
	billboardService := services.NewBillboardService(billboardRepository, db, validate)
	categoryService := services.NewCategoryService(categoryRepository, db, validate)
	colorService := services.NewColorService(colorRepository, db, validate)
	sizeService := services.NewSizeService(sizeRepository, db, validate)
	imageService := services.NewImageService(imageRepository, db, validate)
	orderItemService := services.NewOrderItemService(orderItemRepository, db, validate)
	productService := services.NewProductService(productRepository, imageService, db, validate)
	orderService := services.NewOrderService(orderRepository, orderItemService, db, validate)

	// Initialize controllers
	userController := controllers.NewUserController(userService)
	storeController := controllers.NewStoreController(storeService)
	billboardController := controllers.NewBillboardController(billboardService)
	categoryController := controllers.NewCategoryController(categoryService)
	colorController := controllers.NewColorController(colorService)
	sizeController := controllers.NewSizeController(sizeService)
	productController := controllers.NewProductController(productService)
	orderController := controllers.NewOrderController(orderService)

	// Initialize router
	r := router.RouterInit(
		userController,
		storeController,
		billboardController,
		categoryController,
		colorController,
		sizeController,
		productController,
		orderController,
	)

	// Perform database migration
	database.DBMigrate()

	// Initialize middleware and server
	authRouter := middleware.NewAuthMiddleware(r)

	server := http.Server{
		Addr:    "localhost:8000",
		Handler: authRouter,
	}

	// Start the server
	fmt.Println("Starting server on port :8000")
	err := server.ListenAndServe()
	helpers.PanicIfError(err)
}
