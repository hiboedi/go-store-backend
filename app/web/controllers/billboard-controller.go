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

func (c *BillboardControllerImpl) Create(w http.ResponseWriter, r *http.Request) {
	billboardCreateRequest := models.BillboardCreate{}
	helpers.ToRequestBody(r, &billboardCreateRequest)

	vars := mux.Vars(r)
	storeId := vars["storeId"]
	billboardCreateRequest.StoreID = storeId

	billboardResponse := c.BillboardService.Create(r.Context(), billboardCreateRequest)
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   billboardResponse,
	}

	helpers.WriteResponseBody(w, webResponse)
}

func (c *BillboardControllerImpl) Update(w http.ResponseWriter, r *http.Request) {
	billboardUpdateRequest := models.BillboardUpdate{}
	helpers.ToRequestBody(r, &billboardUpdateRequest)

	vars := mux.Vars(r)
	billboardId := vars["billboardId"]
	billboardUpdateRequest.ID = billboardId

	billboardResponse := c.BillboardService.Update(r.Context(), billboardUpdateRequest)
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   billboardResponse,
	}

	helpers.WriteResponseBody(w, webResponse)
}

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

func (c *BillboardControllerImpl) FindAll(w http.ResponseWriter, r *http.Request) {
	billboardResponse := c.BillboardService.FindAll(r.Context())
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   billboardResponse,
	}
	helpers.WriteResponseBody(w, webResponse)
}
