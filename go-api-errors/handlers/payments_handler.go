package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"go-api-errors/apperrors"
	"go-api-errors/httputils"
	"go-api-errors/models"
)

// paymentsStore holds our in-memory list of payments.
// In a real application, this would be a database.
var paymentsStore = []models.Payment{
	{PaymentDate: time.Date(2024, time.July, 1, 10, 0, 0, 0, time.UTC), PaymentReference: "REF001", Amount: 100.50},
	{PaymentDate: time.Date(2024, time.July, 15, 11, 30, 0, 0, time.UTC), PaymentReference: "REF002", Amount: 75.25},
}

// GetPayments returns a list of payments from the in-memory store.
func GetPayments(w http.ResponseWriter, r *http.Request) error {
	httputils.RespondWithJSON(w, http.StatusOK, paymentsStore)
	return nil
}

// AddPayment adds a new payment to the in-memory store.
func AddPayment(w http.ResponseWriter, r *http.Request) error {
	var newPayment models.Payment
	if err := json.NewDecoder(r.Body).Decode(&newPayment); err != nil {
		// Example of a specific validation error for payment date (if it were part of the request)
		// if newPayment.PaymentDate.IsZero() { // Assuming PaymentDate might be validated
		// 	 return apperrors.NewValidationError("PaymentDateValidationError", "Payment date is required and cannot be empty.")
		// }
		return apperrors.ErrInvalidRequestBody // Generic invalid body error
	}

	// Basic validation example (can be expanded)
	if newPayment.PaymentReference == "" {
		return apperrors.NewValidationError("PaymentReferenceMissing", "Payment reference is required.")
	}
	if newPayment.Amount <= 0 {
		return apperrors.NewValidationError("InvalidPaymentAmount", "Payment amount must be greater than zero.")
	}

	newPayment.PaymentDate = time.Now().UTC() // Set payment date to current time

	paymentsStore = append(paymentsStore, newPayment)

	httputils.RespondWithJSON(w, http.StatusCreated, newPayment)
	return nil
}
