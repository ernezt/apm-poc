package repository

import (
	"context"

	"apm/internal/models"
)

// UserRepository defines the interface for user-related database operations
type UserRepository interface {
	Create(ctx context.Context, user models.User) (models.User, error)
	GetByID(ctx context.Context, id string) (models.User, error)
	GetByEmail(ctx context.Context, email string) (models.User, error)
	List(ctx context.Context, limit, offset int) ([]models.User, error)
	Update(ctx context.Context, user models.User) error
	Delete(ctx context.Context, id string) error
}

// UserGroupRepository defines the interface for user group-related database operations
type UserGroupRepository interface {
	Create(ctx context.Context, group models.UserGroup) (models.UserGroup, error)
	GetByID(ctx context.Context, id string) (models.UserGroup, error)
	List(ctx context.Context, limit, offset int) ([]models.UserGroup, error)
	Update(ctx context.Context, group models.UserGroup) error
	Delete(ctx context.Context, id string) error
}

// StakeholderRepository defines the interface for stakeholder-related database operations
type StakeholderRepository interface {
	Create(ctx context.Context, stakeholder models.Stakeholder) (models.Stakeholder, error)
	GetByID(ctx context.Context, id string) (models.Stakeholder, error)
	GetByUserID(ctx context.Context, userID string) ([]models.Stakeholder, error)
	List(ctx context.Context, limit, offset int) ([]models.Stakeholder, error)
	Update(ctx context.Context, stakeholder models.Stakeholder) error
	Delete(ctx context.Context, id string) error
}

// EntityRepository defines the interface for entity-related database operations
type EntityRepository interface {
	Create(ctx context.Context, entity models.Entity) (models.Entity, error)
	GetByID(ctx context.Context, id string) (models.Entity, error)
	List(ctx context.Context, limit, offset int) ([]models.Entity, error)
	Update(ctx context.Context, entity models.Entity) error
	Delete(ctx context.Context, id string) error
}

// SoftwareRepository defines the interface for software-related database operations
type SoftwareRepository interface {
	Create(ctx context.Context, software models.Software) (models.Software, error)
	GetByID(ctx context.Context, id string) (models.Software, error)
	List(ctx context.Context, limit, offset int) ([]models.Software, error)
	Update(ctx context.Context, software models.Software) error
	Delete(ctx context.Context, id string) error
}

// FunctionalCategoryRepository defines the interface for functional category-related database operations
type FunctionalCategoryRepository interface {
	Create(ctx context.Context, category models.FunctionalCategory) (models.FunctionalCategory, error)
	GetByID(ctx context.Context, id string) (models.FunctionalCategory, error)
	List(ctx context.Context, limit, offset int) ([]models.FunctionalCategory, error)
	Update(ctx context.Context, category models.FunctionalCategory) error
	Delete(ctx context.Context, id string) error
}

// SoftwareGroupRepository defines the interface for software group-related database operations
type SoftwareGroupRepository interface {
	Create(ctx context.Context, group models.SoftwareGroup) (models.SoftwareGroup, error)
	GetByID(ctx context.Context, id string) (models.SoftwareGroup, error)
	List(ctx context.Context, limit, offset int) ([]models.SoftwareGroup, error)
	Update(ctx context.Context, group models.SoftwareGroup) error
	Delete(ctx context.Context, id string) error
}

// StatusRepository defines the interface for status-related database operations
type StatusRepository interface {
	Create(ctx context.Context, status models.Status) (models.Status, error)
	GetByID(ctx context.Context, id string) (models.Status, error)
	List(ctx context.Context, limit, offset int) ([]models.Status, error)
	Update(ctx context.Context, status models.Status) error
	Delete(ctx context.Context, id string) error
}

// StatusLogRepository defines the interface for status log-related database operations
type StatusLogRepository interface {
	Create(ctx context.Context, log models.StatusLog) (models.StatusLog, error)
	GetByID(ctx context.Context, id string) (models.StatusLog, error)
	List(ctx context.Context, limit, offset int) ([]models.StatusLog, error)
	Update(ctx context.Context, log models.StatusLog) error
	Delete(ctx context.Context, id string) error
}

// RankRepository defines the interface for rank-related database operations
type RankRepository interface {
	Create(ctx context.Context, rank models.Rank) (models.Rank, error)
	GetByID(ctx context.Context, id string) (models.Rank, error)
	List(ctx context.Context, limit, offset int) ([]models.Rank, error)
	Update(ctx context.Context, rank models.Rank) error
	Delete(ctx context.Context, id string) error
}

// NewsArticleRepository defines the interface for news article-related database operations
type NewsArticleRepository interface {
	Create(ctx context.Context, news models.NewsArticle) (models.NewsArticle, error)
	GetByID(ctx context.Context, id string) (models.NewsArticle, error)
	List(ctx context.Context, limit, offset int) ([]models.NewsArticle, error)
	Update(ctx context.Context, news models.NewsArticle) error
	Delete(ctx context.Context, id string) error
}

// MediaRepository defines the interface for media-related database operations
type MediaRepository interface {
	Create(ctx context.Context, media models.Media) (models.Media, error)
	GetByID(ctx context.Context, id string) (models.Media, error)
	List(ctx context.Context, limit, offset int) ([]models.Media, error)
	Update(ctx context.Context, media models.Media) error
	Delete(ctx context.Context, id string) error
}

// ProductDocumentationRepository defines the interface for product documentation-related database operations
type ProductDocumentationRepository interface {
	Create(ctx context.Context, doc models.ProductDocumentation) (models.ProductDocumentation, error)
	GetByID(ctx context.Context, id string) (models.ProductDocumentation, error)
	GetByForeignKey(ctx context.Context, foreignKey string) ([]models.ProductDocumentation, error)
	List(ctx context.Context, limit, offset int) ([]models.ProductDocumentation, error)
	Update(ctx context.Context, doc models.ProductDocumentation) error
	Delete(ctx context.Context, id string) error
}

// LogRepository defines the interface for log-related database operations
type LogRepository interface {
	Create(ctx context.Context, log models.Log) (models.Log, error)
	GetByID(ctx context.Context, id string) (models.Log, error)
	List(ctx context.Context, limit, offset int) ([]models.Log, error)
	Delete(ctx context.Context, id string) error
}
