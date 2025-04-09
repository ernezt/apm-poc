package handlers

import (
	"errors"
	"net/http"

	"apm/internal/models"
	"apm/internal/services"

	"github.com/gin-gonic/gin"
)

// StatusHandler handles HTTP requests for statuses
type StatusHandler struct {
	service services.StatusService
}

// NewStatusHandler creates a new status handler
func NewStatusHandler(service services.StatusService) *StatusHandler {
	return &StatusHandler{
		service: service,
	}
}

// Register registers the routes for statuses
func (h *StatusHandler) Register(router *gin.RouterGroup) {
	statuses := router.Group("/statuses")
	{
		statuses.POST("", h.Create)
		statuses.GET("", h.List)
		statuses.GET("/:id", h.GetByID)
		statuses.PUT("/:id", h.Update)
		statuses.DELETE("/:id", h.Delete)
	}
}

// Create handles the creation of a new status
func (h *StatusHandler) Create(c *gin.Context) {
	var req models.CreateStatusRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		RespondWithError(c, http.StatusBadRequest, err, "Invalid request body")
		return
	}

	resp, err := h.service.Create(c.Request.Context(), req)
	if err != nil {
		RespondWithError(c, http.StatusInternalServerError, err, "Failed to create status")
		return
	}

	c.JSON(http.StatusCreated, resp)
}

// GetByID handles the retrieval of a status by ID
func (h *StatusHandler) GetByID(c *gin.Context) {
	id := ExtractIDParam(c)
	if id == "" {
		RespondWithError(c, http.StatusBadRequest, errors.New("missing or invalid ID"), "Missing or invalid ID")
		return
	}

	resp, err := h.service.GetByID(c.Request.Context(), id)
	if err != nil {
		RespondWithError(c, http.StatusNotFound, err, "Status not found")
		return
	}

	c.JSON(http.StatusOK, resp)
}

// List handles the retrieval of a list of statuses
func (h *StatusHandler) List(c *gin.Context) {
	limit, offset := SetPagination(c)

	resp, err := h.service.List(c.Request.Context(), limit, offset)
	if err != nil {
		RespondWithError(c, http.StatusInternalServerError, err, "Failed to retrieve statuses")
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":   resp,
		"limit":  limit,
		"offset": offset,
		"count":  len(resp),
	})
}

// Update handles the update of a status
func (h *StatusHandler) Update(c *gin.Context) {
	id := ExtractIDParam(c)
	if id == "" {
		RespondWithError(c, http.StatusBadRequest, errors.New("missing or invalid ID"), "Missing or invalid ID")
		return
	}

	var req models.UpdateStatusRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		RespondWithError(c, http.StatusBadRequest, err, "Invalid request body")
		return
	}

	if err := h.service.Update(c.Request.Context(), id, req); err != nil {
		RespondWithError(c, http.StatusInternalServerError, err, "Failed to update status")
		return
	}

	c.Status(http.StatusNoContent)
}

// Delete handles the deletion of a status
func (h *StatusHandler) Delete(c *gin.Context) {
	id := ExtractIDParam(c)
	if id == "" {
		RespondWithError(c, http.StatusBadRequest, errors.New("missing or invalid ID"), "Missing or invalid ID")
		return
	}

	if err := h.service.Delete(c.Request.Context(), id); err != nil {
		RespondWithError(c, http.StatusInternalServerError, err, "Failed to delete status")
		return
	}

	c.Status(http.StatusNoContent)
}
