package api

import (
	"log"
	"net/http"

	"github.com/axz/control-plane/internal/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// ServiceHandler handles service-related API endpoints
type ServiceHandler struct {
	db *gorm.DB
}

// NewServiceHandler creates a new service handler
func NewServiceHandler(db *gorm.DB) *ServiceHandler {
	return &ServiceHandler{db: db}
}

// RegisterService creates a new service
func (h *ServiceHandler) RegisterService(c *gin.Context) {
	var req models.CreateServiceRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	log.Printf("Received service registration request: %+v", req)

	service := models.Service{
		Name:        req.Name,
		Description: req.Description,
		URL:         req.URL,
	}

	if err := h.db.Create(&service).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create service"})
		return
	}

	c.JSON(http.StatusCreated, service)
} 