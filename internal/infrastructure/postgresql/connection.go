package postgresql

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/otie173/skyland-auth/internal/config"
)

func NewConnection(cfg *config.Config) (*sqlx.DB, error) {
	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName,
	)

	db, err := sqlx.Connect("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)
	db.SetConnMaxLifetime(5 * time.Minute)

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}
	log.Println("Database connected successfully")

	if err := runMigrations(db); err != nil {
		return nil, fmt.Errorf("migrations failed: %w", err)
	}

	return db, nil
}

func runMigrations(db *sqlx.DB) error {
	files, err := filepath.Glob("migrate/*.up.sql")
	if err != nil {
		return fmt.Errorf("failed to fing migrations file: %w", err)
	}

	if len(files) == 0 {
		fmt.Println("No migrations files found")
		return nil
	}

	for _, file := range files {
		log.Printf("Applying migration: %s\n", filepath.Base(file))

		schema, err := os.ReadFile(file)
		if err != nil {
			return fmt.Errorf("failed to read migration file %s: %w", file, err)
		}

		if _, err := db.Exec(string(schema)); err != nil {
			return fmt.Errorf("failed to execute migration %s: %w", file, err)
		}
		log.Printf("Applied: %s", filepath.Base(file))
	}

	log.Println("All migrations applied successfully")
	return nil
}
