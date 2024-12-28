package storage

import (
	"bufio"
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/imyazip/GoSIEM/log-storage/internal/models"
	pb "github.com/imyazip/GoSIEM/log-storage/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
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

func (s *Storage) InsertSerializedLog(source string, logSerialized []byte, createdAt time.Time, sensorId string) error {
	query := `INSERT INTO serialized_logs (log_source, log_serialized, created_at_system, sensor_id) VALUES (?, ?, ?, ?)`
	_, err := s.db.Exec(query, source, logSerialized, createdAt, sensorId)
	if err != nil {
		return err
	}

	return nil
}

func (s *Storage) GetNewLogs(ctx context.Context, limit int32) ([]*pb.LogEntry, error) {
	rows, err := s.db.Query(`
		SELECT id, log_source, log_serialized, DATE_FORMAT(created_at_system, '%Y-%m-%dT%H:%i:%sZ') AS created_at_system, sensor_id
		FROM serialized_logs
		WHERE read_flag = FALSE
		LIMIT ?`, limit)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	var logs []*pb.LogEntry
	for rows.Next() {
		var id int
		var logSource, sensorID string
		var serializedData []byte
		var systemCreatedAt string

		if err := rows.Scan(&id, &logSource, &serializedData, &systemCreatedAt, &sensorID); err != nil {
			return nil, err
		}
		createdAt, err := time.Parse(time.RFC3339, systemCreatedAt)
		if err != nil {
			return nil, err
		}

		logs = append(logs, &pb.LogEntry{
			Id:              int64(id),
			LogSource:       logSource,
			LogSerialized:   serializedData,
			SystemCreatedAt: timestamppb.New(time.Time(createdAt)),
			SensorId:        sensorID,
		})

		// Помечаем лог как обработанный
		_, err = s.db.Exec("UPDATE serialized_logs SET read_flag = TRUE WHERE id = ?", id)
		if err != nil {
			return nil, err
		}
	}

	return logs, nil
}

func (s *Storage) AddSecurityEvent(ctx context.Context, event models.SecurityEvent) error {
	query := `INSERT INTO security_events ( event_type, event_description)
        VALUES (?, ?)`
	_, err := s.db.Exec(query, event.EventType, event.EventDescription)
	if err != nil {
		log.Printf("Ошибка при добавлении события: %v", err)
		return err
	}
	return nil
}

func (s *Storage) GetSecurityEvents(ctx context.Context, limit int32, onlyUnread bool) ([]*pb.SecurityEvent, error) {
	query := `
		SELECT id, event_type, event_description, detected_at, read_flag
		FROM security_events
	`
	if onlyUnread {
		query += " WHERE read_flag = FALSE"
	}
	query += " LIMIT ?"

	rows, err := s.db.QueryContext(ctx, query, limit)
	if err != nil {
		return nil, fmt.Errorf("failed to query security events: %w", err)
	}
	defer rows.Close()

	var events []*pb.SecurityEvent
	for rows.Next() {
		var id int64
		var eventType, eventDescription string
		var detectedAtRaw []byte
		var readFlag bool

		// Сканируем значения
		if err := rows.Scan(&id, &eventType, &eventDescription, &detectedAtRaw, &readFlag); err != nil {
			return nil, fmt.Errorf("failed to scan security event: %w", err)
		}

		// Преобразуем `detectedAtRaw` в строку и затем в `time.Time`
		detectedAtStr := string(detectedAtRaw)
		detectedAt, err := time.Parse("2006-01-02 15:04:05", detectedAtStr)
		if err != nil {
			return nil, fmt.Errorf("failed to parse detected_at: %w", err)
		}

		// Создаем объект `pb.SecurityEvent`
		events = append(events, &pb.SecurityEvent{
			Id:               id,
			EventType:        eventType,
			EventDescription: eventDescription,
			DetectedAt:       timestamppb.New(detectedAt),
			ReadFlag:         readFlag,
		})
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over rows: %w", err)
	}

	return events, nil
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

func (s *Storage) AddRule(ctx context.Context, ruleJSON string) (int64, error) {
	query := `INSERT INTO rules (rule) VALUES (?)`
	result, err := s.db.ExecContext(ctx, query, ruleJSON)
	if err != nil {
		return 0, fmt.Errorf("failed to add rule: %w", err)
	}

	ruleID, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("failed to get inserted rule ID: %w", err)
	}

	return ruleID, nil
}

func (s *Storage) DeleteRule(ctx context.Context, ruleID int64) error {
	query := `DELETE FROM rules WHERE id = ?`
	_, err := s.db.ExecContext(ctx, query, ruleID)
	if err != nil {
		return fmt.Errorf("failed to delete rule with ID %d: %w", ruleID, err)
	}
	return nil
}

func (s *Storage) GetRules(ctx context.Context) ([]models.Rule, error) {
	query := `SELECT id, rule FROM rules`
	rows, err := s.db.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed to query rules: %w", err)
	}
	defer rows.Close()

	var rules []models.Rule
	for rows.Next() {
		var rule models.Rule
		if err := rows.Scan(&rule.ID, &rule.JSON); err != nil {
			return nil, fmt.Errorf("failed to scan rule: %w", err)
		}
		rules = append(rules, rule)
	}

	return rules, nil
}
