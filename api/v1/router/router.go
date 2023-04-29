package router

import (
	"database/sql"

	"github.com/ccallazans/url-shortener/api/v1/handlers"
	"github.com/ccallazans/url-shortener/api/v1/middleware"
	"github.com/ccallazans/url-shortener/internal/database/sqliteImpl"
	service "github.com/ccallazans/url-shortener/internal/domain/service/impl"

	"github.com/gin-gonic/gin"
)

// type Router struct {
// 	router *gin.Engine
// }

func Config(db *sql.DB) *gin.Engine {

	// Repositories
	userRepository := sqliteImpl.NewSqliteUserRepository(db)

	// Services
	userService := service.NewUserService(userRepository)

	// Handlers
	userHandler := handlers.NewUserHandler(userService)
	authHandler := handlers.NewAuthHandler(userService)

	router := gin.Default()

	v1Router := router.Group("/v1")
	{
		userRouter := v1Router.Group("/user", middleware.AuthMiddleware())
		{
			userRouter.GET("/", userHandler.GetAllUsers)
			userRouter.GET("/:id", userHandler.GetUser)
		}

		authRouter := v1Router.Group("/auth")
		{
			authRouter.POST("/login", middleware.ValidateUserRequestMiddleware(), authHandler.AuthUser)
			authRouter.POST("/register", middleware.ValidateUserRequestMiddleware(), userHandler.CreateUser)
		}
	}

	return router
}
