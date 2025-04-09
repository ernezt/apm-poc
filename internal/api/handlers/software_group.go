package handlers

import (
	"errors"
	"net/http"

	"apm/internal/models"
	"apm/internal/services"

	"github.com/gin-gonic/gin"
)

// SoftwareGroupHandler handles HTTP requests for software groups
type SoftwareGroupHandler struct {
	service services.SoftwareGroupService
}

// NewSoftwareGroupHandler creates a new software group handler
func NewSoftwareGroupHandler(service services.SoftwareGroupService) *SoftwareGroupHandler {
	return &SoftwareGroupHandler{
		service: service,
	}
}

// Register registers the routes for software groups
func (h *SoftwareGroupHandler) Register(router *gin.RouterGroup) {
	groups := router.Group("/software-groups")
	{
		groups.POST("", h.Create)
		groups.GET("", h.List)
		groups.GET("/:id", h.GetByID)
		groups.PUT("/:id", h.Update)
		groups.DELETE("/:id", h.Delete)
	}
}

// Create handles the creation of a new software group
func (h *SoftwareGroupHandler) Create(c *gin.Context) {
	var req models.CreateSoftwareGroupRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		RespondWithError(c, http.StatusBadRequest, err, "Invalid request body")
		return
	}

	resp, err := h.service.Create(c.Request.Context(), req)
	if err != nil {
		RespondWithError(c, http.StatusInternalServerError, err, "Failed to create software group")
		return
	}

	c.JSON(http.StatusCreated, resp)
}

// GetByID handles the retrieval of a software group by ID
func (h *SoftwareGroupHandler) GetByID(c *gin.Context) {
	id := ExtractIDParam(c)
	if id == "" {
		RespondWithError(c, http.StatusBadRequest, errors.New("missing or invalid ID"), "Missing or invalid ID")
		return
	}

	resp, err := h.service.GetByID(c.Request.Context(), id)
	if err != nil {
		RespondWithError(c, http.StatusNotFound, err, "Software group not found")
		return
	}

	c.JSON(http.StatusOK, resp)
}

// List handles the retrieval of a list of software groups
func (h *SoftwareGroupHandler) List(c *gin.Context) {
	limit, offset := SetPagination(c)

	resp, err := h.service.List(c.Request.Context(), limit, offset)
	if err != nil {
		RespondWithError(c, http.StatusInternalServerError, err, "Failed to retrieve software groups")
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":   resp,
		"limit":  limit,
		"offset": offset,
		"count":  len(resp),
	})
}

// Update handles the update of a software group
func (h *SoftwareGroupHandler) Update(c *gin.Context) {
	id := ExtractIDParam(c)
	if id == "" {
		RespondWithError(c, http.StatusBadRequest, errors.New("missing or invalid ID"), "Missing or invalid ID")
		return
	}

	var req models.UpdateSoftwareGroupRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		RespondWithError(c, http.StatusBadRequest, err, "Invalid request body")
		return
	}

	if err := h.service.Update(c.Request.Context(), id, req); err != nil {
		RespondWithError(c, http.StatusInternalServerError, err, "Failed to update software group")
		return
	}

	c.Status(http.StatusNoContent)
}

// Delete handles the deletion of a software group
func (h *SoftwareGroupHandler) Delete(c *gin.Context) {
	id := ExtractIDParam(c)
	if id == "" {
		RespondWithError(c, http.StatusBadRequest, errors.New("missing or invalid ID"), "Missing or invalid ID")
		return
	}

	if err := h.service.Delete(c.Request.Context(), id); err != nil {
		RespondWithError(c, http.StatusInternalServerError, err, "Failed to delete software group")
		return
	}

	c.Status(http.StatusNoContent)
}
