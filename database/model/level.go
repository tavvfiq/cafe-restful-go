package model

import "time"

// Level user level
type Level struct {
	ID        uint8
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `gorm:"index" json:"deleted_at"`
	Name      string
	User      []User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
