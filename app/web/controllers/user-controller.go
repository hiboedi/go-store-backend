package controllers

import (
	"net/http"

	"github.com/hiboedi/go-store-backend/app/helpers"
	"github.com/hiboedi/go-store-backend/app/web"
	"github.com/hiboedi/go-store-backend/app/web/models"
	"github.com/hiboedi/go-store-backend/app/web/services"
)

type UserControllerImpl struct {
	UserService services.UserService
}

type UserController interface {
	SignUp(w http.ResponseWriter, r *http.Request)
	Login(w http.ResponseWriter, r *http.Request)
}

func NewUserController(userService services.UserService) UserController {
	return &UserControllerImpl{
		UserService: userService,
	}
}

func (c *UserControllerImpl) SignUp(w http.ResponseWriter, r *http.Request) {
	userSignUp := models.UserCreate{}
	helpers.ToRequestBody(r, &userSignUp)

	userResponse := c.UserService.Create(r.Context(), userSignUp)
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   userResponse,
	}

	helpers.WriteResponseBody(w, webResponse)
	http.Redirect(w, r, "/api/users/login", http.StatusOK)
}

func (c *UserControllerImpl) Login(w http.ResponseWriter, r *http.Request) {
	userLogin := models.UserLogin{}
	helpers.ToRequestBody(r, &userLogin)

	userResponse, loggedIn := c.UserService.Login(r.Context(), userLogin)
	if loggedIn {
		helpers.SetUserCookie(w, r, userResponse)
		webResponse := web.WebResponse{
			Code:   http.StatusOK,
			Status: "Ok",
			Data:   userResponse,
		}
		helpers.WriteResponseBody(w, webResponse)
	} else {
		webResponse := web.WebResponse{
			Code:   http.StatusUnauthorized,
			Status: "Unauthorized",
		}
		helpers.WriteResponseBody(w, webResponse)
	}
}
