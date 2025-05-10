package httputils

import (
	"encoding/json"
	"log"
	"net/http"

	"go-api-errors/apperrors"
)

// RespondWithError handles writing error responses in a structured JSON format.
func RespondWithError(w http.ResponseWriter, err error) {
	var appErr *apperrors.AppError
	var ok bool

	// Check if the error is an APIError
	if appErr, ok = err.(*apperrors.AppError); !ok {
		// If not, default to an internal server error
		log.Printf("Unhandled error: %v", err) // Log the original error for debugging
		appErr = apperrors.ErrInternalServer
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(appErr.StatusCode)

	response := apperrors.JSONErrorResponse{
		Notice: apperrors.ErrorNotice{
			ID:          appErr.ID,
			Description: appErr.Description,
		},
	}

	if encodeErr := json.NewEncoder(w).Encode(response); encodeErr != nil {
		// If encoding fails, log it and send a plain text error as a last resort.
		log.Printf("Error encoding JSON error response: %v", encodeErr)
		w.WriteHeader(http.StatusInternalServerError) // Ensure status code is set again if header was overwritten
		// Note: We might have already written headers, so this http.Error might not write the body if so.
		// It's a best-effort at this point.
		http.Error(w, "{\"notice\":{\"id\":\"InternalServerError\",\"description\":\"Error preparing error response.\"}}", http.StatusInternalServerError)
	}
}

// RespondWithJSON handles writing successful JSON responses.
func RespondWithJSON(w http.ResponseWriter, statusCode int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(statusCode)
	if payload != nil {
		if err := json.NewEncoder(w).Encode(payload); err != nil {
			log.Printf("Error encoding JSON success response: %v", err)
			// If encoding fails, we might be in a bad state. Try to send an internal server error.
			// This might not work if headers/status are already sent.
			RespondWithError(w, apperrors.ErrInternalServer)
		}
	}
}
