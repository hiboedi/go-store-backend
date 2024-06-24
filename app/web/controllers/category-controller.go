package controllers

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/hiboedi/go-store-backend/app/helpers"
	"github.com/hiboedi/go-store-backend/app/web"
	"github.com/hiboedi/go-store-backend/app/web/models"
	"github.com/hiboedi/go-store-backend/app/web/services"
)

type CategoryControllerImpl struct {
	CategoryService services.CategoryService
}

type CategoryController interface {
	Create(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
	FindAll(w http.ResponseWriter, r *http.Request)
	FindById(w http.ResponseWriter, r *http.Request)
}

func NewCategoryController(categoryService services.CategoryService) CategoryController {
	return &CategoryControllerImpl{
		CategoryService: categoryService,
	}
}

func (c *CategoryControllerImpl) Create(w http.ResponseWriter, r *http.Request) {
	categoryCreateRequest := models.CategoryCreate{}
	helpers.ToRequestBody(r, &categoryCreateRequest)

	vars := mux.Vars(r)
	storeId := vars["storeId"]

	categoryCreateRequest.StoreID = storeId

	categoryResponse := c.CategoryService.Create(r.Context(), categoryCreateRequest)
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   categoryResponse,
	}

	helpers.WriteResponseBody(w, webResponse)
}

func (c *CategoryControllerImpl) Update(w http.ResponseWriter, r *http.Request) {
	categoryUpdateRequest := models.CategoryUpdate{}
	helpers.ToRequestBody(r, &categoryUpdateRequest)

	vars := mux.Vars(r)
	categoryId := vars["categoryId"]

	categoryResponse := c.CategoryService.Update(r.Context(), categoryUpdateRequest, categoryId)
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   categoryResponse,
	}

	helpers.WriteResponseBody(w, webResponse)
}

func (c *CategoryControllerImpl) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	categoryId := vars["categoryId"]

	c.CategoryService.Delete(r.Context(), categoryId)
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "Ok",
	}
	helpers.WriteResponseBody(w, webResponse)
}

func (c *CategoryControllerImpl) FindById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	categoryId := vars["categoryId"]

	categoryResponse := c.CategoryService.FindById(r.Context(), categoryId)
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   categoryResponse,
	}
	helpers.WriteResponseBody(w, webResponse)
}

func (c *CategoryControllerImpl) FindAll(w http.ResponseWriter, r *http.Request) {
	categoryResponse := c.CategoryService.FindAll(r.Context())
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   categoryResponse,
	}
	helpers.WriteResponseBody(w, webResponse)
}
