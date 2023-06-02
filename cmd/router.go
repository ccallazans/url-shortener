package main

import (
	"myapi/internal/app/application/usecase"
	"myapi/internal/app/interfaces/handlers"
	"myapi/internal/app/interfaces/repository"

	"gorm.io/gorm"

	"github.com/ccallazans/url-shortener/api/v1/handlers"
	"github.com/ccallazans/url-shortener/api/v1/middleware"
	"github.com/gin-gonic/gin"
)

// type Router struct {
// 	router *gin.Engine
// }

func Config(db *gorm.DB) *gin.Engine {

	// Repositories
	userRepository := repository.NewUserRepository(db)
	// shortenerRepository := repository.NewShortenerRepository(db)

	// Services
	userUsecase := usecase.NewUserUsecase(userRepository)

	// Handlers
	urlHandler := handlers.NewUserHandler(userUsecase)

	router := gin.Default()

	router.GET("/:hash", urlHandler.RedirectUrl)

	v1Router := router.Group("/v1")
	{

		urlRouter := v1Router.Group("/url", middleware.AuthMiddleware())
		{
			urlRouter.POST("/", middleware.ValidateUrlMiddleware(), urlHandler.CreateUrl)
		}

		userRouter := v1Router.Group("/user", middleware.AuthMiddleware(), middleware.RoleMiddleware(roles.ADMIN))
		{
			userRouter.GET("/", userHandler.GetAllUsers)
			userRouter.GET("/:id", userHandler.GetUser)
		}

		authRouter := v1Router.Group("/auth", middleware.ValidateUserRequestMiddleware())
		{
			authRouter.POST("/login", authHandler.AuthUser)
			authRouter.POST("/register", userHandler.CreateUser)
		}
	}

	return router
}
