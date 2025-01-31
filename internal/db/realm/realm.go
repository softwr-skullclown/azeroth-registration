package realm

import (
	"context"
	"database/sql"
)

type Service struct {
	DB *sql.DB
}

// GetOnlineCharacters returns a list of online characters from the realm
func (s *Service) GetOnlineCharacters(ctx context.Context) error {
	return nil
}
