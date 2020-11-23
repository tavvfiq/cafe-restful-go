package model

import (
	"time"

	"github.com/google/uuid"
)

// History table
type History struct {
	ID           uuid.UUID
	Cashier      string
	UserID       uint
	Amount       uint
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	OrderHistory []OrderHistory `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
