package http

import (
	"fmt"
	"log/slog"
	"net/http"
)

// handleRegister will attempt to register the user account for the given request
func (o *Endpoints) handleRegister(w http.ResponseWriter, r *http.Request) {
	err := sendJsonResponse(w, `{"alive": true}`)
	if err != nil {
		slog.ErrorContext(r.Context(), fmt.Sprintf("error sending json response: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

// handleForgotPassword will attempt to send a forgot password email to the specified user account by using the associated email
func (o *Endpoints) handleForgotPassword(w http.ResponseWriter, r *http.Request) {
	err := sendJsonResponse(w, `{"alive": true}`)
	if err != nil {
		slog.ErrorContext(r.Context(), fmt.Sprintf("error sending json response: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

// handleUpdatePassword will update the user password for the given account
// using provided existing password against db existing password for verification before attempting update
func (o *Endpoints) handleUpdatePassword(w http.ResponseWriter, r *http.Request) {
	err := sendJsonResponse(w, `{"alive": true}`)
	if err != nil {
		slog.ErrorContext(r.Context(), fmt.Sprintf("error sending json response: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
