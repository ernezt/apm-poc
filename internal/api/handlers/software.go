package handlers

import (
	"errors"
	"net/http"

	"apm/internal/models"
	"apm/internal/services"

	"github.com/gin-gonic/gin"
)

// SoftwareHandler handles HTTP requests for software
type SoftwareHandler struct {
	service services.SoftwareService
}

// NewSoftwareHandler creates a new software handler
func NewSoftwareHandler(service services.SoftwareService) *SoftwareHandler {
	return &SoftwareHandler{
		service: service,
	}
}

// Register registers the routes for software
func (h *SoftwareHandler) Register(router *gin.RouterGroup) {
	software := router.Group("/software")
	{
		software.POST("", h.Create)
		software.GET("", h.List)
		software.GET("/:id", h.GetByID)
		software.PUT("/:id", h.Update)
		software.DELETE("/:id", h.Delete)
	}
}

// Create handles the creation of new software
func (h *SoftwareHandler) Create(c *gin.Context) {
	var req models.CreateSoftwareRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		RespondWithError(c, http.StatusBadRequest, err, "Invalid request body")
		return
	}

	resp, err := h.service.Create(c.Request.Context(), req)
	if err != nil {
		RespondWithError(c, http.StatusInternalServerError, err, "Failed to create software")
		return
	}

	c.JSON(http.StatusCreated, resp)
}

// GetByID handles the retrieval of software by ID
func (h *SoftwareHandler) GetByID(c *gin.Context) {
	id := ExtractIDParam(c)
	if id == "" {
		RespondWithError(c, http.StatusBadRequest, errors.New("missing or invalid ID"), "Missing or invalid ID")
		return
	}

	resp, err := h.service.GetByID(c.Request.Context(), id)
	if err != nil {
		RespondWithError(c, http.StatusNotFound, err, "Software not found")
		return
	}

	c.JSON(http.StatusOK, resp)
}

// List handles the retrieval of a list of software
func (h *SoftwareHandler) List(c *gin.Context) {
	limit, offset := SetPagination(c)

	resp, err := h.service.List(c.Request.Context(), limit, offset)
	if err != nil {
		RespondWithError(c, http.StatusInternalServerError, err, "Failed to retrieve software list")
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":   resp,
		"limit":  limit,
		"offset": offset,
		"count":  len(resp),
	})
}

// Update handles the update of software
func (h *SoftwareHandler) Update(c *gin.Context) {
	id := ExtractIDParam(c)
	if id == "" {
		RespondWithError(c, http.StatusBadRequest, errors.New("missing or invalid ID"), "Missing or invalid ID")
		return
	}

	var req models.UpdateSoftwareRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		RespondWithError(c, http.StatusBadRequest, err, "Invalid request body")
		return
	}

	if err := h.service.Update(c.Request.Context(), id, req); err != nil {
		RespondWithError(c, http.StatusInternalServerError, err, "Failed to update software")
		return
	}

	c.Status(http.StatusNoContent)
}

// Delete handles the deletion of software
func (h *SoftwareHandler) Delete(c *gin.Context) {
	id := ExtractIDParam(c)
	if id == "" {
		RespondWithError(c, http.StatusBadRequest, errors.New("missing or invalid ID"), "Missing or invalid ID")
		return
	}

	if err := h.service.Delete(c.Request.Context(), id); err != nil {
		RespondWithError(c, http.StatusInternalServerError, err, "Failed to delete software")
		return
	}

	c.Status(http.StatusNoContent)
}
