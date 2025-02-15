package http

import (
	"context"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/softwr-skullclown/azeroth-registration/domain"
)

type Config struct {
	ListenAddress         string
	UseOSFilesystem       bool
	RealmIds              []int
	AllowMultipleAccounts bool
	UIConfig              UIConfig
}

type UIConfig struct {
	SiteTitle   string `json:"siteTitle"`
	RealmList   string `json:"realmList"`
	GameVersion string `json:"gameVersion"`
}

type AuthDBService interface {
	RealmList(ctx context.Context, ids []int) ([]domain.Realm, error)
	RegisterAccount(ctx context.Context, email string, username string, password string) (*domain.Account, error)
	UpdatePassword(ctx context.Context, username string, exitingPassword string, newPassword string) error
	GetAccountByName(ctx context.Context, username string) (*domain.Account, error)
	CheckEmailHasAccount(ctx context.Context, email string) (bool, error)
}

type RealmDBService interface {
	GetOnlineCharacters(ctx context.Context) ([]domain.Character, error)
}

type EmailService interface {
	SendWelcome(ctx context.Context, email string, username string) error
	SendPasswordReset(ctx context.Context, email string, username string, token string) error
	SendPasswordUpdated(ctx context.Context, email string, username string) error
}

// Endpoints represents the http service and its endpoints
type Endpoints struct {
	config          Config
	router          *http.ServeMux
	authDBSvc       AuthDBService
	realmDBServices map[int]RealmDBService
	emailService    EmailService
	validator       *validator.Validate
}

type registrationRequest struct {
	// Email must be valid email format
	Email string `json:"email" validate:"required,email"`
	// Username must be >= 2 chars and <= 16 chars matching regex /^[0-9A-Z-_]+$/
	Username string `json:"username" validate:"required,gte=2,lte=16,alphanum_dash_underscore"`
	// Password must between >= 4 chars and <= 16 chars
	Password string `json:"password" validate:"required,gte=4,lte=16"`
	// RePassword must match password and its requirements
	RePassword string `json:"repassword" validate:"required,eqfield=Password"`
}
