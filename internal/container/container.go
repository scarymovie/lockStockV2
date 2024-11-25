package container

import (
	"database/sql"
	"fmt"
	"lockStockV2/internal/config"
	"log"

	_ "github.com/lib/pq" // PostgreSQL драйвер
)

type Container struct {
	DB *sql.DB
}

// Initialize создает и возвращает контейнер с инициализированными зависимостями
func Initialize(cfg *config.Config) *Container {
	// Создаём подключение к базе данных
	db, err := initializeDB(cfg)
	if err != nil {
		log.Fatalf("could not initialize database: %v", err)
	}

	// Возвращаем контейнер с инициализированными зависимостями
	return &Container{
		DB: db,
	}
}

// initializeDB создает подключение к PostgreSQL
func initializeDB(cfg *config.Config) (*sql.DB, error) {
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.Database.Host, cfg.Database.Port, cfg.Database.User, cfg.Database.Password, cfg.Database.Name)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("failed to open database connection: %w", err)
	}

	// Проверяем подключение
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	log.Println("Database connection established")
	return db, nil
}
