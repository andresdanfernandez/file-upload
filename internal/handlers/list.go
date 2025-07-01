package handlers

import (
	"database/sql"
	"file-upload/internal/models"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func ListFilesHandler(c *gin.Context) {
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "DB connection failed"})
		return
	}
	defer db.Close()

	rows, err := db.Query("SELECT id, filename, url, size, uploaded_at FROM files ORDER BY uploaded_at DESC")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "query failed"})
		return
	}
	defer rows.Close()

	var files []models.File
	for rows.Next() {
		var f models.File
		if err := rows.Scan(&f.ID, &f.Filename, &f.URL, &f.Size, &f.UploadedAt); err == nil {
			files = append(files, f)
		}
	}

	c.JSON(http.StatusOK, files)
}
