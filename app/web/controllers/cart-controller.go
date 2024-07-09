package controllers

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/hiboedi/go-store-backend/app/helpers"
	"github.com/hiboedi/go-store-backend/app/web"
	"github.com/hiboedi/go-store-backend/app/web/models"
	"github.com/hiboedi/go-store-backend/app/web/services"
)

type CartController struct {
	cartService services.CartService
}

func NewCartController(cartService services.CartService) *CartController {
	return &CartController{cartService}
}

func (c *CartController) AddToCart(w http.ResponseWriter, r *http.Request) {
	cartItemRequest := models.CartItemCreate{}
	helpers.ToRequestBody(r, &cartItemRequest)

	cookie, err := helpers.GetUserCookie(w, r)
	if err != nil {
		http.Redirect(w, r, "/api/login", http.StatusUnauthorized)
	}
	userID := cookie.Value

	addCartResponse := c.cartService.AddToCart(r.Context(), cartItemRequest, userID)

	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   addCartResponse,
	}

	helpers.WriteResponseBody(w, webResponse)
}

func (c *CartController) CheckoutCart(w http.ResponseWriter, r *http.Request) {
	orderCreateRequest := models.OrderCreate{}
	helpers.ToRequestBody(r, &orderCreateRequest)

	cookie, err := helpers.GetUserCookie(w, r)
	if err != nil {
		http.Redirect(w, r, "/api/login", http.StatusUnauthorized)
	}
	userID := cookie.Value

	vars := mux.Vars(r)
	storeId := vars["storeId"]
	orderCreateRequest.StoreID = storeId

	orderResponse := c.cartService.CheckoutCart(r.Context(), orderCreateRequest, userID)

	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   orderResponse,
	}

	helpers.WriteResponseBody(w, webResponse)
}
