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

// SignUp godoc
// @Summary Sign up a new user
// @Description Create a new user account
// @Tags Users
// @Accept json
// @Produce json
// @Param user body models.UserCreate true "User Sign Up"
// @Success 200 {object} web.WebResponse{data=models.UserResponse}
// @Failure 400 {object} web.WebResponse
// @Router /api/signup [post]
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
	http.Redirect(w, r, "/api/login", http.StatusOK)
}

// Login godoc
// @Summary Log in a user
// @Description Authenticate a user and set a session cookie
// @Tags Users
// @Accept json
// @Produce json
// @Param user body models.UserLogin true "User Login"
// @Success 200 {object} web.WebResponse{data=models.UserResponse}
// @Failure 401 {object} web.WebResponse
// @Router /api/login [post]
func (c *UserControllerImpl) Login(w http.ResponseWriter, r *http.Request) {
	userLogin := models.UserLogin{}
	helpers.ToRequestBody(r, &userLogin)

	userResponse, loggedIn := c.UserService.Login(r.Context(), userLogin)
	if loggedIn {
		helpers.SetUserCookie(w, r, userResponse.ID)
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
