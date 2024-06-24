package controllers

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/hiboedi/go-store-backend/app/helpers"
	"github.com/hiboedi/go-store-backend/app/web"
	"github.com/hiboedi/go-store-backend/app/web/models"
	"github.com/hiboedi/go-store-backend/app/web/services"
)

type ProductControllerImpl struct {
	ProductService services.ProductService
}

type ProductController interface {
	Create(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
	FindAll(w http.ResponseWriter, r *http.Request)
	FindById(w http.ResponseWriter, r *http.Request)
}

func NewProductController(productService services.ProductService) ProductController {
	return &ProductControllerImpl{
		ProductService: productService,
	}
}

func (c *ProductControllerImpl) Create(w http.ResponseWriter, r *http.Request) {
	productCreateRequest := models.ProductCreate{}
	helpers.ToRequestBody(r, &productCreateRequest)

	vars := mux.Vars(r)
	storeId := vars["storeId"]
	productCreateRequest.StoreID = storeId

	productResponse := c.ProductService.Create(r.Context(), productCreateRequest)
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   productResponse,
	}

	helpers.WriteResponseBody(w, webResponse)
}

func (c *ProductControllerImpl) Update(w http.ResponseWriter, r *http.Request) {
	productUpdateRequest := models.ProductUpdate{}
	helpers.ToRequestBody(r, &productUpdateRequest)

	vars := mux.Vars(r)
	productId := vars["productId"]

	storeId := vars["storeId"]
	productUpdateRequest.StoreID = storeId

	productResponse := c.ProductService.Update(r.Context(), productUpdateRequest, productId)
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   productResponse,
	}

	helpers.WriteResponseBody(w, webResponse)
}

func (c *ProductControllerImpl) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	productId := vars["productId"]

	c.ProductService.Delete(r.Context(), productId)
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "Ok",
	}
	helpers.WriteResponseBody(w, webResponse)
}

func (c *ProductControllerImpl) FindById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	productId := vars["productId"]

	productResponse := c.ProductService.FindById(r.Context(), productId)
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   productResponse,
	}
	helpers.WriteResponseBody(w, webResponse)
}

func (c *ProductControllerImpl) FindAll(w http.ResponseWriter, r *http.Request) {
	productResponse := c.ProductService.FindAll(r.Context())
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   productResponse,
	}
	helpers.WriteResponseBody(w, webResponse)
}
