package handlers

import (
	"file-upload/internal/database"
	"file-upload/internal/services"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UploadHandler(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "file required"})
		return
	}

	src, err := file.Open()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not read file"})
		return
	}
	defer src.Close()

	url, err := services.UploadFile(file.Filename, src)
	if err != nil {
		log.Println("Failed to upload file:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "upload failed"})
		return
	}

	db, err := database.Connect()
	if err != nil {
		log.Println("Failed to connect to DB:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "DB connection failed"})
		return
	}
	defer db.Close()

	userID := c.GetInt("user_id")
	_, err = db.Exec(`
        INSERT INTO files (user_id, filename, url, size)
        VALUES ($1, $2, $3, $4)
    `, userID, file.Filename, url, file.Size)
	if err != nil {
		log.Println("Failed to insert file metadata:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to save file metadata"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"url": url})
}
