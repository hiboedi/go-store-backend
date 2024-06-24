package controllers

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/hiboedi/go-store-backend/app/helpers"
	"github.com/hiboedi/go-store-backend/app/web"
	"github.com/hiboedi/go-store-backend/app/web/models"
	"github.com/hiboedi/go-store-backend/app/web/services"
)

type ImageControllerImpl struct {
	ImageService services.ImageService
}

type ImageController interface {
	Create(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
	FindAll(w http.ResponseWriter, r *http.Request)
	FindById(w http.ResponseWriter, r *http.Request)
}

func NewImageController(imageService services.ImageService) ImageController {
	return &ImageControllerImpl{
		ImageService: imageService,
	}
}

func (c *ImageControllerImpl) Create(w http.ResponseWriter, r *http.Request) {
	imageCreateRequest := models.ImageCreate{}
	helpers.ToRequestBody(r, &imageCreateRequest)

	vars := mux.Vars(r)
	productId := vars["productId"]
	imageCreateRequest.ProductID = productId

	imageResponse := c.ImageService.Create(r.Context(), imageCreateRequest)
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   imageResponse,
	}

	helpers.WriteResponseBody(w, webResponse)
}

func (c *ImageControllerImpl) Update(w http.ResponseWriter, r *http.Request) {
	imageUpdateRequest := models.ImageUpdate{}
	helpers.ToRequestBody(r, &imageUpdateRequest)

	vars := mux.Vars(r)
	imageId := vars["imageId"]

	imageResponse := c.ImageService.Update(r.Context(), imageUpdateRequest, imageId)
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   imageResponse,
	}

	helpers.WriteResponseBody(w, webResponse)
}

func (c *ImageControllerImpl) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	imageId := vars["imageId"]

	c.ImageService.Delete(r.Context(), imageId)
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "Ok",
	}
	helpers.WriteResponseBody(w, webResponse)
}

func (c *ImageControllerImpl) FindById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	imageId := vars["imageId"]

	imageResponse := c.ImageService.FindById(r.Context(), imageId)
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   imageResponse,
	}
	helpers.WriteResponseBody(w, webResponse)
}

func (c *ImageControllerImpl) FindAll(w http.ResponseWriter, r *http.Request) {
	imageResponse := c.ImageService.FindAll(r.Context())
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   imageResponse,
	}
	helpers.WriteResponseBody(w, webResponse)
}
