package handlers

import (
	"errors"
	"net/http"

	"apm/internal/models"
	"apm/internal/services"

	"github.com/gin-gonic/gin"
)

// MediaHandler handles HTTP requests for media
type MediaHandler struct {
	service services.MediaService
}

// NewMediaHandler creates a new media handler
func NewMediaHandler(service services.MediaService) *MediaHandler {
	return &MediaHandler{
		service: service,
	}
}

// Register registers the routes for media
func (h *MediaHandler) Register(router *gin.RouterGroup) {
	media := router.Group("/media")
	{
		media.POST("", h.Create)
		media.GET("", h.List)
		media.GET("/:id", h.GetByID)
		media.PUT("/:id", h.Update)
		media.DELETE("/:id", h.Delete)
	}
}

// Create handles the creation of new media
func (h *MediaHandler) Create(c *gin.Context) {
	var req models.CreateMediaRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		RespondWithError(c, http.StatusBadRequest, err, "Invalid request body")
		return
	}

	resp, err := h.service.Create(c.Request.Context(), req)
	if err != nil {
		RespondWithError(c, http.StatusInternalServerError, err, "Failed to create media")
		return
	}

	c.JSON(http.StatusCreated, resp)
}

// GetByID handles the retrieval of media by ID
func (h *MediaHandler) GetByID(c *gin.Context) {
	id := ExtractIDParam(c)
	if id == "" {
		RespondWithError(c, http.StatusBadRequest, errors.New("missing or invalid ID"), "Missing or invalid ID")
		return
	}

	resp, err := h.service.GetByID(c.Request.Context(), id)
	if err != nil {
		RespondWithError(c, http.StatusNotFound, err, "Media not found")
		return
	}

	c.JSON(http.StatusOK, resp)
}

// List handles the retrieval of a list of media
func (h *MediaHandler) List(c *gin.Context) {
	limit, offset := SetPagination(c)

	resp, err := h.service.List(c.Request.Context(), limit, offset)
	if err != nil {
		RespondWithError(c, http.StatusInternalServerError, err, "Failed to retrieve media list")
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":   resp,
		"limit":  limit,
		"offset": offset,
		"count":  len(resp),
	})
}

// Update handles the update of media
func (h *MediaHandler) Update(c *gin.Context) {
	id := ExtractIDParam(c)
	if id == "" {
		RespondWithError(c, http.StatusBadRequest, errors.New("missing or invalid ID"), "Missing or invalid ID")
		return
	}

	var req models.UpdateMediaRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		RespondWithError(c, http.StatusBadRequest, err, "Invalid request body")
		return
	}

	if err := h.service.Update(c.Request.Context(), id, req); err != nil {
		RespondWithError(c, http.StatusInternalServerError, err, "Failed to update media")
		return
	}

	c.Status(http.StatusNoContent)
}

// Delete handles the deletion of media
func (h *MediaHandler) Delete(c *gin.Context) {
	id := ExtractIDParam(c)
	if id == "" {
		RespondWithError(c, http.StatusBadRequest, errors.New("missing or invalid ID"), "Missing or invalid ID")
		return
	}

	if err := h.service.Delete(c.Request.Context(), id); err != nil {
		RespondWithError(c, http.StatusInternalServerError, err, "Failed to delete media")
		return
	}

	c.Status(http.StatusNoContent)
}
