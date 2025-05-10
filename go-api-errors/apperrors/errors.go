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
type AppError struct {
	ID          string
	Description string
	StatusCode  int
}

// Error makes APIError satisfy the error interface.
func (e *AppError) Error() string {
	return e.Description // Standard error interface returns the description
}

// New creates a new APIError.
func New(statusCode int, id, description string) *AppError {
	return &AppError{
		StatusCode:  statusCode,
		ID:          id,
		Description: description,
	}
}

func NewBadRequestError(id, description string) *AppError {
	return New(http.StatusBadRequest, id, description)
}

func NewNotFoundError(id, description string) *AppError {
	return New(http.StatusNotFound, id, description)
}

func NewInternalServerError(id, description string) *AppError {
	return New(http.StatusInternalServerError, id, description)
}

func NewUnauthorizedError(id, description string) *AppError {
	return New(http.StatusUnauthorized, id, description)
}

// Predefined application errors
var (
	ErrInvalidRequestBody = NewBadRequestError("ErrInvalidRequestBody", "The request body is malformed or contains invalid data.")
	ErrInternalServer     = NewInternalServerError("ErrInternalServer", "An unexpected internal server error occurred.")
	// Example specific error
	ErrPaymentNotFound = NewNotFoundError("ErrPaymentNotFound", "payment id is not found")
	ErrValidation      = NewBadRequestError("ErrValidation", "Input validation failed.") // Generic validation
)

// Specific validation errors (can be created dynamically or predefined if common)
func NewValidationError(id, description string) *AppError {
	return New(http.StatusBadRequest, id, description)
}
