package handlers

import (
	"file-upload/internal/database"
	"file-upload/internal/services"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func sanitizeFilename(name string) string {
	// Remove path separators and dangerous characters
	name = strings.ReplaceAll(name, "..", "")
	name = strings.ReplaceAll(name, "/", "_")
	name = strings.ReplaceAll(name, "\\", "_")
	name = strings.ReplaceAll(name, "\x00", "_")
	return name
}

func allowedFileType(filename string) bool {
	allowed := []string{".jpg", ".jpeg", ".png", ".gif", ".pdf", ".txt", ".zip"}
	ext := strings.ToLower(filename[strings.LastIndex(filename, ".") : ])
	for _, a := range allowed {
		if ext == a {
			return true
		}
	}
	return false
}

func UploadHandler(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "file required"})
		return
	}

	// Enforce max file size (10MB)
	if file.Size > 10*1024*1024 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "file too large (max 10MB)"})
		return
	}

	// Enforce allowed file types
	if !allowedFileType(file.Filename) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "file type not allowed"})
		return
	}

	safeFilename := sanitizeFilename(file.Filename)

	src, err := file.Open()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not read file"})
		return
	}
	defer src.Close()

	url, err := services.UploadFile(safeFilename, src)
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
    `, userID, safeFilename, url, file.Size)
	if err != nil {
		log.Println("Failed to insert file metadata:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to save file metadata"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"url": url})
}

// DeleteFileHandler deletes a file by id if it belongs to the user
func DeleteFileHandler(c *gin.Context) {
	fileID := c.Param("id")
	userID := c.GetInt("user_id")

	db, err := database.Connect()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "DB connection failed"})
		return
	}
	defer db.Close()

	// Get file info and check ownership
	var s3url string
	var dbUserID int
	err = db.QueryRow(`SELECT url, user_id FROM files WHERE id = $1`, fileID).Scan(&s3url, &dbUserID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "File not found"})
		return
	}
	if dbUserID != userID {
		c.JSON(http.StatusForbidden, gin.H{"error": "Not authorized"})
		return
	}

	// Extract S3 key from URL
	parts := strings.Split(s3url, "/")
	key := parts[len(parts)-1]

	// Delete from S3
	err = services.DeleteFileFromS3(key)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete from S3"})
		return
	}

	// Delete from DB
	_, err = db.Exec(`DELETE FROM files WHERE id = $1`, fileID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete from DB"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true})
}
