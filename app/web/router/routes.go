package router

import (
	"github.com/gorilla/mux"
	"github.com/hiboedi/go-store-backend/app/middleware"
	"github.com/hiboedi/go-store-backend/app/web/controllers"
)

func RouterInit(userController controllers.UserController, storeController controllers.StoreController, billboardController controllers.BillboardController, categoryController controllers.CategoryController, colorController controllers.CategoryController) *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/api/login", userController.Login).Methods("POST")
	router.HandleFunc("/api/signup", userController.SignUp).Methods("POST")

	router.HandleFunc("/api", storeController.Create).Methods("POST")
	router.HandleFunc("/api", storeController.FindAll).Methods("GET")
	router.HandleFunc("/api/{storeId}", storeController.Update).Methods("PUT")
	router.HandleFunc("/api/{storeId}", storeController.Update).Methods("GET")
	router.HandleFunc("/api/{storeId}", storeController.Delete).Methods("DELETE")

	router.HandleFunc("/api/{storeId}/billboards", billboardController.Create).Methods("POST")
	router.HandleFunc("/api/{storeId}/billboards", billboardController.FindAll).Methods("GET")
	router.HandleFunc("/api/{storeId}/billboards/{billboardId}", billboardController.Update).Methods("PUT")
	router.HandleFunc("/api/{storeId}/billboards/{billboardId}", billboardController.Update).Methods("GET")
	router.HandleFunc("/api/{storeId}/billboards/{billboardId}", billboardController.Delete).Methods("DELETE")

	router.HandleFunc("/api/{storeId}/categories", categoryController.Create).Methods("POST")
	router.HandleFunc("/api/{storeId}/categories", categoryController.FindAll).Methods("GET")
	router.HandleFunc("/api/{storeId}/categories/{categoryId}", categoryController.Update).Methods("PUT")
	router.HandleFunc("/api/{storeId}/categories/{categoryId}", categoryController.Update).Methods("GET")
	router.HandleFunc("/api/{storeId}/categories/{categoryId}", categoryController.Delete).Methods("DELETE")

	router.HandleFunc("/api/{storeId}/colors", colorController.Create).Methods("POST")
	router.HandleFunc("/api/{storeId}/colors", colorController.FindAll).Methods("GET")
	router.HandleFunc("/api/{storeId}/colors/{colorId}", colorController.Update).Methods("PUT")
	router.HandleFunc("/api/{storeId}/colors/{colorId}", colorController.Update).Methods("GET")
	router.HandleFunc("/api/{storeId}/colors/{colorId}", colorController.Delete).Methods("DELETE")

	router.Use(middleware.RecoverMiddleware)

	return router
}
