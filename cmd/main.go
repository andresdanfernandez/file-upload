package main

import (
	"file-upload/config"
	"file-upload/internal/handlers"
	"file-upload/internal/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadEnv()

	router := gin.Default()

	// Health check
	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "healthy"})
	})

	// Public routes
	router.POST("/register", handlers.RegisterHandler)
	router.POST("/login", handlers.LoginHandler)

	// Protected routes
	protected := router.Group("/")
	protected.Use(middleware.AuthMiddleware())
	{
		protected.POST("/upload", handlers.UploadHandler)
		protected.GET("/files", handlers.ListFilesHandler)
		protected.GET("/download/:key", handlers.DownloadHandler)
	}

	router.Run(":" + config.Env("PORT", "8080"))
}
