package models

import (
	"gorm.io/gorm"
)

// Note struct
type Note struct {
	gorm.Model
	Title       *string	`json:"title" gorm:"not null"`
	Text 	  	*string	`json:"text" gorm:"not null"`
}
