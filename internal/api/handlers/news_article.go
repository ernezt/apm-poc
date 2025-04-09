package handlers

import (
	"errors"
	"net/http"

	"apm/internal/models"
	"apm/internal/services"

	"github.com/gin-gonic/gin"
)

// NewsArticleHandler handles HTTP requests for news articles
type NewsArticleHandler struct {
	service services.NewsArticleService
}

// NewNewsArticleHandler creates a new news article handler
func NewNewsArticleHandler(service services.NewsArticleService) *NewsArticleHandler {
	return &NewsArticleHandler{
		service: service,
	}
}

// Register registers the routes for news articles
func (h *NewsArticleHandler) Register(router *gin.RouterGroup) {
	news := router.Group("/news")
	{
		news.POST("", h.Create)
		news.GET("", h.List)
		news.GET("/:id", h.GetByID)
		news.PUT("/:id", h.Update)
		news.DELETE("/:id", h.Delete)
	}
}

// Create handles the creation of a new news article
func (h *NewsArticleHandler) Create(c *gin.Context) {
	var req models.CreateNewsArticleRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		RespondWithError(c, http.StatusBadRequest, err, "Invalid request body")
		return
	}

	resp, err := h.service.Create(c.Request.Context(), req)
	if err != nil {
		RespondWithError(c, http.StatusInternalServerError, err, "Failed to create news article")
		return
	}

	c.JSON(http.StatusCreated, resp)
}

// GetByID handles the retrieval of a news article by ID
func (h *NewsArticleHandler) GetByID(c *gin.Context) {
	id := ExtractIDParam(c)
	if id == "" {
		RespondWithError(c, http.StatusBadRequest, errors.New("missing or invalid ID"), "Missing or invalid ID")
		return
	}

	resp, err := h.service.GetByID(c.Request.Context(), id)
	if err != nil {
		RespondWithError(c, http.StatusNotFound, err, "News article not found")
		return
	}

	c.JSON(http.StatusOK, resp)
}

// List handles the retrieval of a list of news articles
func (h *NewsArticleHandler) List(c *gin.Context) {
	limit, offset := SetPagination(c)

	resp, err := h.service.List(c.Request.Context(), limit, offset)
	if err != nil {
		RespondWithError(c, http.StatusInternalServerError, err, "Failed to retrieve news articles")
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":   resp,
		"limit":  limit,
		"offset": offset,
		"count":  len(resp),
	})
}

// Update handles the update of a news article
func (h *NewsArticleHandler) Update(c *gin.Context) {
	id := ExtractIDParam(c)
	if id == "" {
		RespondWithError(c, http.StatusBadRequest, errors.New("missing or invalid ID"), "Missing or invalid ID")
		return
	}

	var req models.UpdateNewsArticleRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		RespondWithError(c, http.StatusBadRequest, err, "Invalid request body")
		return
	}

	if err := h.service.Update(c.Request.Context(), id, req); err != nil {
		RespondWithError(c, http.StatusInternalServerError, err, "Failed to update news article")
		return
	}

	c.Status(http.StatusNoContent)
}

// Delete handles the deletion of a news article
func (h *NewsArticleHandler) Delete(c *gin.Context) {
	id := ExtractIDParam(c)
	if id == "" {
		RespondWithError(c, http.StatusBadRequest, errors.New("missing or invalid ID"), "Missing or invalid ID")
		return
	}

	if err := h.service.Delete(c.Request.Context(), id); err != nil {
		RespondWithError(c, http.StatusInternalServerError, err, "Failed to delete news article")
		return
	}

	c.Status(http.StatusNoContent)
}
