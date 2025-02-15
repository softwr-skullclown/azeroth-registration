package http

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/softwr-skullclown/azeroth-registration/ui"
)

func (o *Endpoints) handle() {
	uiFS := ui.New(o.config.UseOSFilesystem)
	// k8s healthcheck /healthz as per convention
	o.router.HandleFunc("GET /health", o.handleHealthz)

	apiRouter := http.NewServeMux()

	apiRouter.HandleFunc("GET /config", o.handleUIConfig)
	apiRouter.HandleFunc("POST /register", o.handleRegister)
	apiRouter.HandleFunc("POST /updatepwd", o.handleUpdatePassword)
	apiRouter.HandleFunc("POST /forgotpwd", o.handleForgotPassword)
	apiRouter.HandleFunc("GET /realms", o.handleRealmList)
	apiRouter.HandleFunc("GET /realms/{id}/online-characters", o.handleRealmOnlineCharacters)

	o.router.Handle("/api/", http.StripPrefix("/api", apiRouter))

	// index/static
	o.router.Handle("/", http.StripPrefix("/", spaHandler(uiFS)))
}

func (o *Endpoints) ListenAndServe(ctx context.Context) error {
	slog.InfoContext(ctx, fmt.Sprintf("Listening on %s", o.config.ListenAddress))
	return http.ListenAndServe(o.config.ListenAddress, o.router)
}

func New(config Config, authDbService AuthDBService, realmServices map[int]RealmDBService, emailService EmailService) *Endpoints {
	router := http.NewServeMux()
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
