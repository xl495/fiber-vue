package model

import (
	"gorm.io/gorm"
	"time"
)

// User represents a user in the system.
// swagger:model
type User struct {
	gorm.Model
	// Nickname of the user
	// example: JohnDoe
	// required: true
	Nickname string `gorm:"unique" json:"nickname"`

	// Password of the user
	// required: true
	Password string `json:"password"`

	// Email of the user
	// format: email
	// required: false
	Email string `gorm:"unique" json:"email"`

	// Phone number of the user
	// format: int64
	// required: false
	Phone string `gorm:"unique" json:"phone"`

	// Avatar URL of the user
	Avatar string `json:"avatar"`
}

// GormModel gorm.Model represents the base model for gorm models.
// swagger:model
type GormModel struct {
	ID        uint `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
