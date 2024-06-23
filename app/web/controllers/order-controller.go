package controllers

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/hiboedi/go-store-backend/app/helpers"
	"github.com/hiboedi/go-store-backend/app/web"
	"github.com/hiboedi/go-store-backend/app/web/models"
	"github.com/hiboedi/go-store-backend/app/web/services"
)

type OrderControllerImpl struct {
	OrderService services.OrderService
}

type OrderController interface {
	Create(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
	FindAll(w http.ResponseWriter, r *http.Request)
	FindById(w http.ResponseWriter, r *http.Request)
}

func NewOrderController(orderService services.OrderService) OrderController {
	return &OrderControllerImpl{
		OrderService: orderService,
	}
}

func (c *OrderControllerImpl) Create(w http.ResponseWriter, r *http.Request) {
	orderCreateRequest := models.OrderCreate{}
	helpers.ToRequestBody(r, &orderCreateRequest)

	vars := mux.Vars(r)
	storeId := vars["storeId"]
	orderCreateRequest.StoreID = storeId

	orderResponse := c.OrderService.Create(r.Context(), orderCreateRequest)
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   orderResponse,
	}

	helpers.WriteResponseBody(w, webResponse)
}

func (c *OrderControllerImpl) Update(w http.ResponseWriter, r *http.Request) {
	orderUpdateRequest := models.OrderUpdate{}
	helpers.ToRequestBody(r, &orderUpdateRequest)

	vars := mux.Vars(r)
	orderId := vars["orderId"]
	orderUpdateRequest.ID = orderId

	orderResponse := c.OrderService.Update(r.Context(), orderUpdateRequest)
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   orderResponse,
	}

	helpers.WriteResponseBody(w, webResponse)
}

func (c *OrderControllerImpl) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	orderId := vars["orderId"]

	c.OrderService.Delete(r.Context(), orderId)
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "Ok",
	}
	helpers.WriteResponseBody(w, webResponse)
}

func (c *OrderControllerImpl) FindById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	orderId := vars["orderId"]

	orderResponse := c.OrderService.FindById(r.Context(), orderId)
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   orderResponse,
	}
	helpers.WriteResponseBody(w, webResponse)
}

func (c *OrderControllerImpl) FindAll(w http.ResponseWriter, r *http.Request) {
	orderResponse := c.OrderService.FindAll(r.Context())
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   orderResponse,
	}
	helpers.WriteResponseBody(w, webResponse)
}
