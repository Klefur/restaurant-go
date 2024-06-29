package models

import (
	"gorm.io/gorm"
)

// Food struct
type Food struct {
	gorm.Model
	Name        *string			`json:"name" gorm:"not null"`	 	
	Price       *int			`json:"price" gorm:"not null"`
	Food_image 	*string			`json:"food_image" gorm:"not null"`
	Menu_id    	*uint   		`json:"menu_id" gorm:"not null"`
	OrderItems 	*[]OrderItem	`json:"order_items"`
}