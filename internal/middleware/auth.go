package middleware

import (
	"file-upload/internal/services"
	"log"
	"net/http"
	"strings"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Allow OPTIONS requests to pass through for CORS preflight
		if c.Request.Method == "OPTIONS" {
			log.Printf("Auth middleware: Allowing OPTIONS request to pass through")
			c.Next()
			return
		}
		
		authHeader := c.GetHeader("Authorization")
		log.Printf("Auth middleware: Authorization header present: %t", authHeader != "")
		
		if authHeader == "" {
			log.Printf("Auth middleware: No authorization header found")
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header required"})
			c.Abort()
			return
		}
		
		// Extract token from "Bearer <token>"
		tokenParts := strings.Split(authHeader, " ")
		if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
			log.Printf("Auth middleware: Invalid authorization header format. Parts: %d, First part: %s", len(tokenParts), tokenParts[0])
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid authorization header format"})
			c.Abort()
			return
		}
		
		token := tokenParts[1]
		log.Printf("Auth middleware: Token length: %d", len(token))
		
		claims, err := services.ValidateSupabaseToken(token)
		if err != nil {
			log.Printf("Auth middleware: Token validation failed: %v", err)
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}
		
		log.Printf("Auth middleware: Token validated successfully for user: %s (ID: %d)", claims.Email, claims.UserID)
		c.Set("user_id", claims.UserID)
		c.Set("email", claims.Email)
		c.Next()
	}
}