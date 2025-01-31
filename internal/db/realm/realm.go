package realm

import (
	"context"
	"database/sql"
	"log/slog"

	"github.com/softwr-skullclown/azeroth-registration/domain"
)

type Service struct {
	DB *sql.DB
}

// GetOnlineCharacters returns a list of online characters from the realm
func (s *Service) GetOnlineCharacters(ctx context.Context) ([]domain.Character, error) {
	characters := make([]domain.Character, 0)

	rows, err := s.DB.QueryContext(ctx, `
		SELECT guid, name, race, class, level FROM characters WHERE online = 1;
	`)

	if err != nil {
		return characters, err
	}

	defer rows.Close()

	for rows.Next() {
		c := domain.Character{}
		if err := rows.Scan(&c.Guid, &c.Name, &c.Race, &c.Class, &c.Level); err != nil {
			slog.Error("error scanning character", slog.Any("error", err))
		} else {
			characters = append(characters, c)
		}
	}

	return characters, nil
}
