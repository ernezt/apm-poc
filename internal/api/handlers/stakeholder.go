package handlers

import (
	"errors"
	"net/http"

	"apm/internal/models"
	"apm/internal/services"

	"github.com/gin-gonic/gin"
)

// StakeholderHandler handles HTTP requests for stakeholders
type StakeholderHandler struct {
	service services.StakeholderService
}

// NewStakeholderHandler creates a new stakeholder handler
func NewStakeholderHandler(service services.StakeholderService) *StakeholderHandler {
	return &StakeholderHandler{
		service: service,
	}
}

// Register registers the routes for stakeholders
func (h *StakeholderHandler) Register(router *gin.RouterGroup) {
	stakeholders := router.Group("/stakeholders")
	{
		stakeholders.POST("", h.Create)
		stakeholders.GET("", h.List)
		stakeholders.GET("/:id", h.GetByID)
		stakeholders.PUT("/:id", h.Update)
		stakeholders.DELETE("/:id", h.Delete)
	}
}

// Create handles the creation of a new stakeholder
func (h *StakeholderHandler) Create(c *gin.Context) {
	var req models.CreateStakeholderRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		RespondWithError(c, http.StatusBadRequest, err, "Invalid request body")
		return
	}

	resp, err := h.service.Create(c.Request.Context(), req)
	if err != nil {
		RespondWithError(c, http.StatusInternalServerError, err, "Failed to create stakeholder")
		return
	}

	c.JSON(http.StatusCreated, resp)
}

// GetByID handles the retrieval of a stakeholder by ID
func (h *StakeholderHandler) GetByID(c *gin.Context) {
	id := ExtractIDParam(c)
	if id == "" {
		RespondWithError(c, http.StatusBadRequest, errors.New("missing or invalid ID"), "Missing or invalid ID")
		return
	}

	resp, err := h.service.GetByID(c.Request.Context(), id)
	if err != nil {
		RespondWithError(c, http.StatusNotFound, err, "Stakeholder not found")
		return
	}

	c.JSON(http.StatusOK, resp)
}

// List handles the retrieval of a list of stakeholders
func (h *StakeholderHandler) List(c *gin.Context) {
	limit, offset := SetPagination(c)

	resp, err := h.service.List(c.Request.Context(), limit, offset)
	if err != nil {
		RespondWithError(c, http.StatusInternalServerError, err, "Failed to retrieve stakeholders")
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":   resp,
		"limit":  limit,
		"offset": offset,
		"count":  len(resp),
	})
}

// Update handles the update of a stakeholder
func (h *StakeholderHandler) Update(c *gin.Context) {
	id := ExtractIDParam(c)
	if id == "" {
		RespondWithError(c, http.StatusBadRequest, errors.New("missing or invalid ID"), "Missing or invalid ID")
		return
	}

	var req models.UpdateStakeholderRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		RespondWithError(c, http.StatusBadRequest, err, "Invalid request body")
		return
	}

	if err := h.service.Update(c.Request.Context(), id, req); err != nil {
		RespondWithError(c, http.StatusInternalServerError, err, "Failed to update stakeholder")
		return
	}

	c.Status(http.StatusNoContent)
}

// Delete handles the deletion of a stakeholder
func (h *StakeholderHandler) Delete(c *gin.Context) {
	id := ExtractIDParam(c)
	if id == "" {
		RespondWithError(c, http.StatusBadRequest, errors.New("missing or invalid ID"), "Missing or invalid ID")
		return
	}

	if err := h.service.Delete(c.Request.Context(), id); err != nil {
		RespondWithError(c, http.StatusInternalServerError, err, "Failed to delete stakeholder")
		return
	}

	c.Status(http.StatusNoContent)
}
