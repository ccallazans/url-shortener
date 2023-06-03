package handlers

import (
	"myapi/internal/app/adapter/middleware"
	"myapi/internal/app/adapter/repository"
	"myapi/internal/app/application/usecase"

	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
)

func RouterConfig(db *gorm.DB) *gin.Engine {

	// Repositories
	userRepository := repository.NewUserRepository(db)
	shortenerRepository := repository.NewShortenerRepository(db)

	// Services
	userUsecase := usecase.NewUserUsecase(userRepository)
	shortenerUsecase := usecase.NewShortenerUsecase(shortenerRepository)

	// Handlers
	userHandler := NewUserHandler(userUsecase)
	shortenerHandler := NewShortenerHandler(shortenerUsecase)

	router := gin.Default()
	router.GET("/:hash", shortenerHandler.Redirect)
	v1Router := router.Group("/v1")
	{
		shortenerRouter := v1Router.Group("/shortener", middleware.AuthenticationMiddleware())
		shortenerRouter.POST("/", shortenerHandler.CreateShortener)

		userRouter := v1Router.Group("/user", middleware.AuthenticationMiddleware())
		userRouter.GET("/", userHandler.GetAllUsers)
		userRouter.GET("/:uuid", userHandler.GetUser)

		authRouter := v1Router.Group("/auth")
		authRouter.POST("/login", userHandler.AuthUser)
		authRouter.POST("/register", userHandler.CreateUser)
	}

	return router
}
