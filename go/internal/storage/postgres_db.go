package storage

import (
	"database/sql"
	"fmt"

	"{{MODULE_NAME}}/internal/config"

	_ "github.com/lib/pq"
)

type PostgresDB struct {
	DB *sql.DB
}

func NewPostgresDB() (*PostgresDB, error) {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		config.DBHost,
		config.DBPort,
		config.DBUser,
		config.DBPassword,
		config.DBName,
		config.DBSSLMode,
	)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("error al abrir la conexión: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("error al conectar con postgres: %w", err)
	}

	return &PostgresDB{DB: db}, nil
}
