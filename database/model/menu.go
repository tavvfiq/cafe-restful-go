package model

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

// Menu table
type Menu struct {
	ID           uuid.UUID
	Name         string
	Image        sql.NullString
	Price        uint
	Quantity     uint
	CategoryID   uint8
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    time.Time      `gorm:"index" json:"deleted_at"`
	OrderHistory []OrderHistory `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
