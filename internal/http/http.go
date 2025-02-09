package http

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"github.com/softwr-skullclown/azeroth-registration/ui"
)

func (o *Endpoints) handle() {
	uiFS := ui.New(o.config.UseOSFilesystem)
	// k8s healthcheck /healthz as per convention
	o.router.HandleFunc("/healthz", o.handleHealthz).Methods(http.MethodGet)

	apiRouter := o.router.PathPrefix("/api").Subrouter()

	apiRouter.HandleFunc("/config", o.handleUIConfig).Methods(http.MethodGet)
	apiRouter.HandleFunc("/register", o.handleRegister).Methods(http.MethodPost)
	apiRouter.HandleFunc("/updatepwd", o.handleUpdatePassword).Methods(http.MethodPost)
	apiRouter.HandleFunc("/forgotpwd", o.handleForgotPassword).Methods(http.MethodPost)
	apiRouter.HandleFunc("/realms", o.handleRealmList).Methods(http.MethodGet)
	apiRouter.HandleFunc("/realms/{id}/online-characters", o.handleRealmOnlineCharacters).Methods(http.MethodGet)

	// index/static
	o.router.PathPrefix("/").Handler(http.StripPrefix("/", spaHandler(uiFS)))
}

func (o *Endpoints) ListenAndServe(ctx context.Context) error {
	slog.InfoContext(ctx, fmt.Sprintf("Listening on %s", o.config.ListenAddress))
	return http.ListenAndServe(o.config.ListenAddress, o.router)
}

func New(config Config, authDbService AuthDBService, realmServices map[int]RealmDBService, emailService EmailService) *Endpoints {
	router := mux.NewRouter()
	validate := validator.New(validator.WithRequiredStructEnabled())
	validate.RegisterValidation("alphanum_dash_underscore", isAlphanumDashUnderscore)

	e := &Endpoints{
		config:          config,
		router:          router,
		authDBSvc:       authDbService,
		realmDBServices: realmServices,
		emailService:    emailService,
		validator:       validate,
	}
	e.handle()
	return e
}

func spaHandler(uiFS http.FileSystem) http.Handler {
	fileServer := http.FileServer(uiFS)

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Get the absolute path to check if file exists
		path := r.URL.Path

		// Try to open the file from the filesystem
		f, err := uiFS.Open(strings.TrimPrefix(path, "/"))
		if err != nil {
			// If file doesn't exist, serve index.html
			if os.IsNotExist(err) {
				// Reset the path to serve index.html
				r.URL.Path = "/"
			}
		} else {
			// Don't forget to close the file
			f.Close()
		}

		// Serve either the requested file or index.html
		fileServer.ServeHTTP(w, r)
	})
}

// handleUIConfig returns a json response containing ui configuration values
func (o *Endpoints) handleUIConfig(w http.ResponseWriter, r *http.Request) {
	err := sendJsonResponse(w, o.config.UIConfig)
	if err != nil {
		slog.ErrorContext(r.Context(), fmt.Sprintf("error sending ui config json response: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
