package http

import (
	"context"

	"github.com/gorilla/mux"
	"github.com/softwr-skullclown/azeroth-registration/domain"
)

type Config struct {
	ListenAddress   string
	UseOSFilesystem bool
	RealmIds        []int
}

type AuthDBService interface {
	RealmList(ctx context.Context, ids []int) ([]domain.Realm, error)
	RegisterAccount(ctx context.Context, email string, username string, password string) error
	UpdatePassword(ctx context.Context, username string, exitingPassword string, newPassword string) error
	GetAccountByName(ctx context.Context, username string) error
}

type RealmDBService interface {
	GetOnlineCharacters(ctx context.Context) error
}

// Endpoints represents the http service and its endpoints
type Endpoints struct {
	config          Config
	router          *mux.Router
	authDBSvc       AuthDBService
	realmDBServices map[int]RealmDBService
}
