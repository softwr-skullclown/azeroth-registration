package auth

import (
	"context"
	"database/sql"
	"fmt"
	"log/slog"

	"github.com/softwr-skullclown/azeroth-registration/domain"
)

type Service struct {
	DB *sql.DB
}

// RealmList returns a list of realms from the Auth database
func (s *Service) RealmList(ctx context.Context, ids []int) ([]domain.Realm, error) {
	realms := make([]domain.Realm, 0)

	rows, err := s.DB.QueryContext(ctx, `
		SELECT id, name, flag, icon, population WHERE id IN($1);
	`, ids)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		r := domain.Realm{}
		if err := rows.Scan(&r.Id, &r.Name, &r.Flag, &r.Icon, &r.Population); err != nil {
			slog.Error(fmt.Sprintf("error scanning realm: %v", err))
		} else {
			realms = append(realms, r)
		}
	}

	return realms, nil
}

// RegisterAccount attempts to create a new auth account
func (s *Service) RegisterAccount(ctx context.Context, email string, username string, password string) error {
	return nil
}

// UpdatePassword updates an accounts password if the existing passwords match
func (s *Service) UpdatePassword(ctx context.Context, username string, exitingPassword string, newPassword string) error {
	return nil
}

// GetAccountByName gets the auth account by name to be used in forgot password flow
func (s *Service) GetAccountByName(ctx context.Context, username string) error {
	return nil
}
