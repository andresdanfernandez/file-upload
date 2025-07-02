package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"file-upload/internal/models"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestRegisterHandler(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.New()
	router.POST("/register", RegisterHandler)

	tests := []struct {
		name       string
		payload    models.RegisterRequest
		wantStatus int
	}{
		{
			name: "valid registration",
			payload: models.RegisterRequest{
				Email:    "test@example.com",
				Password: "password123",
			},
			wantStatus: http.StatusCreated,
		},
		{
			name: "invalid email",
			payload: models.RegisterRequest{
				Email:    "invalid-email",
				Password: "password123",
			},
			wantStatus: http.StatusBadRequest,
		},
		{
			name: "short password",
			payload: models.RegisterRequest{
				Email:    "test@example.com",
				Password: "123",
			},
			wantStatus: http.StatusBadRequest,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			jsonData, _ := json.Marshal(tt.payload)
			req, _ := http.NewRequest("POST", "/register", bytes.NewBuffer(jsonData))
			req.Header.Set("Content-Type", "application/json")

			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)

			assert.Equal(t, tt.wantStatus, w.Code)
		})
	}
}

func TestLoginHandler(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.New()
	router.POST("/login", LoginHandler)

	tests := []struct {
		name       string
		payload    models.LoginRequest
		wantStatus int
	}{
		{
			name: "valid login",
			payload: models.LoginRequest{
				Email:    "test@example.com",
				Password: "password123",
			},
			wantStatus: http.StatusOK,
		},
		{
			name: "missing email",
			payload: models.LoginRequest{
				Password: "password123",
			},
			wantStatus: http.StatusBadRequest,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			jsonData, _ := json.Marshal(tt.payload)
			req, _ := http.NewRequest("POST", "/login", bytes.NewBuffer(jsonData))
			req.Header.Set("Content-Type", "application/json")

			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)

			assert.Equal(t, tt.wantStatus, w.Code)
		})
	}
} 