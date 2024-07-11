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

// Create Size godoc
// @Summary create Size for the store
// @Description create Size for the store
// @Tags Size
// @Accept json
// @Produce json
// @Param Size body models.SizeCreate true "Size create"
// @Param storeId path string true "Store ID"
// @Success 200 {object} web.WebResponse{data=models.SizeResponse}
// @Failure 401 {object} web.WebResponse
// @Router /api/{storeId}/sizes [post]
// @Security BearerAuth
func (c *SizeControllerImpl) Create(w http.ResponseWriter, r *http.Request) {
	sizeCreateRequest := models.SizeCreate{}
	helpers.ToRequestBody(r, &sizeCreateRequest)

	vars := mux.Vars(r)
	storeId := vars["storeId"]

	sizeResponse := c.SizeService.Create(r.Context(), sizeCreateRequest, storeId)
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   sizeResponse,
	}

	helpers.WriteResponseBody(w, webResponse)
}

// Update Size godoc
// @Summary Update Size from the store
// @Description Update Size from the store
// @Tags Size
// @Accept json
// @Produce json
// @Param Size body models.SizeUpdate true "Size Update"
// @Param sizeId path string true "Size ID"
// @Param storeId path string true "Store ID"
// @Success 200 {object} web.WebResponse{data=models.SizeResponse}
// @Failure 401 {object} web.WebResponse
// @Router /api/{storeId}/sizes/{sizeId} [put]
// @Security BearerAuth
func (c *SizeControllerImpl) Update(w http.ResponseWriter, r *http.Request) {
	sizeUpdateRequest := models.SizeUpdate{}
	helpers.ToRequestBody(r, &sizeUpdateRequest)

	vars := mux.Vars(r)
	sizeId := vars["sizeId"]

	sizeResponse := c.SizeService.Update(r.Context(), sizeUpdateRequest, sizeId)
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   sizeResponse,
	}

	helpers.WriteResponseBody(w, webResponse)
}

// Delete Size godoc
// @Summary Delete Size from the store
// @Description Delete Size from the store
// @Tags Size
// @Accept json
// @Produce json
// @Param sizeId path string true "Size ID"
// @Param storeId path string true "Store ID"
// @Success 200 {object} web.WebResponse
// @Failure 401 {object} web.WebResponse
// @Router /api/{storeId}/sizes/{sizeId} [delete]
// @Security BearerAuth
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

// FindById Size godoc
// @Summary FindById Size from the store
// @Description FindById Size from the store
// @Tags Size
// @Accept json
// @Produce json
// @Param sizeId path string true "Size ID"
// @Param storeId path string true "Store ID"
// @Success 200 {object} web.WebResponse{data=models.SizeResponse}
// @Failure 401 {object} web.WebResponse
// @Router /api/{storeId}/sizes/{sizeId} [get]
// @Security BearerAuth
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

// FindAll Size godoc
// @Summary FindAll Size from the store
// @Description FindAll Size from the store
// @Tags Size
// @Accept json
// @Produce json
// @Param storeId path string true "Store ID"
// @Success 200 {object} web.WebResponse{data=models.SizeResponse}
// @Failure 401 {object} web.WebResponse
// @Router /api/{storeId}/sizes [get]
// @Security BearerAuth
func (c *SizeControllerImpl) FindAll(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	storeId := vars["storeId"]

	sizeResponse := c.SizeService.FindAll(r.Context(), storeId)
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   sizeResponse,
	}
	helpers.WriteResponseBody(w, webResponse)
}
