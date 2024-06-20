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

func (c *ColorControllerImpl) Create(w http.ResponseWriter, r *http.Request) {
	colorCreateRequest := models.ColorCreate{}
	helpers.ToRequestBody(r, &colorCreateRequest)

	vars := mux.Vars(r)
	storeId := vars["storeId"]
	colorCreateRequest.ID = storeId

	colorResponse := c.ColorService.Create(r.Context(), colorCreateRequest)
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   colorResponse,
	}

	helpers.WriteResponseBody(w, webResponse)
}

func (c *ColorControllerImpl) Update(w http.ResponseWriter, r *http.Request) {
	colorUpdateRequest := models.ColorUpdate{}
	helpers.ToRequestBody(r, &colorUpdateRequest)

	vars := mux.Vars(r)
	colorId := vars["colorId"]
	colorUpdateRequest.ID = colorId

	colorResponse := c.ColorService.Update(r.Context(), colorUpdateRequest)
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   colorResponse,
	}

	helpers.WriteResponseBody(w, webResponse)
}

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

func (c *ColorControllerImpl) FindAll(w http.ResponseWriter, r *http.Request) {
	colorResponse := c.ColorService.FindAll(r.Context())
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   colorResponse,
	}
	helpers.WriteResponseBody(w, webResponse)
}
