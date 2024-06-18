package router

import (
	"github.com/gorilla/mux"
	"github.com/hiboedi/go-store-backend/app/middleware"
	"github.com/hiboedi/go-store-backend/app/web/controllers"
)

func RouterInit(userController controllers.UserController) *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/api/login", userController.Login).Methods("POST")
	router.HandleFunc("/api", userController.SignUp).Methods("POST")

	router.Use(middleware.RecoverMiddleware)

	return router
}
