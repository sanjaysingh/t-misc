package handlers

import (
	"net/http"

	"go-api-errors/httputils"
)

// GetPaymentPlans returns a list of supported payment plans.
func GetPaymentPlans(w http.ResponseWriter, r *http.Request) error {
	dummyPaymentPlans := []string{"monthly", "bi-weekly", "annually"}
	httputils.RespondWithJSON(w, http.StatusOK, dummyPaymentPlans)
	return nil
}
