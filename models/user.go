package models

import (
	"time"
)

// User represents a user in the system
type User struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Email     string    `json:"email" gorm:"uniqueIndex;not null"`
	Password  string    `json:"-" gorm:"not null"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// TableName specifies the table name for the User model
func (User) TableName() string {
	return "users"
}

// SafeResponse returns a user without sensitive fields
func (u User) SafeResponse() map[string]interface{} {
	return map[string]interface{}{
		"id":         u.ID,
		"email":      u.Email,
		"first_name": u.FirstName,
		"last_name":  u.LastName,
		"created_at": u.CreatedAt,
		"updated_at": u.UpdatedAt,
	}
} 