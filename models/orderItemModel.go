package models

import (
	"gorm.io/gorm"
)

// OrderItem struct
type OrderItem struct {
	gorm.Model
	Quantity 	*int		`json:"quantity" gorm:"check:quantity > 0; not null"`
	Unit_price 	*float64	`json:"unit_price" gorm:"check:unit_price > 0; not null"`
	Order_id	*uint		`json:"order_id" gorm:"not null"`
	Food_id	 	*uint		`json:"food_id" gorm:"not null"`
}