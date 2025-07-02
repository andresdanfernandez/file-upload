package handlers

import (
	"database/sql"
	"file-upload/internal/database"
	"file-upload/internal/models"
	"file-upload/internal/services"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterHandler(c *gin.Context) {
	var req models.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db, err := database.Connect()
	if err != nil {
		log.Println("DB connection failed:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "DB connection failed"})
		return
	}
	defer db.Close()

	// Check if user already exists
	var existingUser models.User
	err = db.QueryRow("SELECT id FROM users WHERE email = $1", req.Email).Scan(&existingUser.ID)
	if err != sql.ErrNoRows {
		c.JSON(http.StatusConflict, gin.H{"error": "User already exists"})
		return
	}

	// Hash password
	hashedPassword, err := services.HashPassword(req.Password)
	if err != nil {
		log.Println("Password hashing failed:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Registration failed"})
		return
	}

	// Insert new user
	var userID int
	err = db.QueryRow(`
		INSERT INTO users (email, password)
		VALUES ($1, $2)
		RETURNING id
	`, req.Email, hashedPassword).Scan(&userID)
	if err != nil {
		log.Println("Failed to insert user:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Registration failed"})
		return
	}

	// Generate token
	token, err := services.GenerateToken(userID, req.Email)
	if err != nil {
		log.Println("Token generation failed:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Registration failed"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "User registered successfully",
		"token":   token,
		"user_id": userID,
	})
}

func LoginHandler(c *gin.Context) {
	var req models.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db, err := database.Connect()
	if err != nil {
		log.Println("DB connection failed:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "DB connection failed"})
		return
	}
	defer db.Close()

	// Get user
	var user models.User
	err = db.QueryRow(`
		SELECT id, email, password
		FROM users
		WHERE email = $1
	`, req.Email).Scan(&user.ID, &user.Email, &user.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	// Check password
	if !services.CheckPasswordHash(req.Password, user.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	// Generate token
	token, err := services.GenerateToken(user.ID, user.Email)
	if err != nil {
		log.Println("Token generation failed:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Login failed"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Login successful",
		"token":   token,
		"user_id": user.ID,
	})
} 