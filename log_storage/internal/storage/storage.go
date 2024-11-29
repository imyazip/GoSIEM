package storage

import (
	"bufio"
	"context"
	"database/sql"
	"log"
	"os"
	"path/filepath"
	"strings"
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

func (s *Storage) ExecuteMigrations(ctx context.Context, migrationsDir string) error {
	// Открываем директорию с миграциями
	files, err := os.ReadDir(migrationsDir)
	if err != nil {
		log.Printf("failed to read migrations directory: %v", err)
		return err
	}

	// Проходим по всем файлам в директории
	for _, file := range files {
		if filepath.Ext(file.Name()) != ".sql" {
			continue // Пропускаем файлы, не являющиеся .sql
		}

		// Открываем каждый файл для чтения
		filePath := filepath.Join(migrationsDir, file.Name())
		migrationFile, err := os.Open(filePath)
		if err != nil {
			log.Printf("failed to open migration file %s: %v", filePath, err)
			return err
		}
		defer migrationFile.Close()

		// Читаем содержимое файла
		var queries []string
		scanner := bufio.NewScanner(migrationFile)
		var currentQuery strings.Builder
		for scanner.Scan() {
			line := scanner.Text()

			// Игнорируем пустые строки и комментарии
			if strings.TrimSpace(line) == "" || strings.HasPrefix(line, "--") {
				continue
			}

			// Собираем SQL-запросы
			currentQuery.WriteString(line)
			if strings.HasSuffix(strings.TrimSpace(line), ";") {
				queries = append(queries, currentQuery.String())
				currentQuery.Reset()
			}
		}

		if err := scanner.Err(); err != nil {
			log.Printf("error reading migration file %s: %v", filePath, err)
			return err
		}

		// Выполняем каждый запрос из миграции
		for _, query := range queries {
			_, err := s.db.ExecContext(ctx, query)
			if err != nil {
				log.Printf("failed to execute query from migration file %s: %v", filePath, err)
				return err
			}
		}

		log.Printf("Successfully executed migration file: %s", file.Name())
	}

	return nil
}
