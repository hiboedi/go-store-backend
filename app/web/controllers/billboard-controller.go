package controllers

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/hiboedi/go-store-backend/app/helpers"
	"github.com/hiboedi/go-store-backend/app/web"
	"github.com/hiboedi/go-store-backend/app/web/models"
	"github.com/hiboedi/go-store-backend/app/web/services"
)

type BillboardControllerImpl struct {
	BillboardService services.BillboardService
}

type BillboardController interface {
	Create(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
	FindAll(w http.ResponseWriter, r *http.Request)
	FindById(w http.ResponseWriter, r *http.Request)
}

func NewBillboardController(billboardService services.BillboardService) BillboardController {
	return &BillboardControllerImpl{
		BillboardService: billboardService,
	}
}

// Create billboard godoc
// @Summary create billboard for the store
// @Description create billboard for the store
// @Tags Billboard
// @Accept json
// @Produce json
// @Param billboard body models.BillboardCreate true "billboard create"
// @Param storeId path string true "Store ID"
// @Success 200 {object} web.WebResponse{data=models.BillboardResponse}
// @Failure 401 {object} web.WebResponse
// @Router /api/{storeId}/billboards [post]
// @Security BearerAuth
func (c *BillboardControllerImpl) Create(w http.ResponseWriter, r *http.Request) {
	billboardCreateRequest := models.BillboardCreate{}
	helpers.ToRequestBody(r, &billboardCreateRequest)

	vars := mux.Vars(r)
	storeId := vars["storeId"]

	billboardResponse := c.BillboardService.Create(r.Context(), billboardCreateRequest, storeId)
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   billboardResponse,
	}

	helpers.WriteResponseBody(w, webResponse)
}

// Update billboard godoc
// @Summary Update billboard for the store
// @Description Update billboard for the store
// @Tags Billboard
// @Accept json
// @Produce json
// @Param billboard body models.BillboardUpdate true "billboad Update"
// @Param storeId path string true "Store ID"
// @Param billboardId path string true "BIllboard ID"
// @Success 200 {object} web.WebResponse{data=models.BillboardResponse}
// @Failure 401 {object} web.WebResponse
// @Router /api/{storeId}/billboards/{billboardId} [put]
// @Security BearerAuth
func (c *BillboardControllerImpl) Update(w http.ResponseWriter, r *http.Request) {
	billboardUpdateRequest := models.BillboardUpdate{}
	helpers.ToRequestBody(r, &billboardUpdateRequest)

	vars := mux.Vars(r)
	billboardId := vars["billboardId"]

	billboardResponse := c.BillboardService.Update(r.Context(), billboardUpdateRequest, billboardId)
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   billboardResponse,
	}

	helpers.WriteResponseBody(w, webResponse)
}

// Delete billboard godoc
// @Summary Delete billboard from the store
// @Description Delete billboard from the store
// @Tags Billboard
// @Accept json
// @Produce json
// @Param storeId path string true "Store ID"
// @Param billboardId path string true "BIllboard ID"
// @Success 200 {object} web.WebResponse
// @Failure 401 {object} web.WebResponse
// @Router /api/{storeId}/billboards/{billboardId} [delete]
// @Security BearerAuth
func (c *BillboardControllerImpl) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	billboardId := vars["billboardId"]

	c.BillboardService.Delete(r.Context(), billboardId)
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "Ok",
	}
	helpers.WriteResponseBody(w, webResponse)
}

// FindById billboard godoc
// @Summary FindById billboard from the store
// @Description FindById billboard from the store
// @Tags Billboard
// @Accept json
// @Produce json
// @Param storeId path string true "Store ID"
// @Param billboardId path string true "BIllboard ID"
// @Success 200 {object} web.WebResponse{data=models.BillboardResponse}
// @Failure 401 {object} web.WebResponse
// @Router /api/{storeId}/billboards/{billboardId} [get]
// @Security BearerAuth
func (c *BillboardControllerImpl) FindById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	billboardId := vars["billboardId"]

	billboardResponse := c.BillboardService.FindById(r.Context(), billboardId)
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   billboardResponse,
	}
	helpers.WriteResponseBody(w, webResponse)
}

// FindALl billboard godoc
// @Summary Get all billboards for the store
// @Description Get all billboards for the store
// @Tags Billboard
// @Accept json
// @Produce json
// @Param storeId path string true "Store ID"
// @Success 200 {object} web.WebResponse{data=models.BillboardResponse}
// @Failure 401 {object} web.WebResponse
// @Router /api/{storeId}/billboards [get]
// @Security BearerAuth
func (c *BillboardControllerImpl) FindAll(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	storeId := vars["storeId"]

	billboardResponse := c.BillboardService.FindAll(r.Context(), storeId)
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   billboardResponse,
	}
	helpers.WriteResponseBody(w, webResponse)
}
