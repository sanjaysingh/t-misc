package models

import "time"

// Payment represents a payment transaction.
type Payment struct {
	PaymentDate      time.Time `json:"paymentDate"`
	PaymentReference string    `json:"paymentReference"`
	Amount           float64   `json:"amount"`
}
