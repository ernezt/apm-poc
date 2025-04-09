package handlers

import (
	"net/http"
	"time"

	"apm/internal/models"
	"apm/internal/services"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

// AuthHandler handles authentication-related requests
type AuthHandler struct {
	userService services.UserService
	jwtSecret   []byte
}

// NewAuthHandler creates a new auth handler
func NewAuthHandler(userService services.UserService, jwtSecret string) *AuthHandler {
	return &AuthHandler{
		userService: userService,
		jwtSecret:   []byte(jwtSecret),
	}
}

// Login handles user login
func (h *AuthHandler) Login(c *gin.Context) {
	var req models.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	// TODO: Implement actual user authentication
	// For now, we'll use a mock user for testing
	if req.Email == "admin@example.com" && req.Password == "password123" {
		// Create a mock user response
		user := models.UserResponse{
			ID:         "1",
			Email:      req.Email,
			FirstName:  "Admin",
			LastName:   "User",
			Role:       models.RoleOrganizationAdmin,
			MFAEnabled: false,
			CreatedAt:  time.Now(),
			UpdatedAt:  time.Now(),
		}

		// Generate JWT token
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"sub": user.ID,
			"exp": time.Now().Add(time.Hour * 24).Unix(),
		})

		tokenString, err := token.SignedString(h.jwtSecret)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
			return
		}

		c.JSON(http.StatusOK, models.LoginResponse{
			User:        user,
			AccessToken: tokenString,
		})
		return
	}

	c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
}

// RegisterRoutes registers the auth routes
func (h *AuthHandler) RegisterRoutes(router *gin.RouterGroup) {
	auth := router.Group("/auth")
	{
		auth.POST("/login", h.Login)
	}
}
