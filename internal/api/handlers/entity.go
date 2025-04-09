package handlers

import (
	"errors"
	"net/http"

	"apm/internal/models"
	"apm/internal/services"

	"github.com/gin-gonic/gin"
)

// EntityHandler handles HTTP requests for entities
type EntityHandler struct {
	service services.EntityService
}

// NewEntityHandler creates a new entity handler
func NewEntityHandler(service services.EntityService) *EntityHandler {
	return &EntityHandler{
		service: service,
	}
}

// Register registers the routes for entities
func (h *EntityHandler) Register(router *gin.RouterGroup) {
	entities := router.Group("/entities")
	{
		entities.POST("", h.Create)
		entities.GET("", h.List)
		entities.GET("/:id", h.GetByID)
		entities.PUT("/:id", h.Update)
		entities.DELETE("/:id", h.Delete)
	}
}

// Create handles the creation of a new entity
func (h *EntityHandler) Create(c *gin.Context) {
	var req models.CreateEntityRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		RespondWithError(c, http.StatusBadRequest, err, "Invalid request body")
		return
	}

	resp, err := h.service.Create(c.Request.Context(), req)
	if err != nil {
		RespondWithError(c, http.StatusInternalServerError, err, "Failed to create entity")
		return
	}

	c.JSON(http.StatusCreated, resp)
}

// GetByID handles the retrieval of an entity by ID
func (h *EntityHandler) GetByID(c *gin.Context) {
	id := ExtractIDParam(c)
	if id == "" {
		RespondWithError(c, http.StatusBadRequest, errors.New("missing or invalid ID"), "Missing or invalid ID")
		return
	}

	resp, err := h.service.GetByID(c.Request.Context(), id)
	if err != nil {
		RespondWithError(c, http.StatusNotFound, err, "Entity not found")
		return
	}

	c.JSON(http.StatusOK, resp)
}

// List handles the retrieval of a list of entities
func (h *EntityHandler) List(c *gin.Context) {
	limit, offset := SetPagination(c)

	resp, err := h.service.List(c.Request.Context(), limit, offset)
	if err != nil {
		RespondWithError(c, http.StatusInternalServerError, err, "Failed to retrieve entities")
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":   resp,
		"limit":  limit,
		"offset": offset,
		"count":  len(resp),
	})
}

// Update handles the update of an entity
func (h *EntityHandler) Update(c *gin.Context) {
	id := ExtractIDParam(c)
	if id == "" {
		RespondWithError(c, http.StatusBadRequest, errors.New("missing or invalid ID"), "Missing or invalid ID")
		return
	}

	var req models.UpdateEntityRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		RespondWithError(c, http.StatusBadRequest, err, "Invalid request body")
		return
	}

	if err := h.service.Update(c.Request.Context(), id, req); err != nil {
		RespondWithError(c, http.StatusInternalServerError, err, "Failed to update entity")
		return
	}

	c.Status(http.StatusNoContent)
}

// Delete handles the deletion of an entity
func (h *EntityHandler) Delete(c *gin.Context) {
	id := ExtractIDParam(c)
	if id == "" {
		RespondWithError(c, http.StatusBadRequest, errors.New("missing or invalid ID"), "Missing or invalid ID")
		return
	}

	if err := h.service.Delete(c.Request.Context(), id); err != nil {
		RespondWithError(c, http.StatusInternalServerError, err, "Failed to delete entity")
		return
	}

	c.Status(http.StatusNoContent)
}
