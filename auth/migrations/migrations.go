package migrations

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

// RunMigrations выполняет SQL-миграции из файла init.sql
func RunMigrations(dsn string) error {
	// Подключаемся к базе данных
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return fmt.Errorf("failed to connect to database: %v", err)
	}
	defer db.Close()

	// Читаем SQL-миграции из файла init.sql
	sqlFile := "/app/auth/migrations/init.sql" // Путь до файла миграции в контейнере
	sqlData, err := ioutil.ReadFile(sqlFile)
	if err != nil {
		return fmt.Errorf("failed to read migration file: %v", err)
	}

	// Выполняем миграции
	_, err = db.Exec(string(sqlData))
	if err != nil {
		return fmt.Errorf("error applying migrations: %v", err)
	}

	log.Println("Migrations applied successfully")
	return nil
}
