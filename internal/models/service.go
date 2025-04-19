package models

import (
	"time"

	"gorm.io/gorm"
)

// Service represents a backend service registered in the platform
type Service struct {
	ID          string         `json:"id" gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	Name        string         `json:"name" gorm:"not null"`
	Description string         `json:"description"`
	URL         string         `json:"url" gorm:"not null"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"-" gorm:"index"`
}

// CreateServiceRequest represents the request body for service creation
type CreateServiceRequest struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
	URL         string `json:"url" binding:"required,url"`
} 