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

// FindAll Order godoc
// @Summary FindAll Order from the store
// @Description FindAll Order from the store
// @Tags Order
// @Accept json
// @Produce json
// @Param storeId path string true "Store ID"
// @Success 200 {object} web.WebResponse{data=models.OrderResponse}
// @Failure 401 {object} web.WebResponse
// @Router /api/{storeId}/orders [get]
// @Security BearerAuth
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
