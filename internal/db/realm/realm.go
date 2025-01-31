package realm

import (
	"context"
	"database/sql"
)

type Service struct {
	DB *sql.DB
}

// GetOnlineCharacterCount returns a count of online characters
func (s *Service) GetOnlineCharacterCount(ctx context.Context) (int, error) {
	count := 0

	err := s.DB.QueryRow(`SELECT count(guid) FROM characters WHERE online = 1`).Scan(&count)
	if err != nil {
		return count, err
	}

	return count, nil
}

// GetOnlineCharacters returns a list of online characters from the realm
func (s *Service) GetOnlineCharacters(ctx context.Context) error {
	return nil
}
