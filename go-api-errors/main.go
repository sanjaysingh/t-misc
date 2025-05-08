package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"go-api-errors/handlers"
	"go-api-errors/httputils"
)

// appHandler is a custom handler type that returns an error.
type appHandler func(http.ResponseWriter, *http.Request) error

// ServeHTTP makes appHandler satisfy the http.Handler interface.
// It calls the underlying appHandler and then uses httputils.RespondWithError
// to format the error response if an error occurs.
func (fn appHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := fn(w, r); err != nil {
		httputils.RespondWithError(w, err)
	}
}

func main() {
	r := chi.NewRouter()

	// Middleware
	r.Use(middleware.Logger)    // Logs requests
	r.Use(middleware.Recoverer) // Recovers from panics

	// API routes
	r.Route("/api", func(r chi.Router) {
		r.Method("GET", "/paymentplans", appHandler(handlers.GetPaymentPlans))
		r.Method("GET", "/payments", appHandler(handlers.GetPayments))
		r.Method("POST", "/payments", appHandler(handlers.AddPayment))
	})

	log.Println("Starting server on :3000...")
	err := http.ListenAndServe(":3000", r)
	if err != nil {
		log.Fatalf("Error starting server: %s\n", err)
	}
}
