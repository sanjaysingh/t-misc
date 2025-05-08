package apperrors

import "net/http"

// ErrorNotice is the structure for the "notice" part of the JSON error response.
type ErrorNotice struct {
	ID          string `json:"id"`
	Description string `json:"description"`
}

// JSONErrorResponse is the full structure for sending JSON error responses.
type JSONErrorResponse struct {
	Notice ErrorNotice `json:"notice"`
}

// APIError represents a custom application error with an ID, description, and HTTP status code.
type APIError struct {
	ID          string
	Description string
	StatusCode  int
}

// Error makes APIError satisfy the error interface.
func (e *APIError) Error() string {
	return e.Description // Standard error interface returns the description
}

// New creates a new APIError.
func New(statusCode int, id, description string) *APIError {
	return &APIError{
		StatusCode:  statusCode,
		ID:          id,
		Description: description,
	}
}

// Predefined application errors
var (
	ErrInvalidRequestBody = New(http.StatusBadRequest, "ErrInvalidRequestBody", "The request body is malformed or contains invalid data.")
	ErrInternalServer     = New(http.StatusInternalServerError, "ErrInternalServer", "An unexpected internal server error occurred.")
	// Example specific error
	ErrPaymentNotFound = New(http.StatusNotFound, "ErrPaymentNotFound", "The requested payment was not found.")
	ErrValidation      = New(http.StatusBadRequest, "ErrValidation", "Input validation failed.") // Generic validation
)

// Specific validation errors (can be created dynamically or predefined if common)
func NewValidationError(id, description string) *APIError {
	return New(http.StatusBadRequest, id, description)
}
