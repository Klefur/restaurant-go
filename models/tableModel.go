package models

import (
	"gorm.io/gorm"
)

// Table struct
type Table struct {
	gorm.Model
	Number_of_guests 	int		`json:"number_of_guests" gorm:"check number_of_guests > 0, not null"`
	Table_number 		int		`json:"table_number" gorm:"check table_number > 0, not null"`
	Orders 				[]Order	`json:"orders"`
}