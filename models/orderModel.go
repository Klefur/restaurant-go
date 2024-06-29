package models

import (
	"time"

	"gorm.io/gorm"
)

// Order struct
type Order struct {
	gorm.Model
	Order_Date 	*time.Time		`json:"order_date" gorm:"not null"`
	Table_id	*uint			`gorm:"not null"`
	Invoices 	*[]Invoice		`json:"invoices"`
	OrderItems 	*[]OrderItem	`json:"order_items"`
}