package main

import (
	"file-upload/config"
	"file-upload/internal/handlers"
	"file-upload/internal/middleware"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadEnv()

	router := gin.Default()

	// CORS configuration for production
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{
			"http://localhost:5173", 
			"http://localhost:3000", 
			"http://localhost:4173",
			"https://amused-creation-production.up.railway.app",
			"https://file-upload-production-9bd0.up.railway.app",
		},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "healthy"})
	})

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
