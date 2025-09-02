package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"go-health-app/internal/service"
)

func TestHealthHandler_Check(t *testing.T) {
	// Setup Gin in test mode
	gin.SetMode(gin.TestMode)

	// Wire handler with service
	hs := service.NewHealthService()
	h := NewHealthHandler(hs)

	// Setup router
	router := gin.Default()
	router.GET("/health", h.Check)

	// Perform request
	req, _ := http.NewRequest("GET", "/health", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Assertions
	if w.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", w.Code)
	}

	expectedBody := `{"status":"ok"}`
	if w.Body.String() != expectedBody {
		t.Errorf("Expected body %s, got %s", expectedBody, w.Body.String())
	}
}
