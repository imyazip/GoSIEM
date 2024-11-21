// Package storage предоставляет функционал работы с базой данных
package storage

import (
	"context"
	"database/sql"
)

type Storage struct {
	db *sql.DB // Объект для работы с базой данных
}

// NewStorage создает новый экземпляр Storage с переданным подключением к базе данных.
func NewStorage(db *sql.DB) *Storage {
	return &Storage{db: db}
}

// FindAPIKeyInStorage проверяет наличие API-ключа в базе данных.
func (s *Storage) FindAPIKeyInStorage(ctx context.Context, apiKey string) (bool, error) {
	var exists bool
	err := s.db.QueryRowContext(ctx, "SELECT EXISTS(SELECT 1 FROM api_keys WHERE key_value = ?)", apiKey).Scan(&exists)
	if err != nil {
		return false, err
	}
	return exists, nil
}
