package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/syaafiudinm/go-starter-kit/internal/handler"
	"github.com/syaafiudinm/go-starter-kit/internal/middleware"
	"github.com/syaafiudinm/go-starter-kit/internal/repository"
	"github.com/syaafiudinm/go-starter-kit/internal/service"
	"gorm.io/gorm"
)

func Setup(router *gin.Engine, db *gorm.DB) {
	// Global middleware
	router.Use(middleware.CORS())
	router.Use(middleware.Recovery())

	// Initialize dependencies
	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	// API v1 group
	v1 := router.Group("/api/v1")
	{
		// Health check
		v1.GET("/health", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"status":  "ok",
				"message": "Service is running",
			})
		})

		// User routes
		users := v1.Group("/users")
		{
			users.POST("", userHandler.Create)
			users.GET("", userHandler.GetAll)
			users.GET("/:id", userHandler.GetByID)
			users.PUT("/:id", userHandler.Update)
			users.DELETE("/:id", userHandler.Delete)
		}
	}
}
