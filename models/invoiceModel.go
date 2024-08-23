package models

import (
	"time"

	"gorm.io/gorm"
)

// Invoice struct
type Invoice struct {
	gorm.Model
	Payment_method   string    `json:"payment_method"`
	Payment_status   string    `json:"payment_status" gorm:"check:payment_status = 'PAID' or payment_status = 'PENDING'; not null"`
	Payment_due_date time.Time `json:"payment_due_date"`
	Order_id         uint      `json:"order_id" gorm:"not null"`
}
