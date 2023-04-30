package router

import (
	"github.com/ccallazans/url-shortener/api/v1/handlers"
	"github.com/ccallazans/url-shortener/api/v1/middleware"
	"github.com/ccallazans/url-shortener/internal/database/sqliteImpl"
	"github.com/ccallazans/url-shortener/internal/domain/models/roles"
	service "github.com/ccallazans/url-shortener/internal/domain/service/impl"
	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
)

// type Router struct {
// 	router *gin.Engine
// }

func Config(db *gorm.DB) *gin.Engine {

	// Repositories
	urlRepository := sqliteImpl.NewSqliteUrlRepository(db)
	userRepository := sqliteImpl.NewSqliteUserRepository(db)

	// Services
	urlService := service.NewUrlService(urlRepository)
	userService := service.NewUserService(userRepository)

	// Handlers
	urlHandler := handlers.NewUrlHandler(urlService)
	userHandler := handlers.NewUserHandler(userService)
	authHandler := handlers.NewAuthHandler(userService)

	router := gin.Default()

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
