// pkg/database/postgres.go

package database

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib" // pgx driver registered for database/sql
)

// Database wraps the sql.DB connection pool
type Database struct {
	Conn *sql.DB
}

// Config holds the values needed to build a DSN
type Config struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	SSLMode  string // "disable" for local dev, "require" for prod
}

// NewPostgresDB opens a connection pool and verifies it with a ping
func NewPostgresDB(cfg Config) (*Database, error) {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.DBName, cfg.SSLMode,
	)

	conn, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to open db connection: %w", err)
	}

	// connection pool tuning
	conn.SetMaxOpenConns(25)
	conn.SetMaxIdleConns(25)
	conn.SetConnMaxLifetime(5 * time.Minute)

	if err := conn.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping db: %w", err)
	}

	log.Println("database connection established")

	return &Database{Conn: conn}, nil
}

// Close closes the underlying connection pool
func (d *Database) Close() error {
	return d.Conn.Close()
}
