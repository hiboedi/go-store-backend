package controllers

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/hiboedi/go-store-backend/app/helpers"
	"github.com/hiboedi/go-store-backend/app/web"
	"github.com/hiboedi/go-store-backend/app/web/models"
	"github.com/hiboedi/go-store-backend/app/web/services"
)

type StoreControllerImpl struct {
	StoreService services.StoreService
}

type StoreController interface {
	Create(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
	FindAll(w http.ResponseWriter, r *http.Request)
	FindById(w http.ResponseWriter, r *http.Request)
}

func NewStoreController(storeService services.StoreService) StoreController {
	return &StoreControllerImpl{
		StoreService: storeService,
	}
}

func (c *StoreControllerImpl) Create(w http.ResponseWriter, r *http.Request) {
	storeCreateRequest := models.StoreCreate{}
	helpers.ToRequestBody(r, &storeCreateRequest)
	cookie, err := helpers.GetUserCookie(w, r)
	if err != nil {
		http.Redirect(w, r, "/api/login", http.StatusUnauthorized)
	}

	userID := cookie.Value
	storeCreateRequest.UserID = userID

	storeResponse := c.StoreService.Create(r.Context(), storeCreateRequest)
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   storeResponse,
	}

	helpers.WriteResponseBody(w, webResponse)
}

func (c *StoreControllerImpl) Update(w http.ResponseWriter, r *http.Request) {
	storeUpdateRequest := models.StoreUpdate{}
	helpers.ToRequestBody(r, &storeUpdateRequest)

	vars := mux.Vars(r)
	storeID := vars["storeId"]

	orderResponse := c.StoreService.Update(r.Context(), storeUpdateRequest, storeID)
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   orderResponse,
	}

	helpers.WriteResponseBody(w, webResponse)
}

func (c *StoreControllerImpl) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	storeID := vars["storeId"]

	c.StoreService.Delete(r.Context(), storeID)
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "Ok",
	}
	helpers.WriteResponseBody(w, webResponse)
}

func (c *StoreControllerImpl) FindById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	storeId := vars["storeId"]

	storeResponse := c.StoreService.FindById(r.Context(), storeId)
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   storeResponse,
	}
	helpers.WriteResponseBody(w, webResponse)
}

func (c *StoreControllerImpl) FindAll(w http.ResponseWriter, r *http.Request) {
	storeResponse := c.StoreService.FindAll(r.Context())
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   storeResponse,
	}
	helpers.WriteResponseBody(w, webResponse)
}
