package controllers

import (
	"net/http"
	"strconv"

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

// Create Product godoc
// @Summary create Product for the store
// @Description create Product for the store
// @Tags Product
// @Accept json
// @Produce json
// @Param Product body models.ProductCreate true "Product create"
// @Param storeId path string true "Store ID"
// @Success 200 {object} web.WebResponse{data=models.ProductResponse}
// @Failure 401 {object} web.WebResponse
// @Router /api/{storeId}/products [post]
// @Security BearerAuth
func (c *ProductControllerImpl) Create(w http.ResponseWriter, r *http.Request) {
	productCreateRequest := models.ProductCreate{}
	helpers.ToRequestBody(r, &productCreateRequest)

	vars := mux.Vars(r)
	storeId := vars["storeId"]

	productResponse := c.ProductService.Create(r.Context(), productCreateRequest, storeId)
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   productResponse,
	}

	helpers.WriteResponseBody(w, webResponse)
}

// Update Product godoc
// @Summary Update Product from the store
// @Description Update Product from the store
// @Tags Product
// @Accept json
// @Produce json
// @Param Product body models.ProductUpdate true "Product Update"
// @Param storeId path string true "Store ID"
// @Param productId path string true "Product ID"
// @Success 200 {object} web.WebResponse{data=models.ProductResponse}
// @Failure 401 {object} web.WebResponse
// @Router /api/{storeId}/products/{productId} [put]
// @Security BearerAuth
func (c *ProductControllerImpl) Update(w http.ResponseWriter, r *http.Request) {
	productUpdateRequest := models.ProductUpdate{}
	helpers.ToRequestBody(r, &productUpdateRequest)

	vars := mux.Vars(r)
	productId := vars["productId"]

	productResponse := c.ProductService.Update(r.Context(), productUpdateRequest, productId)
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   productResponse,
	}

	helpers.WriteResponseBody(w, webResponse)
}

// Delete Product godoc
// @Summary Delete Product from the store
// @Description Delete Product from the store
// @Tags Product
// @Accept json
// @Produce json
// @Param storeId path string true "Store ID"
// @Param productId path string true "Product ID"
// @Success 200 {object} web.WebResponse
// @Failure 401 {object} web.WebResponse
// @Router /api/{storeId}/products/{productId} [delete]
// @Security BearerAuth
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

// FindById Product godoc
// @Summary FindById Product from the store
// @Description FindById Product from the store
// @Tags Product
// @Accept json
// @Produce json
// @Param storeId path string true "Store ID"
// @Param productId path string true "Product ID"
// @Success 200 {object} web.WebResponse{data=models.ProductResponse}
// @Failure 401 {object} web.WebResponse
// @Router /api/{storeId}/products/{productId} [get]
// @Security BearerAuth
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

// FindAll Products godoc
// @Summary FindAll Products from the store
// @Description FindAll Products from the store
// @Tags Product
// @Accept json
// @Produce json
// @Param storeId path string true "Store ID"
// @Success 200 {object} web.WebResponse{data=models.ProductResponse}
// @Failure 401 {object} web.WebResponse
// @Router /api/{storeId}/products [get]
// @Security BearerAuth
func (c *ProductControllerImpl) FindAll(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	storeId := vars["storeId"]

	pageStr := r.URL.Query().Get("page")
	page := int64(1)
	if pageStr != "" {
		pageParams, err := strconv.Atoi(pageStr)
		if err != nil {
			http.Error(w, "Invalid page parameter", http.StatusBadRequest)
			return
		}
		page = int64(pageParams)
	}

	productResponse := c.ProductService.FindAll(r.Context(), storeId, page)
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   productResponse,
	}
	helpers.WriteResponseBody(w, webResponse)
}
