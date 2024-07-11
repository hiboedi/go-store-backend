package controllers

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/hiboedi/go-store-backend/app/helpers"
	"github.com/hiboedi/go-store-backend/app/web"
	"github.com/hiboedi/go-store-backend/app/web/models"
	"github.com/hiboedi/go-store-backend/app/web/services"
)

type ColorControllerImpl struct {
	ColorService services.ColorService
}

type ColorController interface {
	Create(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
	FindAll(w http.ResponseWriter, r *http.Request)
	FindById(w http.ResponseWriter, r *http.Request)
}

func NewColorController(colorService services.ColorService) ColorController {
	return &ColorControllerImpl{
		ColorService: colorService,
	}
}

// Create Color godoc
// @Summary create Color for the store
// @Description create Color for the store
// @Tags Color
// @Accept json
// @Produce json
// @Param Color body models.ColorCreate true "Color create"
// @Param storeId path string true "Store ID"
// @Success 200 {object} web.WebResponse{data=models.ColorResponse}
// @Failure 401 {object} web.WebResponse
// @Router /api/{storeId}/colors [post]
// @Security BearerAuth
func (c *ColorControllerImpl) Create(w http.ResponseWriter, r *http.Request) {
	colorCreateRequest := models.ColorCreate{}
	helpers.ToRequestBody(r, &colorCreateRequest)

	vars := mux.Vars(r)
	storeId := vars["storeId"]

	colorResponse := c.ColorService.Create(r.Context(), colorCreateRequest, storeId)
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   colorResponse,
	}

	helpers.WriteResponseBody(w, webResponse)
}

// Update Color godoc
// @Summary Update Color from the store
// @Description Update Color from the store
// @Tags Color
// @Accept json
// @Produce json
// @Param Color body models.ColorUpdate true "Color Update"
// @Param storeId path string true "Store ID"
// @Param colorId path string true "Color ID"
// @Success 200 {object} web.WebResponse{data=models.ColorResponse}
// @Failure 401 {object} web.WebResponse
// @Router /api/{storeId}/colors/{colorId} [put]
// @Security BearerAuth
func (c *ColorControllerImpl) Update(w http.ResponseWriter, r *http.Request) {
	colorUpdateRequest := models.ColorUpdate{}
	helpers.ToRequestBody(r, &colorUpdateRequest)

	vars := mux.Vars(r)
	colorId := vars["colorId"]

	colorResponse := c.ColorService.Update(r.Context(), colorUpdateRequest, colorId)
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   colorResponse,
	}

	helpers.WriteResponseBody(w, webResponse)
}

// Delete Color godoc
// @Summary Delete Color from the store
// @Description Delete Color from the store
// @Tags Color
// @Accept json
// @Produce json
// @Param storeId path string true "Store ID"
// @Param colorId path string true "Color ID"
// @Success 200 {object} web.WebResponse
// @Failure 401 {object} web.WebResponse
// @Router /api/{storeId}/colors/{colorId} [delete]
// @Security BearerAuth
func (c *ColorControllerImpl) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	colorId := vars["colorId"]

	c.ColorService.Delete(r.Context(), colorId)
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "Ok",
	}
	helpers.WriteResponseBody(w, webResponse)
}

// FindById Color godoc
// @Summary FindById Color from the store
// @Description FindById Color from the store
// @Tags Color
// @Accept json
// @Produce json
// @Param storeId path string true "Store ID"
// @Param colorId path string true "Color ID"
// @Success 200 {object} web.WebResponse{data=models.ColorResponse}
// @Failure 401 {object} web.WebResponse
// @Router /api/{storeId}/colors/{colorId} [get]
// @Security BearerAuth
func (c *ColorControllerImpl) FindById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	colorId := vars["colorId"]

	colorResponse := c.ColorService.FindById(r.Context(), colorId)
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   colorResponse,
	}
	helpers.WriteResponseBody(w, webResponse)
}

// FindAll Color godoc
// @Summary FindAll Color from the store
// @Description FindAll Color from the store
// @Tags Color
// @Accept json
// @Produce json
// @Param storeId path string true "Store ID"
// @Success 200 {object} web.WebResponse{data=models.ColorResponse}
// @Failure 401 {object} web.WebResponse
// @Router /api/{storeId}/colors [get]
// @Security BearerAuth
func (c *ColorControllerImpl) FindAll(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	storeId := vars["storeId"]

	colorResponse := c.ColorService.FindAll(r.Context(), storeId)
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   colorResponse,
	}
	helpers.WriteResponseBody(w, webResponse)
}
