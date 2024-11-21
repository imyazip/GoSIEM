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

// RoleExists проверяет существование роли и возвращает ее ID
func (s *Storage) RoleExists(ctx context.Context, roleName string) (int, error) {
	var roleID int
	err := s.db.QueryRowContext(ctx, "SELECT id FROM roles WHERE role_name = $1", roleName).Scan(&roleID)
	if err != nil {
		return -1, err
	}

	return roleID, nil
}

// UserExists проверяет существование пользователя
// Возвращает false в случае отсутствия пользователя с указанным username в БД
func (s *Storage) UserExists(ctx context.Context, userName string) (bool, error) {
	var exists bool
	err := s.db.QueryRowContext(ctx, "SELECT COUNT(*) FROM users WHERE username = ?", userName).Scan(&exists)
	if err != nil {
		return true, err
	}

	return false, nil
}

// InsertUser добавляет нового пользователя в БД,
// Возвращает err в случае неудачи, nil при успешном добавлении
func (s *Storage) InsertUser(ctx context.Context, userName string, hashedPassword string, roleID int) error {
	_, err := s.db.ExecContext(ctx, "INSERT INTO users (username, password, role_id) VALUES (?, ?, ?)", userName, hashedPassword, roleID)
	if err != nil {
		return err
	}

	return nil
}
