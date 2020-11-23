package model

import (
	"time"

	"github.com/google/uuid"
)

// User model
type User struct {
	ID           uuid.UUID      `json:"id"`
	FirstName    string         `json:"first_name"`
	LastName     string         `json:"last_name"`
	ProfileImage string         `json:"profile_image"`
	PhoneNumber  string         `json:"phone_number"`
	Email        string         `gorm:"unique" json:"email"`
	Password     string         `json:"password"`
	LevelID      uint8          `json:"level_id"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	History      []History      `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	OrderHistory []OrderHistory `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
