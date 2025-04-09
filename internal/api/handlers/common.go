package handlers

import (
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

// ErrorResponse represents a standard error response
type ErrorResponse struct {
	Error   string `json:"error"`
	Message string `json:"message,omitempty"`
	Code    int    `json:"code"`
}

// RespondWithError sends a JSON error response
func RespondWithError(c *gin.Context, status int, err error, message string) {
	c.JSON(status, ErrorResponse{
		Error:   err.Error(),
		Message: message,
		Code:    status,
	})
}

// ExtractIDParam extracts an ID parameter from the request URL
func ExtractIDParam(c *gin.Context) string {
	return strings.TrimSpace(c.Param("id"))
}

// QueryParam gets a query parameter with a default value
func QueryParam(c *gin.Context, key string, defaultValue string) string {
	value := c.Query(key)
	if value == "" {
		return defaultValue
	}
	return value
}

// SetPagination prepares pagination parameters from the request
func SetPagination(c *gin.Context) (limit, offset int) {
	limitStr := c.DefaultQuery("limit", "10")
	offsetStr := c.DefaultQuery("offset", "0")

	// Parse values with defaults
	limit = 10
	offset = 0

	// Try to parse limit
	if val, err := strconv.Atoi(limitStr); err == nil {
		limit = val
	}

	// Try to parse offset
	if val, err := strconv.Atoi(offsetStr); err == nil {
		offset = val
	}

	// Ensure reasonable defaults
	if limit < 1 {
		limit = 10
	}
	if limit > 100 {
		limit = 100
	}
	if offset < 0 {
		offset = 0
	}

	return limit, offset
}
