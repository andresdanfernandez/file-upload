package main

import (
	"file-upload-api/config"
	"file-upload-api/internal/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadEnv()

	router := gin.Default()

	router.POST("/upload", handlers.UploadHandler)
	router.GET("/files", handlers.ListFilesHandler)
	router.GET("/download/:key", handlers.DownloadHandler)

	router.Run(":" + config.Env("PORT", "8080"))
}
