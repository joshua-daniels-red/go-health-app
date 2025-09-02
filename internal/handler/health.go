package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go-health-app/internal/service"
)

// HealthHandler wraps the HealthService
type HealthHandler struct {
	service *service.HealthService
}

// NewHealthHandler initializes the handler with a service
func NewHealthHandler(s *service.HealthService) *HealthHandler {
	return &HealthHandler{service: s}
}

// Check handles GET /health
func (h *HealthHandler) Check(c *gin.Context) {
	status := h.service.Check()
	c.JSON(http.StatusOK, status)
}
