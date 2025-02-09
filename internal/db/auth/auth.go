package auth

import (
	"context"
	"database/sql"
	"fmt"
	"log/slog"
	"slices"
	"strings"

	"github.com/softwr-skullclown/azeroth-registration/domain"
)

type Service struct {
	DB *sql.DB
}

// RealmList returns a list of realms from the Auth database
func (s *Service) RealmList(ctx context.Context, ids []int) ([]domain.Realm, error) {
	realms := make([]domain.Realm, 0)

	rows, err := s.DB.QueryContext(ctx, `
		SELECT id, name, flag, icon, population FROM realmlist;
	`)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		r := domain.Realm{}
		if err := rows.Scan(&r.Id, &r.Name, &r.Flag, &r.Icon, &r.Population); err != nil {
			slog.Error(fmt.Sprintf("error scanning realm: %v", err))
		} else {
			if slices.Contains(ids, r.Id) {
				realms = append(realms, r)
			}
		}
	}

	return realms, nil
}

// RegisterAccount attempts to create a new auth account
func (s *Service) RegisterAccount(ctx context.Context, email string, username string, password string) (*domain.Account, error) {
	account := domain.Account{
		Email:    email,
		Username: username,
	}

	salt, err := salt()
	if err != nil {
		return nil, err
	}

	verifier := calculateSRP6Verifier(username, password, salt)

	_, err = s.DB.ExecContext(ctx, `INSERT INTO account (email, username, salt, verifier) VALUES (?, ?, ?, ?);`,
		strings.ToUpper(email), strings.ToUpper(username), salt, verifier)
	if err != nil {
		return nil, err
	}

	err = s.DB.QueryRowContext(ctx, `SELECT id FROM account WHERE username = ?`, strings.ToUpper(username)).Scan(&account.Id)
	if err != nil {
		return nil, err
	}

	return &account, nil
}

// UpdatePassword updates an accounts password if the existing passwords match
func (s *Service) UpdatePassword(ctx context.Context, username string, exitingPassword string, newPassword string) error {
	return nil
}

// GetAccountByName gets the auth account by name to be used in forgot password flow
func (s *Service) GetAccountByName(ctx context.Context, username string) (*domain.Account, error) {
	account := domain.Account{}
	err := s.DB.QueryRowContext(ctx,
		`SELECT id, username, email FROM account WHERE username = ?`,
		strings.ToUpper(username),
	).Scan(&account.Id, &account.Username, &account.Email)
	if err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	return &account, nil
}

// CheckEmailHasAccount checks if the provided email has any existing accounts
func (s *Service) CheckEmailHasAccount(ctx context.Context, email string) (bool, error) {
	count := 0

	err := s.DB.QueryRowContext(ctx, `SELECT count(id) FROM account WHERE email = ?`, strings.ToUpper(email)).Scan(&count)
	if err != nil {
		return count > 0, err
	}

	return count > 0, nil
}
