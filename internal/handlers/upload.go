package handlers

import (
	"database/sql"
	"file-upload/internal/services"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq" // Needed to initialize the postgres driver
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
		c.JSON(http.StatusInternalServerError, gin.H{"error": "upload failed"})
		return
	}

	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Println("Failed to connect to DB", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "DB connection failed"})
		return
	}
	defer db.Close()

	_, err = db.Exec(`
        INSERT INTO files (filename, url, size)
        VALUES ($1, $2, $3)
    `, file.Filename, url, file.Size)
	if err != nil {
		log.Println("Failed to insert file metadata:", err)
	}

	c.JSON(http.StatusOK, gin.H{"url": url})
}
