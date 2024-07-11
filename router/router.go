package router

import (
	"eira/document/delivery/middleware"
	restTranslation "eira/document/delivery/rest/translation"
	translationService "eira/document/services/translation"
	usecaseTranslation "eira/document/usecase/acs/translation"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB) *gin.Engine {

	AcsTranslationService := translationService.DefACSTranslationService()

	translationUseCase := &usecaseTranslation.TranslationUseCaseImpl{
		AiProviders: AcsTranslationService,
	}

	translationHandler := &restTranslation.TranslationHandler{
		TranslationUsecase: translationUseCase,
	}

	router := gin.Default()

	// Apply CORS middleware globally
	router.Use(middleware.CORSMiddleware())

	// Public routes
	router.POST("/translate-doc", translationHandler.TranslateFile)

	// Protected routes
	// authRoutes := router.Group("/")
	// authRoutes.Use(middleware.JWTAuthMiddleware())
	// {
	// 	authRoutes.POST("/users/:id/upload", userHandler.UploadUserProfileImage)
	// 	authRoutes.GET("/users/:id/predict", userHandler.GetPredictionForUser)
	// }

	// // Admin routes
	// adminRoutes := router.Group("/")
	// adminRoutes.Use(middleware.JWTAuthMiddleware(), middleware.AdminMiddleware())
	// {
	// 	// Add admin-specific routes here
	// }

	return router
}
