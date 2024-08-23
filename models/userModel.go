package models

import (
	"gorm.io/gorm"
)

// User struct
type User struct {
	gorm.Model
	First_name 		string		`json:"first_name" gorm:"not null"`
	Last_name 		string		`json:"last_name" gorm:"not null"`
	Email 			string		`json:"email" gorm:"not null;index"`
	Password 		string		`json:"password" gorm:"not null"`
	Avatar 			string		`json:"avatar"`
	Phone 			string		`json:"phone" gorm:"not null"`
}