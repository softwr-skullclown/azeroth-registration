package http

import (
	"fmt"
	"log/slog"
	"net/http"
)

// handleRealmList will return a list of the realms configured with online character counts
func (o *Endpoints) handleRealmList(w http.ResponseWriter, r *http.Request) {
	realms, err := o.authDBSvc.RealmList(r.Context(), o.config.RealmIds)
	if err != nil {
		slog.ErrorContext(r.Context(), fmt.Sprintf("error getting realm list: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	err = sendJsonResponse(w, realms)
	if err != nil {
		slog.ErrorContext(r.Context(), fmt.Sprintf("error sending json response: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

// handleRealmOnlineCharacters will return a list of the online characters for the realm paginated
func (o *Endpoints) handleRealmOnlineCharacters(w http.ResponseWriter, r *http.Request) {
	err := sendJsonResponse(w, `{"alive": true}`)
	if err != nil {
		slog.ErrorContext(r.Context(), fmt.Sprintf("error sending json response: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
