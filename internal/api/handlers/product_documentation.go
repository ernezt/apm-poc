package handlers

import (
	"errors"
	"net/http"

	"apm/internal/models"
	"apm/internal/services"

	"github.com/gin-gonic/gin"
)

// ProductDocumentationHandler handles HTTP requests for product documentation
type ProductDocumentationHandler struct {
	service services.ProductDocumentationService
}

// NewProductDocumentationHandler creates a new product documentation handler
func NewProductDocumentationHandler(service services.ProductDocumentationService) *ProductDocumentationHandler {
	return &ProductDocumentationHandler{
		service: service,
	}
}

// Register registers the routes for product documentation
func (h *ProductDocumentationHandler) Register(router *gin.RouterGroup) {
	docs := router.Group("/product-documentation")
	{
		docs.POST("", h.Create)
		docs.GET("", h.List)
		docs.GET("/:id", h.GetByID)
		docs.PUT("/:id", h.Update)
		docs.DELETE("/:id", h.Delete)
	}
}

// Create handles the creation of new product documentation
func (h *ProductDocumentationHandler) Create(c *gin.Context) {
	var req models.CreateProductDocumentationRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		RespondWithError(c, http.StatusBadRequest, err, "Invalid request body")
		return
	}

	resp, err := h.service.Create(c.Request.Context(), req)
	if err != nil {
		RespondWithError(c, http.StatusInternalServerError, err, "Failed to create product documentation")
		return
	}

	c.JSON(http.StatusCreated, resp)
}

// GetByID handles the retrieval of product documentation by ID
func (h *ProductDocumentationHandler) GetByID(c *gin.Context) {
	id := ExtractIDParam(c)
	if id == "" {
		RespondWithError(c, http.StatusBadRequest, errors.New("missing or invalid ID"), "Missing or invalid ID")
		return
	}

	resp, err := h.service.GetByID(c.Request.Context(), id)
	if err != nil {
		RespondWithError(c, http.StatusNotFound, err, "Product documentation not found")
		return
	}

	c.JSON(http.StatusOK, resp)
}

// List handles the retrieval of a list of product documentation
func (h *ProductDocumentationHandler) List(c *gin.Context) {
	limit, offset := SetPagination(c)

	resp, err := h.service.List(c.Request.Context(), limit, offset)
	if err != nil {
		RespondWithError(c, http.StatusInternalServerError, err, "Failed to retrieve product documentation list")
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":   resp,
		"limit":  limit,
		"offset": offset,
		"count":  len(resp),
	})
}

// Update handles the update of product documentation
func (h *ProductDocumentationHandler) Update(c *gin.Context) {
	id := ExtractIDParam(c)
	if id == "" {
		RespondWithError(c, http.StatusBadRequest, errors.New("missing or invalid ID"), "Missing or invalid ID")
		return
	}

	var req models.UpdateProductDocumentationRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		RespondWithError(c, http.StatusBadRequest, err, "Invalid request body")
		return
	}

	if err := h.service.Update(c.Request.Context(), id, req); err != nil {
		RespondWithError(c, http.StatusInternalServerError, err, "Failed to update product documentation")
		return
	}

	c.Status(http.StatusNoContent)
}

// Delete handles the deletion of product documentation
func (h *ProductDocumentationHandler) Delete(c *gin.Context) {
	id := ExtractIDParam(c)
	if id == "" {
		RespondWithError(c, http.StatusBadRequest, errors.New("missing or invalid ID"), "Missing or invalid ID")
		return
	}

	if err := h.service.Delete(c.Request.Context(), id); err != nil {
		RespondWithError(c, http.StatusInternalServerError, err, "Failed to delete product documentation")
		return
	}

	c.Status(http.StatusNoContent)
}
