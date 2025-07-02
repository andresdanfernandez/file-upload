package handlers

import (
	"file-upload/internal/database"
	"file-upload/internal/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ListFilesHandler(c *gin.Context) {
	db, err := database.Connect()
	if err != nil {
		log.Println("DB connection failed:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "DB connection failed"})
		return
	}
	defer db.Close()

	userID := c.GetInt("user_id")
	rows, err := db.Query(`
		SELECT id, user_id, filename, url, size, uploaded_at
		FROM files
		WHERE user_id = $1
		ORDER BY uploaded_at DESC
	`, userID)
	if err != nil {
		log.Println("Failed to query files:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "query failed"})
		return
	}
	defer rows.Close()

	var files []models.File
	for rows.Next() {
		var f models.File
		if err := rows.Scan(&f.ID, &f.UserID, &f.Filename, &f.URL, &f.Size, &f.UploadedAt); err != nil {
			log.Println("Failed to scan row:", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to read result"})
			return
		}
		files = append(files, f)
	}

	c.JSON(http.StatusOK, files)
}
