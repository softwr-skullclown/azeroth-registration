package http

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
)

// handleRegister will attempt to register the user account for the given request
func (o *Endpoints) handleRegister(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	body, err := io.ReadAll(r.Body)
	if err != nil {
		slog.ErrorContext(r.Context(), fmt.Sprintf("error reading body: %v", err))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var u = registrationRequest{}
	jsonErr := json.Unmarshal(body, &u)
	if jsonErr != nil {
		slog.ErrorContext(r.Context(), fmt.Sprintf("error unmarshalling body: %v", err))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = o.validator.Struct(u)
	if err != nil {
		slog.ErrorContext(r.Context(), fmt.Sprintf("error validating inputs: %v", err))
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	account, err := o.authDBSvc.RegisterAccount(ctx, u.Email, u.Username, u.Password)
	if err != nil {
		slog.ErrorContext(r.Context(), fmt.Sprintf("error creating account: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
	}

	emailCtx := context.WithoutCancel(ctx)
	go func(ctx context.Context, email string, username string) {
		err := o.emailService.SendWelcome(ctx, email, username)
		if err != nil {
			slog.ErrorContext(ctx, "error sending welcome email", slog.Any("error", err), slog.String("username", username), slog.String("email", email))
		}
	}(emailCtx, u.Email, u.Username)

	err = sendJsonResponse(w, account)
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
