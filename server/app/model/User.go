package model

import "gorm.io/gorm"

// User represents a user in the system.
// swagger:model
type User struct {
	gorm.Model
	// Nickname of the user
	// example: JohnDoe
	// required: false
	Nickname string `gorm:"size:255" json:"nickname"`

	// Username of the user
	// example: xxxxx
	// required: true
	Username string `gorm:"unique,size:20;" json:"username"`

	// Password of the user
	// required: true
	Password string `json:"password" gorm:"size:200;not null"`

	// Email of the user
	// format: email
	// required: false
	Email string `gorm:"size:255" json:"email"`

	// Phone number of the user
	// format: int64
	// required: false
	Phone string `gorm:"size:255;type:varchar(100);" json:"phone"`

	// Avatar URL of the user
	// example: http://xx
	// required: false
	Avatar string `gorm:"size: 255;not null" json:"avatar"`
}
