package services

import (
	"context"

	"apm/internal/models"
)

// UserService defines the service for user-related operations
type UserService interface {
	Create(ctx context.Context, req models.CreateUserRequest) (models.UserResponse, error)
	GetByID(ctx context.Context, id string) (models.UserResponse, error)
	List(ctx context.Context, limit, offset int) ([]models.UserResponse, error)
	Update(ctx context.Context, id string, req models.UpdateUserRequest) error
	Delete(ctx context.Context, id string) error
}

// UserGroupService defines the service for user group-related operations
type UserGroupService interface {
	Create(ctx context.Context, req models.CreateUserGroupRequest) (models.UserGroupResponse, error)
	GetByID(ctx context.Context, id string) (models.UserGroupResponse, error)
	List(ctx context.Context, limit, offset int) ([]models.UserGroupResponse, error)
	Update(ctx context.Context, id string, req models.UpdateUserGroupRequest) error
	Delete(ctx context.Context, id string) error
}

// StakeholderService defines the service for stakeholder-related operations
type StakeholderService interface {
	Create(ctx context.Context, req models.CreateStakeholderRequest) (models.StakeholderResponse, error)
	GetByID(ctx context.Context, id string) (models.StakeholderResponse, error)
	List(ctx context.Context, limit, offset int) ([]models.StakeholderResponse, error)
	Update(ctx context.Context, id string, req models.UpdateStakeholderRequest) error
	Delete(ctx context.Context, id string) error
}

// EntityService defines the service for entity-related operations
type EntityService interface {
	Create(ctx context.Context, req models.CreateEntityRequest) (models.EntityResponse, error)
	GetByID(ctx context.Context, id string) (models.EntityResponse, error)
	List(ctx context.Context, limit, offset int) ([]models.EntityResponse, error)
	Update(ctx context.Context, id string, req models.UpdateEntityRequest) error
	Delete(ctx context.Context, id string) error
}

// SoftwareService defines the service for software-related operations
type SoftwareService interface {
	Create(ctx context.Context, req models.CreateSoftwareRequest) (models.SoftwareResponse, error)
	GetByID(ctx context.Context, id string) (models.SoftwareResponse, error)
	List(ctx context.Context, limit, offset int) ([]models.SoftwareResponse, error)
	Update(ctx context.Context, id string, req models.UpdateSoftwareRequest) error
	Delete(ctx context.Context, id string) error
}

// FunctionalCategoryService defines the service for functional category-related operations
type FunctionalCategoryService interface {
	Create(ctx context.Context, req models.CreateFunctionalCategoryRequest) (models.FunctionalCategoryResponse, error)
	GetByID(ctx context.Context, id string) (models.FunctionalCategoryResponse, error)
	List(ctx context.Context, limit, offset int) ([]models.FunctionalCategoryResponse, error)
	Update(ctx context.Context, id string, req models.UpdateFunctionalCategoryRequest) error
	Delete(ctx context.Context, id string) error
}

// SoftwareGroupService defines the service for software group-related operations
type SoftwareGroupService interface {
	Create(ctx context.Context, req models.CreateSoftwareGroupRequest) (models.SoftwareGroupResponse, error)
	GetByID(ctx context.Context, id string) (models.SoftwareGroupResponse, error)
	List(ctx context.Context, limit, offset int) ([]models.SoftwareGroupResponse, error)
	Update(ctx context.Context, id string, req models.UpdateSoftwareGroupRequest) error
	Delete(ctx context.Context, id string) error
}

// StatusService defines the service for status-related operations
type StatusService interface {
	Create(ctx context.Context, req models.CreateStatusRequest) (models.StatusResponse, error)
	GetByID(ctx context.Context, id string) (models.StatusResponse, error)
	List(ctx context.Context, limit, offset int) ([]models.StatusResponse, error)
	Update(ctx context.Context, id string, req models.UpdateStatusRequest) error
	Delete(ctx context.Context, id string) error
}

// StatusLogService defines the service for status log-related operations
type StatusLogService interface {
	Create(ctx context.Context, req models.CreateStatusLogRequest) (models.StatusLogResponse, error)
	GetByID(ctx context.Context, id string) (models.StatusLogResponse, error)
	List(ctx context.Context, limit, offset int) ([]models.StatusLogResponse, error)
	Update(ctx context.Context, id string, req models.UpdateStatusLogRequest) error
	Delete(ctx context.Context, id string) error
}

// RankService defines the service for rank-related operations
type RankService interface {
	Create(ctx context.Context, req models.CreateRankRequest) (models.RankResponse, error)
	GetByID(ctx context.Context, id string) (models.RankResponse, error)
	List(ctx context.Context, limit, offset int) ([]models.RankResponse, error)
	Update(ctx context.Context, id string, req models.UpdateRankRequest) error
	Delete(ctx context.Context, id string) error
}

// NewsArticleService defines the service for news article-related operations
type NewsArticleService interface {
	Create(ctx context.Context, req models.CreateNewsArticleRequest) (models.NewsArticleResponse, error)
	GetByID(ctx context.Context, id string) (models.NewsArticleResponse, error)
	List(ctx context.Context, limit, offset int) ([]models.NewsArticleResponse, error)
	Update(ctx context.Context, id string, req models.UpdateNewsArticleRequest) error
	Delete(ctx context.Context, id string) error
}

// MediaService defines the service for media-related operations
type MediaService interface {
	Create(ctx context.Context, req models.CreateMediaRequest) (models.MediaResponse, error)
	GetByID(ctx context.Context, id string) (models.MediaResponse, error)
	List(ctx context.Context, limit, offset int) ([]models.MediaResponse, error)
	Update(ctx context.Context, id string, req models.UpdateMediaRequest) error
	Delete(ctx context.Context, id string) error
}

// ProductDocumentationService defines the service for product documentation-related operations
type ProductDocumentationService interface {
	Create(ctx context.Context, req models.CreateProductDocumentationRequest) (models.ProductDocumentationResponse, error)
	GetByID(ctx context.Context, id string) (models.ProductDocumentationResponse, error)
	List(ctx context.Context, limit, offset int) ([]models.ProductDocumentationResponse, error)
	Update(ctx context.Context, id string, req models.UpdateProductDocumentationRequest) error
	Delete(ctx context.Context, id string) error
}

// LogService defines the service for log-related operations
type LogService interface {
	Create(ctx context.Context, req models.CreateLogRequest) (models.LogResponse, error)
	GetByID(ctx context.Context, id string) (models.LogResponse, error)
	List(ctx context.Context, limit, offset int) ([]models.LogResponse, error)
	Delete(ctx context.Context, id string) error
}
