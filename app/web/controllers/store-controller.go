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

// Create godoc
// @Summary Create store for the user
// @Description Create store for the authenticated user
// @Tags Store
// @Accept json
// @Produce json
// @Param store body models.StoreCreate true "Store create"
// @Success 200 {object} web.WebResponse{data=models.StoreResponse}
// @Failure 401 {object} web.WebResponse
// @Router /api [post]
// @Security BearerAuth
func (c *StoreControllerImpl) Create(w http.ResponseWriter, r *http.Request) {
	storeCreateRequest := models.StoreCreate{}
	helpers.ToRequestBody(r, &storeCreateRequest)
	cookie, err := helpers.GetUserCookie(w, r)
	if err != nil {
		http.Redirect(w, r, "/api/login", http.StatusUnauthorized)
	}

	userID := cookie.Value

	storeResponse := c.StoreService.Create(r.Context(), storeCreateRequest, userID)
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   storeResponse,
	}

	helpers.WriteResponseBody(w, webResponse)
}

// Update godoc
// @Summary Update store for the user
// @Description Update store for the authenticated user
// @Tags Store
// @Accept json
// @Produce json
// @Param store body models.StoreUpdate true "Store update"
// @Param storeId path string true "Store ID"
// @Success 200 {object} web.WebResponse{data=models.StoreResponse}
// @Failure 401 {object} web.WebResponse
// @Router /api/{storeId} [put]
// @Security BearerAuth
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

// Delete godoc
// @Summary Delete store for the user
// @Description Delete store for the authenticated user
// @Tags Store
// @Accept json
// @Produce json
// @Param storeId path string true "Store ID"
// @Success 200 {object} web.WebResponse
// @Failure 401 {object} web.WebResponse
// @Router /api/{storeId} [delete]
// @Security BearerAuth
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

// FindById godoc
// @Summary FindById store for the user
// @Description FindById store for the authenticated user
// @Tags Store
// @Accept json
// @Produce json
// @Param storeId path string true "Store ID"
// @Success 200 {object} web.WebResponse{data=models.StoreResponse}
// @Failure 401 {object} web.WebResponse
// @Router /api/{storeId} [get]
// @Security BearerAuth
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

// FindALl Store godoc
// @Summary Get all stores for the user
// @Description Get all stores for the authenticated user
// @Tags Store
// @Accept json
// @Produce json
// @Success 200 {object} web.WebResponse{data=models.StoreResponse}
// @Failure 401 {object} web.WebResponse
// @Router /api [get]
// @Security BearerAuth
func (c *StoreControllerImpl) FindAll(w http.ResponseWriter, r *http.Request) {
	cookie, err := helpers.GetUserCookie(w, r)
	if err != nil {
		http.Redirect(w, r, "/api/login", http.StatusUnauthorized)
	}

	userID := cookie.Value

	storeResponse := c.StoreService.FindAllByUserId(r.Context(), userID)
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   storeResponse,
	}
	helpers.WriteResponseBody(w, webResponse)
}
