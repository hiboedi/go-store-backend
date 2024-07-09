package controllers

import (
	"net/http"

	"github.com/hiboedi/go-store-backend/app/helpers"
	"github.com/hiboedi/go-store-backend/app/web"
	"github.com/hiboedi/go-store-backend/app/web/services"
)

type OrderController interface {
	FindAllOrder(w http.ResponseWriter, r *http.Request)
}

type OrderControllerImpl struct {
	OrderService services.OrderService
}

func NewOrderController(orderService services.OrderService) OrderController {
	return &OrderControllerImpl{
		OrderService: orderService,
	}
}

func (c *OrderControllerImpl) FindAllOrder(w http.ResponseWriter, r *http.Request) {

	data, err := c.OrderService.FindAllOrder(r.Context())
	helpers.PanicIfError(err)

	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   data,
	}

	helpers.WriteResponseBody(w, webResponse)
}
