package storage

import (
	"context"
	"database/sql"
)

type Storage struct {
	db *sql.DB
}

func NewStorage(db *sql.DB) *Storage {
	return &Storage{db: db}
}

func (s *Storage) ValidateAPIKey(ctx context.Context, apiKey string) (bool, error) {
	var exists bool
	err := s.db.QueryRowContext(ctx, "SELECT EXISTS(SELECT 1 FROM api_keys WHERE key_value = ?)", apiKey).Scan(&exists)
	if err != nil {
		return false, err
	}
	return exists, nil
}
