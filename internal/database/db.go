package database

import (
	"database/sql"
	"demo09/internal/config"
	"fmt"
	"time"

	_ "github.com/lib/pq"
)

func NewPostgresDB(cfg config.DataBaseConfig) (*sql.DB, error) {
	conn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.Name, cfg.SSLMode)

	db, err := sql.Open("postgres", conn)
	if err != nil {
		return nil, fmt.Errorf("Abriendo Postgres conexión: %s", err)
	}

	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(5)
	db.SetConnMaxLifetime(30 * time.Minute)

	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("Ping Postgres %s", err)
	}

	return db, nil
}
