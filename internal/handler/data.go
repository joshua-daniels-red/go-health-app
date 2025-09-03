package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"go-health-app/internal/service"
)

type DataHandler struct {
	service *service.DataService
}

func NewDataHandler(s *service.DataService) *DataHandler {
	return &DataHandler{service: s}
}

func (h *DataHandler) GetData(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))

	data := h.service.GetPaginatedData(page, pageSize)
	c.JSON(http.StatusOK, data)
}
