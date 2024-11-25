// Package storage предоставляет функционал работы с базой данных
package storage

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"

	"github.com/imyazip/GoSIEM/auth/internal/models"
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
func (s *Storage) InsertUser(ctx context.Context, userName string, hashedPassword string, roleID int64) error {
	_, err := s.db.ExecContext(ctx, "INSERT INTO users (username, password, role_id) VALUES (?, ?, ?)", userName, hashedPassword, roleID)
	if err != nil {
		return err
	}

	return nil
}

func (s *Storage) GetUserByUsername(ctx context.Context, userName string) (*models.User, error) {
	query := `
        SELECT id, username, password, role_id, created_at, updated_at 
        FROM users 
        WHERE username = ?;
    `
	row := s.db.QueryRowContext(ctx, query, userName)

	user := &models.User{}

	err := row.Scan(&user.ID, &user.Username, &user.Password, &user.RoleID, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			log.Printf("Пользователь с username '%s' не найден", userName)
			return nil, err
		}
		log.Printf("Ошибка выполнения запроса для пользователя '%s': %v", userName, err)
		return nil, err
	}

	return user, nil
}

// GetRoleNameByID возвращает имя роли по её ID.
func (s *Storage) GetRoleNameByID(ctx context.Context, roleID int64) (string, error) {
	var roleName string

	// SQL-запрос для получения имени роли
	query := "SELECT role_name FROM roles WHERE id = ?"

	// Выполнение запроса с передачей roleID
	err := s.db.QueryRowContext(ctx, query, roleID).Scan(&roleName)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", fmt.Errorf("role with ID %d not found", roleID)
		}
		return "", fmt.Errorf("failed to query role name: %w", err)
	}

	return roleName, nil
}

func (s *Storage) CheckSensorExists(ctx context.Context, sensorID string) (bool, error) {
	// Запрос для проверки существования сенсора в базе данных
	var count int
	err := s.db.QueryRowContext(ctx, "SELECT COUNT(*) FROM sensors WHERE sensor_id = ?", sensorID).Scan(&count)
	if err != nil {
		return false, fmt.Errorf("error checking if sensor exists: %v", err)
	}

	// Возвращаем true, если сенсор найден (count > 0), иначе false
	return count > 0, nil
}

func (s *Storage) InsertSensor(ctx context.Context, sensor models.Sensor) error {
	_, err := s.db.ExecContext(ctx, "INSERT INTO sensors (sensor_id, name, hostname,os_version, sensor_type, agent_version) VALUES (?, ?, ?, ?, ?, ?)", sensor.Sensor_id, sensor.Name, sensor.Hostname, sensor.Os_version, sensor.Sensor_type, sensor.Agent_version)
	if err != nil {
		return err
	}

	return nil
}

// UpdateSensor обновляет данные о сенсоре в базе данных по sensor_id
func (s *Storage) UpdateSensor(ctx context.Context, sensor models.Sensor) error {
	// Проверка, что sensor_id указан
	if sensor.Sensor_id == "" {
		return errors.New("sensor_id is required")
	}

	// SQL-запрос для обновления данных о сенсоре
	query := `
		UPDATE sensors
		SET 
			name = ?, 
			hostname = ?,
			os_version = ?, 
			sensor_type = ?, 
			agent_version = ?
		WHERE sensor_id = ?`

	// Выполнение запроса с использованием контекста
	_, err := s.db.ExecContext(ctx, query,
		sensor.Name,
		sensor.Hostname,
		sensor.Os_version,
		sensor.Sensor_type,
		sensor.Agent_version,
		sensor.Sensor_id,
	)
	if err != nil {
		return fmt.Errorf("failed to update sensor: %w", err)
	}

	return nil
}
