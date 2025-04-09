package handlers

import (
	"errors"
	"net/http"

	"apm/internal/models"
	"apm/internal/services"

	"github.com/gin-gonic/gin"
)

// UserGroupHandler handles HTTP requests for user groups
type UserGroupHandler struct {
	service services.UserGroupService
}

// NewUserGroupHandler creates a new user group handler
func NewUserGroupHandler(service services.UserGroupService) *UserGroupHandler {
	return &UserGroupHandler{
		service: service,
	}
}

// Register registers the routes for user groups
func (h *UserGroupHandler) Register(router *gin.RouterGroup) {
	groups := router.Group("/user-groups")
	{
		groups.POST("", h.Create)
		groups.GET("", h.List)
		groups.GET("/:id", h.GetByID)
		groups.PUT("/:id", h.Update)
		groups.DELETE("/:id", h.Delete)
	}
}

// Create handles the creation of a new user group
func (h *UserGroupHandler) Create(c *gin.Context) {
	var req models.CreateUserGroupRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		RespondWithError(c, http.StatusBadRequest, err, "Invalid request body")
		return
	}

	resp, err := h.service.Create(c.Request.Context(), req)
	if err != nil {
		RespondWithError(c, http.StatusInternalServerError, err, "Failed to create user group")
		return
	}

	c.JSON(http.StatusCreated, resp)
}

// GetByID handles the retrieval of a user group by ID
func (h *UserGroupHandler) GetByID(c *gin.Context) {
	id := ExtractIDParam(c)
	if id == "" {
		RespondWithError(c, http.StatusBadRequest, errors.New("missing or invalid ID"), "Missing or invalid ID")
		return
	}

	resp, err := h.service.GetByID(c.Request.Context(), id)
	if err != nil {
		RespondWithError(c, http.StatusNotFound, err, "User group not found")
		return
	}

	c.JSON(http.StatusOK, resp)
}

// List handles the retrieval of a list of user groups
func (h *UserGroupHandler) List(c *gin.Context) {
	limit, offset := SetPagination(c)

	resp, err := h.service.List(c.Request.Context(), limit, offset)
	if err != nil {
		RespondWithError(c, http.StatusInternalServerError, err, "Failed to retrieve user groups")
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":   resp,
		"limit":  limit,
		"offset": offset,
		"count":  len(resp),
	})
}

// Update handles the update of a user group
func (h *UserGroupHandler) Update(c *gin.Context) {
	id := ExtractIDParam(c)
	if id == "" {
		RespondWithError(c, http.StatusBadRequest, errors.New("missing or invalid ID"), "Missing or invalid ID")
		return
	}

	var req models.UpdateUserGroupRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		RespondWithError(c, http.StatusBadRequest, err, "Invalid request body")
		return
	}

	if err := h.service.Update(c.Request.Context(), id, req); err != nil {
		RespondWithError(c, http.StatusInternalServerError, err, "Failed to update user group")
		return
	}

	c.Status(http.StatusNoContent)
}

// Delete handles the deletion of a user group
func (h *UserGroupHandler) Delete(c *gin.Context) {
	id := ExtractIDParam(c)
	if id == "" {
		RespondWithError(c, http.StatusBadRequest, errors.New("missing or invalid ID"), "Missing or invalid ID")
		return
	}

	if err := h.service.Delete(c.Request.Context(), id); err != nil {
		RespondWithError(c, http.StatusInternalServerError, err, "Failed to delete user group")
		return
	}

	c.Status(http.StatusNoContent)
}
