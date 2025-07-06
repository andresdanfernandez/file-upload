package services

import (
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"file-upload/internal/database"
	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	UserID int    `json:"user_id"`
	Email  string `json:"email"`
	jwt.RegisteredClaims
}

// SupabaseTokenClaims represents the structure of Supabase JWT tokens
type SupabaseTokenClaims struct {
	Aud   string `json:"aud"`
	Exp   int64  `json:"exp"`
	Iat   int64  `json:"iat"`
	Iss   string `json:"iss"`
	Sub   string `json:"sub"`
	Email string `json:"email"`
	Role  string `json:"role"`
	jwt.RegisteredClaims
}

// ValidateSupabaseToken validates a Supabase JWT token
func ValidateSupabaseToken(tokenString string) (*Claims, error) {
	log.Printf("Validating Supabase token...")
	
	// Get Supabase JWT secret from environment
	jwtSecret := os.Getenv("SUPABASE_JWT_SECRET")
	if jwtSecret == "" {
		log.Printf("SUPABASE_JWT_SECRET not found, using unverified parsing for development")
		// For development, use unverified parsing
		return validateTokenUnverified(tokenString)
	}
	
	// Parse and validate the token with the secret
	token, err := jwt.ParseWithClaims(tokenString, &SupabaseTokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		// Validate the signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(jwtSecret), nil
	})
	
	if err != nil {
		log.Printf("Token validation failed: %v", err)
		return nil, fmt.Errorf("token validation failed: %v", err)
	}
	
	if claims, ok := token.Claims.(*SupabaseTokenClaims); ok && token.Valid {
		log.Printf("Token validated successfully. Email: %s, Sub: %s", claims.Email, claims.Sub)
		
		// Get or create user in local database
		userID, err := getOrCreateUser(claims.Email, claims.Sub)
		if err != nil {
			log.Printf("Failed to get/create user: %v", err)
			return nil, fmt.Errorf("failed to get/create user: %v", err)
		}
		
		log.Printf("User ID resolved: %d", userID)
	
		return &Claims{
			UserID: userID,
			Email:  claims.Email,
			RegisteredClaims: jwt.RegisteredClaims{
				ExpiresAt: jwt.NewNumericDate(time.Unix(claims.Exp, 0)),
				IssuedAt:  jwt.NewNumericDate(time.Unix(claims.Iat, 0)),
			},
		}, nil
	}

	log.Printf("Failed to extract claims from token")
	return nil, errors.New("invalid Supabase token")
}

// validateTokenUnverified parses token without signature validation (for development)
func validateTokenUnverified(tokenString string) (*Claims, error) {
	parsedToken, _, err := new(jwt.Parser).ParseUnverified(tokenString, &SupabaseTokenClaims{})
	if err != nil {
		log.Printf("Failed to parse token: %v", err)
		return nil, fmt.Errorf("failed to parse token: %v", err)
	}

	if claims, ok := parsedToken.Claims.(*SupabaseTokenClaims); ok {
		log.Printf("Token parsed successfully (unverified). Email: %s, Sub: %s", claims.Email, claims.Sub)
		
		// Check if token is expired
		if claims.Exp > 0 && time.Now().Unix() > claims.Exp {
			log.Printf("Token is expired")
			return nil, errors.New("token is expired")
		}
		
		// Get or create user in local database
		userID, err := getOrCreateUser(claims.Email, claims.Sub)
		if err != nil {
			log.Printf("Failed to get/create user: %v", err)
			return nil, fmt.Errorf("failed to get/create user: %v", err)
		}
		
		log.Printf("User ID resolved: %d", userID)
	
		return &Claims{
			UserID: userID,
			Email:  claims.Email,
			RegisteredClaims: jwt.RegisteredClaims{
				ExpiresAt: jwt.NewNumericDate(time.Unix(claims.Exp, 0)),
				IssuedAt:  jwt.NewNumericDate(time.Unix(claims.Iat, 0)),
			},
		}, nil
	}

	log.Printf("Failed to extract claims from token")
	return nil, errors.New("invalid Supabase token")
}

// getOrCreateUser gets or creates a user in the local database
func getOrCreateUser(email, supabaseID string) (int, error) {
	db, err := database.Connect()
	if err != nil {
		return 0, fmt.Errorf("failed to connect to database: %v", err)
	}
	defer db.Close()

	// First, try to find existing user by email
	var userID int
	err = db.QueryRow("SELECT id FROM users WHERE email = $1", email).Scan(&userID)
	if err == nil {
		// User exists, return the ID
		return userID, nil
	}

	// User doesn't exist, create a new one
	// Note: We're not storing the password since Supabase handles auth
	err = db.QueryRow(`
		INSERT INTO users (email, password, supabase_id)
		VALUES ($1, $2, $3)
		RETURNING id
	`, email, "supabase_auth", supabaseID).Scan(&userID)
	if err != nil {
		return 0, fmt.Errorf("failed to create user: %v", err)
	}

	return userID, nil
} 