package model

import "time"

// Category table
type Category struct {
	ID        uint8
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `gorm:"index" json:"deleted_at"`
	Name      string
	Menu      []Menu `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
