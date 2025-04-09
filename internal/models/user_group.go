package models

import (
	"time"
)

// UserGroup represents a group of users in the system
type UserGroup struct {
	ID          string    `json:"id"`
	DisplayName string    `json:"display_name"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// CreateUserGroupRequest represents the request to create a new user group
type CreateUserGroupRequest struct {
	DisplayName string `json:"display_name" validate:"required"`
}

// UpdateUserGroupRequest represents the request to update a user group
type UpdateUserGroupRequest struct {
	DisplayName string `json:"display_name,omitempty"`
}

// UserGroupResponse represents the response when returning user group data
type UserGroupResponse struct {
	ID          string    `json:"id"`
	DisplayName string    `json:"display_name"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// UserToGroup represents the many-to-many relationship between users and groups
type UserToGroup struct {
	UserID      string    `json:"user_id"`
	UserGroupID string    `json:"user_group_id"`
	CreatedAt   time.Time `json:"created_at"`
}
