package services

import (
	"log"

	"apm/internal/db"
	"apm/internal/db/repository"
)

// Services holds all service instances
type Services struct {
	UserService                 UserService
	UserGroupService            UserGroupService
	StakeholderService          StakeholderService
	EntityService               EntityService
	SoftwareService             SoftwareService
	FunctionalCategoryService   FunctionalCategoryService
	SoftwareGroupService        SoftwareGroupService
	StatusService               StatusService
	StatusLogService            StatusLogService
	RankService                 RankService
	NewsArticleService          NewsArticleService
	MediaService                MediaService
	ProductDocumentationService ProductDocumentationService
	LogService                  LogService
}

// NewServices creates a new services manager
func NewServices(db *db.Database, logger *log.Logger) *Services {
	// Instantiate repositories needed by services
	softwareRepo := repository.NewPostgresSoftwareRepository(db.Pool)
	// userRepo := repository.NewPostgresUserRepository(db.Pool)
	// ... instantiate other repos ...

	// TODO: Uncomment and implement other service initializations as needed
	return &Services{
		// UserService: NewUserService(userRepo, logger),
		// UserGroupService: NewUserGroupService(userRepo, logger),
		// StakeholderService: NewStakeholderService(stakeholderRepo, logger),
		// EntityService: NewEntityService(entityRepo, logger),

		// Initialize software service with the repository instance
		SoftwareService: NewSoftwareService(softwareRepo, logger),

		// FunctionalCategoryService: NewFunctionalCategoryService(functionalCategoryRepo, logger),
		// SoftwareGroupService: NewSoftwareGroupService(softwareGroupRepo, logger),
		// StatusService: NewStatusService(statusRepo, logger),
		// StatusLogService: NewStatusLogService(statusLogRepo, logger),
		// RankService: NewRankService(rankRepo, logger),
		// NewsArticleService: NewNewsArticleService(newsArticleRepo, logger),
		// MediaService: NewMediaService(mediaRepo, logger),
		// ProductDocumentationService: NewProductDocumentationService(productDocumentationRepo, logger),
		// LogService: NewLogService(logRepo, logger),
	}
}
