package handlers

import (
	"errors"
	"net/http"

	"apm/internal/models"
	"apm/internal/services"

	"github.com/gin-gonic/gin"
)

// RankHandler handles HTTP requests for ranks
type RankHandler struct {
	service services.RankService
}

// NewRankHandler creates a new rank handler
func NewRankHandler(service services.RankService) *RankHandler {
	return &RankHandler{
		service: service,
	}
}

// Register registers the routes for ranks
func (h *RankHandler) Register(router *gin.RouterGroup) {
	ranks := router.Group("/ranks")
	{
		ranks.POST("", h.Create)
		ranks.GET("", h.List)
		ranks.GET("/:id", h.GetByID)
		ranks.PUT("/:id", h.Update)
		ranks.DELETE("/:id", h.Delete)
	}
}

// Create handles the creation of a new rank
func (h *RankHandler) Create(c *gin.Context) {
	var req models.CreateRankRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		RespondWithError(c, http.StatusBadRequest, err, "Invalid request body")
		return
	}

	resp, err := h.service.Create(c.Request.Context(), req)
	if err != nil {
		RespondWithError(c, http.StatusInternalServerError, err, "Failed to create rank")
		return
	}

	c.JSON(http.StatusCreated, resp)
}

// GetByID handles the retrieval of a rank by ID
func (h *RankHandler) GetByID(c *gin.Context) {
	id := ExtractIDParam(c)
	if id == "" {
		RespondWithError(c, http.StatusBadRequest, errors.New("missing or invalid ID"), "Missing or invalid ID")
		return
	}

	resp, err := h.service.GetByID(c.Request.Context(), id)
	if err != nil {
		RespondWithError(c, http.StatusNotFound, err, "Rank not found")
		return
	}

	c.JSON(http.StatusOK, resp)
}

// List handles the retrieval of a list of ranks
func (h *RankHandler) List(c *gin.Context) {
	limit, offset := SetPagination(c)

	resp, err := h.service.List(c.Request.Context(), limit, offset)
	if err != nil {
		RespondWithError(c, http.StatusInternalServerError, err, "Failed to retrieve ranks")
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":   resp,
		"limit":  limit,
		"offset": offset,
		"count":  len(resp),
	})
}

// Update handles the update of a rank
func (h *RankHandler) Update(c *gin.Context) {
	id := ExtractIDParam(c)
	if id == "" {
		RespondWithError(c, http.StatusBadRequest, errors.New("missing or invalid ID"), "Missing or invalid ID")
		return
	}

	var req models.UpdateRankRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		RespondWithError(c, http.StatusBadRequest, err, "Invalid request body")
		return
	}

	if err := h.service.Update(c.Request.Context(), id, req); err != nil {
		RespondWithError(c, http.StatusInternalServerError, err, "Failed to update rank")
		return
	}

	c.Status(http.StatusNoContent)
}

// Delete handles the deletion of a rank
func (h *RankHandler) Delete(c *gin.Context) {
	id := ExtractIDParam(c)
	if id == "" {
		RespondWithError(c, http.StatusBadRequest, errors.New("missing or invalid ID"), "Missing or invalid ID")
		return
	}

	if err := h.service.Delete(c.Request.Context(), id); err != nil {
		RespondWithError(c, http.StatusInternalServerError, err, "Failed to delete rank")
		return
	}

	c.Status(http.StatusNoContent)
}
