package http

import (
	"fmt"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
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
// @TODO - paginate results
func (o *Endpoints) handleRealmOnlineCharacters(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	rawRealmId := vars["id"]
	realmId, err := strconv.Atoi(rawRealmId)
	if err != nil {
		slog.Warn("bad realm id", slog.String("realm_id", rawRealmId))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	characters, err := o.realmDBServices[realmId].GetOnlineCharacters(r.Context())
	if err != nil {
		slog.ErrorContext(r.Context(), fmt.Sprintf("error getting realm online characters list: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = sendJsonResponse(w, characters)
	if err != nil {
		slog.ErrorContext(r.Context(), fmt.Sprintf("error sending json response: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
