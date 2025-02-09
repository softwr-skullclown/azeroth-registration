package http

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"github.com/softwr-skullclown/azeroth-registration/ui"
)

func (o *Endpoints) handle() {
	uiFS := ui.New(o.config.UseOSFilesystem)
	// k8s healthcheck /healthz as per convention
	o.router.HandleFunc("/healthz", o.handleHealthz).Methods(http.MethodGet)

	apiRouter := o.router.PathPrefix("/api").Subrouter()

	apiRouter.HandleFunc("/register", o.handleRegister).Methods(http.MethodPost)
	apiRouter.HandleFunc("/updatepwd", o.handleUpdatePassword).Methods(http.MethodPost)
	apiRouter.HandleFunc("/forgotpwd", o.handleForgotPassword).Methods(http.MethodPost)
	apiRouter.HandleFunc("/realms", o.handleRealmList).Methods(http.MethodGet)
	apiRouter.HandleFunc("/realms/{id}/online-characters", o.handleRealmOnlineCharacters).Methods(http.MethodGet)

	// index/static
	o.router.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(uiFS)))
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
