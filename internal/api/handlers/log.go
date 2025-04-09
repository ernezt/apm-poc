package handlers

import (
	"errors"
	"net/http"

	"apm/internal/models"
	"apm/internal/services"

	"github.com/gin-gonic/gin"
)

// LogHandler handles HTTP requests for logs
type LogHandler struct {
	service services.LogService
}

// NewLogHandler creates a new log handler
func NewLogHandler(service services.LogService) *LogHandler {
	return &LogHandler{
		service: service,
	}
}

// Register registers the routes for logs
func (h *LogHandler) Register(router *gin.RouterGroup) {
	logs := router.Group("/logs")
	{
		logs.POST("", h.Create)
		logs.GET("", h.List)
		logs.GET("/:id", h.GetByID)
		logs.DELETE("/:id", h.Delete)
	}
}

// Create handles the creation of a new log
func (h *LogHandler) Create(c *gin.Context) {
	var req models.CreateLogRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		RespondWithError(c, http.StatusBadRequest, err, "Invalid request body")
		return
	}

	resp, err := h.service.Create(c.Request.Context(), req)
	if err != nil {
		RespondWithError(c, http.StatusInternalServerError, err, "Failed to create log")
		return
	}

	c.JSON(http.StatusCreated, resp)
}

// GetByID handles the retrieval of a log by ID
func (h *LogHandler) GetByID(c *gin.Context) {
	id := ExtractIDParam(c)
	if id == "" {
		RespondWithError(c, http.StatusBadRequest, errors.New("missing or invalid ID"), "Missing or invalid ID")
		return
	}

	resp, err := h.service.GetByID(c.Request.Context(), id)
	if err != nil {
		RespondWithError(c, http.StatusNotFound, err, "Log not found")
		return
	}

	c.JSON(http.StatusOK, resp)
}

// List handles the retrieval of a list of logs
func (h *LogHandler) List(c *gin.Context) {
	limit, offset := SetPagination(c)

	resp, err := h.service.List(c.Request.Context(), limit, offset)
	if err != nil {
		RespondWithError(c, http.StatusInternalServerError, err, "Failed to retrieve logs")
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":   resp,
		"limit":  limit,
		"offset": offset,
		"count":  len(resp),
	})
}

// Delete handles the deletion of a log
func (h *LogHandler) Delete(c *gin.Context) {
	id := ExtractIDParam(c)
	if id == "" {
		RespondWithError(c, http.StatusBadRequest, errors.New("missing or invalid ID"), "Missing or invalid ID")
		return
	}

	if err := h.service.Delete(c.Request.Context(), id); err != nil {
		RespondWithError(c, http.StatusInternalServerError, err, "Failed to delete log")
		return
	}

	c.Status(http.StatusNoContent)
}
