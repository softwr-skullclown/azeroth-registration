package http

import (
	"fmt"
	"log/slog"
	"net/http"
)

// handleRealmList will return a list of the realms configured with online player counts
func (o *Endpoints) handleRealmList(w http.ResponseWriter, r *http.Request) {
	err := sendJsonResponse(w, `{"alive": true}`)
	if err != nil {
		slog.ErrorContext(r.Context(), fmt.Sprintf("error sending json response: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

// handleRealmPlayers will return a list of the players for the realm paginated
func (o *Endpoints) handleRealmPlayers(w http.ResponseWriter, r *http.Request) {
	err := sendJsonResponse(w, `{"alive": true}`)
	if err != nil {
		slog.ErrorContext(r.Context(), fmt.Sprintf("error sending json response: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
