package controllers

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/hiboedi/go-store-backend/app/helpers"
	"github.com/hiboedi/go-store-backend/app/web"
	"github.com/hiboedi/go-store-backend/app/web/models"
	"github.com/hiboedi/go-store-backend/app/web/services"
)

type SizeControllerImpl struct {
	SizeService services.SizeService
}

type SizeController interface {
	Create(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
	FindAll(w http.ResponseWriter, r *http.Request)
	FindById(w http.ResponseWriter, r *http.Request)
}

func NewSizeController(sizeService services.SizeService) SizeController {
	return &SizeControllerImpl{
		SizeService: sizeService,
	}
}

func (c *SizeControllerImpl) Create(w http.ResponseWriter, r *http.Request) {
	sizeCreateRequest := models.SizeCreate{}
	helpers.ToRequestBody(r, &sizeCreateRequest)

	vars := mux.Vars(r)
	storeId := vars["storeId"]
	sizeCreateRequest.StoreID = storeId

	sizeResponse := c.SizeService.Create(r.Context(), sizeCreateRequest)
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   sizeResponse,
	}

	helpers.WriteResponseBody(w, webResponse)
}

func (c *SizeControllerImpl) Update(w http.ResponseWriter, r *http.Request) {
	sizeUpdateRequest := models.SizeUpdate{}
	helpers.ToRequestBody(r, &sizeUpdateRequest)

	vars := mux.Vars(r)
	sizeId := vars["sizeId"]
	sizeUpdateRequest.ID = sizeId

	sizeResponse := c.SizeService.Update(r.Context(), sizeUpdateRequest)
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   sizeResponse,
	}

	helpers.WriteResponseBody(w, webResponse)
}

func (c *SizeControllerImpl) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	sizeId := vars["sizeId"]

	c.SizeService.Delete(r.Context(), sizeId)
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "Ok",
	}
	helpers.WriteResponseBody(w, webResponse)
}

func (c *SizeControllerImpl) FindById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	sizeId := vars["sizeId"]

	sizeResponse := c.SizeService.FindById(r.Context(), sizeId)
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   sizeResponse,
	}
	helpers.WriteResponseBody(w, webResponse)
}

func (c *SizeControllerImpl) FindAll(w http.ResponseWriter, r *http.Request) {
	sizeResponse := c.SizeService.FindAll(r.Context())
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   sizeResponse,
	}
	helpers.WriteResponseBody(w, webResponse)
}
