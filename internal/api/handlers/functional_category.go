package handlers

import (
	"errors"
	"net/http"

	"apm/internal/models"
	"apm/internal/services"

	"github.com/gin-gonic/gin"
)

// FunctionalCategoryHandler handles HTTP requests for functional categories
type FunctionalCategoryHandler struct {
	service services.FunctionalCategoryService
}

// NewFunctionalCategoryHandler creates a new functional category handler
func NewFunctionalCategoryHandler(service services.FunctionalCategoryService) *FunctionalCategoryHandler {
	return &FunctionalCategoryHandler{
		service: service,
	}
}

// Register registers the routes for functional categories
func (h *FunctionalCategoryHandler) Register(router *gin.RouterGroup) {
	categories := router.Group("/functional-categories")
	{
		categories.POST("", h.Create)
		categories.GET("", h.List)
		categories.GET("/:id", h.GetByID)
		categories.PUT("/:id", h.Update)
		categories.DELETE("/:id", h.Delete)
	}
}

// Create handles the creation of a new functional category
func (h *FunctionalCategoryHandler) Create(c *gin.Context) {
	var req models.CreateFunctionalCategoryRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		RespondWithError(c, http.StatusBadRequest, err, "Invalid request body")
		return
	}

	resp, err := h.service.Create(c.Request.Context(), req)
	if err != nil {
		RespondWithError(c, http.StatusInternalServerError, err, "Failed to create functional category")
		return
	}

	c.JSON(http.StatusCreated, resp)
}

// GetByID handles the retrieval of a functional category by ID
func (h *FunctionalCategoryHandler) GetByID(c *gin.Context) {
	id := ExtractIDParam(c)
	if id == "" {
		RespondWithError(c, http.StatusBadRequest, errors.New("missing or invalid ID"), "Missing or invalid ID")
		return
	}

	resp, err := h.service.GetByID(c.Request.Context(), id)
	if err != nil {
		RespondWithError(c, http.StatusNotFound, err, "Functional category not found")
		return
	}

	c.JSON(http.StatusOK, resp)
}

// List handles the retrieval of a list of functional categories
func (h *FunctionalCategoryHandler) List(c *gin.Context) {
	limit, offset := SetPagination(c)

	resp, err := h.service.List(c.Request.Context(), limit, offset)
	if err != nil {
		RespondWithError(c, http.StatusInternalServerError, err, "Failed to retrieve functional categories")
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":   resp,
		"limit":  limit,
		"offset": offset,
		"count":  len(resp),
	})
}

// Update handles the update of a functional category
func (h *FunctionalCategoryHandler) Update(c *gin.Context) {
	id := ExtractIDParam(c)
	if id == "" {
		RespondWithError(c, http.StatusBadRequest, errors.New("missing or invalid ID"), "Missing or invalid ID")
		return
	}

	var req models.UpdateFunctionalCategoryRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		RespondWithError(c, http.StatusBadRequest, err, "Invalid request body")
		return
	}

	if err := h.service.Update(c.Request.Context(), id, req); err != nil {
		RespondWithError(c, http.StatusInternalServerError, err, "Failed to update functional category")
		return
	}

	c.Status(http.StatusNoContent)
}

// Delete handles the deletion of a functional category
func (h *FunctionalCategoryHandler) Delete(c *gin.Context) {
	id := ExtractIDParam(c)
	if id == "" {
		RespondWithError(c, http.StatusBadRequest, errors.New("missing or invalid ID"), "Missing or invalid ID")
		return
	}

	if err := h.service.Delete(c.Request.Context(), id); err != nil {
		RespondWithError(c, http.StatusInternalServerError, err, "Failed to delete functional category")
		return
	}

	c.Status(http.StatusNoContent)
}
