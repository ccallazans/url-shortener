package router

import (
	"database/sql"

	"github.com/ccallazans/url-shortener/api/v1/handlers"
	"github.com/ccallazans/url-shortener/internal/database/sqliteImpl"
	service "github.com/ccallazans/url-shortener/internal/domain/service/impl"

	"github.com/gin-gonic/gin"
)

type Router struct {
	router *gin.Engine
}

func Config(db *sql.DB) *gin.Engine {

	// Repositories
	userRepository := sqliteImpl.NewSqliteUserRepository(db)

	// Services
	userService := service.NewUserService(userRepository)

	// Handlers
	userHandler := handlers.NewUserHandler(userService)

	
	router := gin.Default()


	v1Router := router.Group("/v1")
	{
		userRouter := v1Router.Group("/user")
		{
			userRouter.GET("/:id", userHandler.GetUser)
			userRouter.POST("/", userHandler.CreateUser)
		}
	}


	return router
}
