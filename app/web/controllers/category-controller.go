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

// Create Category godoc
// @Summary create Category for the store
// @Description create Category for the store
// @Tags Category
// @Accept json
// @Produce json
// @Param category body models.CategoryCreate true "Category create"
// @Param storeId path string true "Store ID"
// @Success 200 {object} web.WebResponse{data=models.CategoryResponse}
// @Failure 401 {object} web.WebResponse
// @Router /api/{storeId}/categories [post]
// @Security BearerAuth
func (c *CategoryControllerImpl) Create(w http.ResponseWriter, r *http.Request) {
	categoryCreateRequest := models.CategoryCreate{}
	helpers.ToRequestBody(r, &categoryCreateRequest)

	vars := mux.Vars(r)
	storeId := vars["storeId"]

	categoryResponse := c.CategoryService.Create(r.Context(), categoryCreateRequest, storeId)
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   categoryResponse,
	}

	helpers.WriteResponseBody(w, webResponse)
}

// Update Category godoc
// @Summary Update Category for the store
// @Description Update Category for the store
// @Tags Category
// @Accept json
// @Produce json
// @Param category body models.CategoryUpdate true "Category Update"
// @Param storeId path string true "Store ID"
// @Param categoryId path string true "Category ID"
// @Success 200 {object} web.WebResponse{data=models.CategoryResponse}
// @Failure 401 {object} web.WebResponse
// @Router /api/{storeId}/categories/{categoryId} [put]
// @Security BearerAuth
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

// Delete Category godoc
// @Summary Delete Category from the store
// @Description Delete Category from the store
// @Tags Category
// @Accept json
// @Produce json
// @Param storeId path string true "Store ID"
// @Param categoryId path string true "Category ID"
// @Success 200 {object} web.WebResponse
// @Failure 401 {object} web.WebResponse
// @Router /api/{storeId}/categories/{categoryId} [delete]
// @Security BearerAuth
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

// FindById Category godoc
// @Summary FindById Category from the store
// @Description FindById Category from the store
// @Tags Category
// @Accept json
// @Produce json
// @Param storeId path string true "Store ID"
// @Param categoryId path string true "Category ID"
// @Success 200 {object} web.WebResponse{data=models.CategoryResponse}
// @Failure 401 {object} web.WebResponse
// @Router /api/{storeId}/categories/{categoryId} [get]
// @Security BearerAuth
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

// FindAll Categories godoc
// @Summary FindAll Categories from the store
// @Description FindAll Categories from the store
// @Tags Category
// @Accept json
// @Produce json
// @Param storeId path string true "Store ID"
// @Success 200 {object} web.WebResponse{data=models.CategoryResponse}
// @Failure 401 {object} web.WebResponse
// @Router /api/{storeId}/categories [get]
// @Security BearerAuth
func (c *CategoryControllerImpl) FindAll(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	storeId := vars["storeId"]

	categoryResponse := c.CategoryService.FindAll(r.Context(), storeId)
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   categoryResponse,
	}
	helpers.WriteResponseBody(w, webResponse)
}
