package models

import (
	"time"

	"gorm.io/gorm"
)

// Menu struct
type Menu struct {
	gorm.Model
	Name        *string		`json:"name" gorm:"not null"`
	Category 	*string		`json:"category" gorm:"not null"`
	Foods 		*[]Food		`json:"foods"`
	Start_date 	*time.Time	`json:"start_date"`
	End_date 	*time.Time	`json:"end_date"`
}
