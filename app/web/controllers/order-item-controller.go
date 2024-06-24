package controllers

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/hiboedi/go-store-backend/app/helpers"
	"github.com/hiboedi/go-store-backend/app/web"
	"github.com/hiboedi/go-store-backend/app/web/models"
	"github.com/hiboedi/go-store-backend/app/web/services"
)

type OrderItemControllerImpl struct {
	OrderItemService services.OrderItemService
}

type OrderItemController interface {
	Create(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
	FindByID(w http.ResponseWriter, r *http.Request)
	FindAllByOrderID(w http.ResponseWriter, r *http.Request)
}

func NewOrderItemController(orderItemService services.OrderItemService) OrderItemController {
	return &OrderItemControllerImpl{
		OrderItemService: orderItemService,
	}
}

func (c *OrderItemControllerImpl) Create(w http.ResponseWriter, r *http.Request) {
	orderItemCreateRequest := models.OrderItemCreate{}
	helpers.ToRequestBody(r, &orderItemCreateRequest)

	orderItemResponse := c.OrderItemService.Create(r.Context(), orderItemCreateRequest)
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   orderItemResponse,
	}

	helpers.WriteResponseBody(w, webResponse)
}

func (c *OrderItemControllerImpl) Update(w http.ResponseWriter, r *http.Request) {
	orderItemUpdateRequest := models.OrderItemUpdate{}
	helpers.ToRequestBody(r, &orderItemUpdateRequest)
	vars := mux.Vars(r)
	orderItemId := vars["orderItemId"]

	orderItemResponse := c.OrderItemService.Update(r.Context(), orderItemUpdateRequest, orderItemId)
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   orderItemResponse,
	}

	helpers.WriteResponseBody(w, webResponse)
}

func (c *OrderItemControllerImpl) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	orderItemId := vars["orderItemId"]

	c.OrderItemService.Delete(r.Context(), orderItemId)
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "Ok",
	}
	helpers.WriteResponseBody(w, webResponse)
}

func (c *OrderItemControllerImpl) FindByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	orderItemId := vars["orderItemId"]

	orderItemResponse := c.OrderItemService.FindByID(r.Context(), orderItemId)
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   orderItemResponse,
	}
	helpers.WriteResponseBody(w, webResponse)
}

func (c *OrderItemControllerImpl) FindAllByOrderID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	orderId := vars["orderId"]

	orderItemResponses := c.OrderItemService.FindAllByOrderID(r.Context(), orderId)
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   orderItemResponses,
	}
	helpers.WriteResponseBody(w, webResponse)
}
