package handlers

import (
	"errors"
	"net/http"

	"apm/internal/models"
	"apm/internal/services"

	"github.com/gin-gonic/gin"
)

// StatusLogHandler handles HTTP requests for status logs
type StatusLogHandler struct {
	service services.StatusLogService
}

// NewStatusLogHandler creates a new status log handler
func NewStatusLogHandler(service services.StatusLogService) *StatusLogHandler {
	return &StatusLogHandler{
		service: service,
	}
}

// Register registers the routes for status logs
func (h *StatusLogHandler) Register(router *gin.RouterGroup) {
	logs := router.Group("/status-logs")
	{
		logs.POST("", h.Create)
		logs.GET("", h.List)
		logs.GET("/:id", h.GetByID)
		logs.PUT("/:id", h.Update)
		logs.DELETE("/:id", h.Delete)
	}
}

// Create handles the creation of a new status log
func (h *StatusLogHandler) Create(c *gin.Context) {
	var req models.CreateStatusLogRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		RespondWithError(c, http.StatusBadRequest, err, "Invalid request body")
		return
	}

	resp, err := h.service.Create(c.Request.Context(), req)
	if err != nil {
		RespondWithError(c, http.StatusInternalServerError, err, "Failed to create status log")
		return
	}

	c.JSON(http.StatusCreated, resp)
}

// GetByID handles the retrieval of a status log by ID
func (h *StatusLogHandler) GetByID(c *gin.Context) {
	id := ExtractIDParam(c)
	if id == "" {
		RespondWithError(c, http.StatusBadRequest, errors.New("missing or invalid ID"), "Missing or invalid ID")
		return
	}

	resp, err := h.service.GetByID(c.Request.Context(), id)
	if err != nil {
		RespondWithError(c, http.StatusNotFound, err, "Status log not found")
		return
	}

	c.JSON(http.StatusOK, resp)
}

// List handles the retrieval of a list of status logs
func (h *StatusLogHandler) List(c *gin.Context) {
	limit, offset := SetPagination(c)

	resp, err := h.service.List(c.Request.Context(), limit, offset)
	if err != nil {
		RespondWithError(c, http.StatusInternalServerError, err, "Failed to retrieve status logs")
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":   resp,
		"limit":  limit,
		"offset": offset,
		"count":  len(resp),
	})
}

// Update handles the update of a status log
func (h *StatusLogHandler) Update(c *gin.Context) {
	id := ExtractIDParam(c)
	if id == "" {
		RespondWithError(c, http.StatusBadRequest, errors.New("missing or invalid ID"), "Missing or invalid ID")
		return
	}

	var req models.UpdateStatusLogRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		RespondWithError(c, http.StatusBadRequest, err, "Invalid request body")
		return
	}

	if err := h.service.Update(c.Request.Context(), id, req); err != nil {
		RespondWithError(c, http.StatusInternalServerError, err, "Failed to update status log")
		return
	}

	c.Status(http.StatusNoContent)
}

// Delete handles the deletion of a status log
func (h *StatusLogHandler) Delete(c *gin.Context) {
	id := ExtractIDParam(c)
	if id == "" {
		RespondWithError(c, http.StatusBadRequest, errors.New("missing or invalid ID"), "Missing or invalid ID")
		return
	}

	if err := h.service.Delete(c.Request.Context(), id); err != nil {
		RespondWithError(c, http.StatusInternalServerError, err, "Failed to delete status log")
		return
	}

	c.Status(http.StatusNoContent)
}
