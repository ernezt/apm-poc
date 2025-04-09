package handlers

import (
	"apm/internal/services"

	"github.com/gin-gonic/gin"
)

// Factory handles the creation and registration of all API handlers
type Factory struct {
	// Services
	userService                 services.UserService
	userGroupService            services.UserGroupService
	stakeholderService          services.StakeholderService
	entityService               services.EntityService
	softwareService             services.SoftwareService
	functionalCategoryService   services.FunctionalCategoryService
	softwareGroupService        services.SoftwareGroupService
	statusService               services.StatusService
	statusLogService            services.StatusLogService
	rankService                 services.RankService
	newsArticleService          services.NewsArticleService
	mediaService                services.MediaService
	productDocumentationService services.ProductDocumentationService
	logService                  services.LogService

	// Handlers
	userGroupHandler            *UserGroupHandler
	stakeholderHandler          *StakeholderHandler
	entityHandler               *EntityHandler
	softwareHandler             *SoftwareHandler
	functionalCategoryHandler   *FunctionalCategoryHandler
	softwareGroupHandler        *SoftwareGroupHandler
	statusHandler               *StatusHandler
	statusLogHandler            *StatusLogHandler
	rankHandler                 *RankHandler
	newsArticleHandler          *NewsArticleHandler
	mediaHandler                *MediaHandler
	productDocumentationHandler *ProductDocumentationHandler
	logHandler                  *LogHandler
}

// NewFactory creates a new handler factory
func NewFactory(
	userService services.UserService,
	userGroupService services.UserGroupService,
	stakeholderService services.StakeholderService,
	entityService services.EntityService,
	softwareService services.SoftwareService,
	functionalCategoryService services.FunctionalCategoryService,
	softwareGroupService services.SoftwareGroupService,
	statusService services.StatusService,
	statusLogService services.StatusLogService,
	rankService services.RankService,
	newsArticleService services.NewsArticleService,
	mediaService services.MediaService,
	productDocumentationService services.ProductDocumentationService,
	logService services.LogService,
) *Factory {
	f := &Factory{
		userService:                 userService,
		userGroupService:            userGroupService,
		stakeholderService:          stakeholderService,
		entityService:               entityService,
		softwareService:             softwareService,
		functionalCategoryService:   functionalCategoryService,
		softwareGroupService:        softwareGroupService,
		statusService:               statusService,
		statusLogService:            statusLogService,
		rankService:                 rankService,
		newsArticleService:          newsArticleService,
		mediaService:                mediaService,
		productDocumentationService: productDocumentationService,
		logService:                  logService,
	}

	f.initHandlers()
	return f
}

// initHandlers initializes all handler instances
func (f *Factory) initHandlers() {
	f.userGroupHandler = NewUserGroupHandler(f.userGroupService)
	f.stakeholderHandler = NewStakeholderHandler(f.stakeholderService)
	f.entityHandler = NewEntityHandler(f.entityService)
	f.softwareHandler = NewSoftwareHandler(f.softwareService)
	f.functionalCategoryHandler = NewFunctionalCategoryHandler(f.functionalCategoryService)
	f.softwareGroupHandler = NewSoftwareGroupHandler(f.softwareGroupService)
	f.statusHandler = NewStatusHandler(f.statusService)
	f.statusLogHandler = NewStatusLogHandler(f.statusLogService)
	f.rankHandler = NewRankHandler(f.rankService)
	f.newsArticleHandler = NewNewsArticleHandler(f.newsArticleService)
	f.mediaHandler = NewMediaHandler(f.mediaService)
	f.productDocumentationHandler = NewProductDocumentationHandler(f.productDocumentationService)
	f.logHandler = NewLogHandler(f.logService)
}

// RegisterRoutes registers all API routes
func (f *Factory) RegisterRoutes(router *gin.RouterGroup) {
	apiV1 := router.Group("/v1")

	// Register routes for each handler
	f.userGroupHandler.Register(apiV1)
	f.stakeholderHandler.Register(apiV1)
	f.entityHandler.Register(apiV1)
	f.softwareHandler.Register(apiV1)
	f.functionalCategoryHandler.Register(apiV1)
	f.softwareGroupHandler.Register(apiV1)
	f.statusHandler.Register(apiV1)
	f.statusLogHandler.Register(apiV1)
	f.rankHandler.Register(apiV1)
	f.newsArticleHandler.Register(apiV1)
	f.mediaHandler.Register(apiV1)
	f.productDocumentationHandler.Register(apiV1)
	f.logHandler.Register(apiV1)
}
