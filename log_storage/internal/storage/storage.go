package storage

import (
	"database/sql"
	"time"
)

type Storage struct {
	db *sql.DB
}

func NewStorage(db *sql.DB) *Storage {
	return &Storage{db: db}
}

func (s *Storage) InsertRawLog(source string, logString string, createdAt time.Time, sensorId string) error {
	query := `INSERT INTO raw_logs (log_source, log_string, created_at_system, sensor_id) VALUES (?, ?, ?, ?)`
	_, err := s.db.Exec(query, source, logString, createdAt, sensorId)
	if err != nil {
		return err
	}

	return nil
}

func (s *Storage) InsertSerializedLog(source string, logSerialized string, createdAt time.Time, sensorId string) error {
	query := `INSERT INTO serialized_logs (log_source, log_serialized, created_at_system, sensor_id) VALUES (?, ?, ?, ?)`
	_, err := s.db.Exec(query, source, logSerialized, createdAt, sensorId)
	if err != nil {
		return err
	}

	return nil
}
