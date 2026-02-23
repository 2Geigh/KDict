package db

import (
	"database/sql"
	"embed"
	"fmt"
	"os"

	_ "github.com/mattn/go-sqlite3"
	"github.com/pressly/goose/v3"
)

var (
	// Use go:embed to bundle migrations into the binary
	// This assumes your migrations are in a folder named 'migrations'
	//go:embed migrations/*.sql
	embedMigrations embed.FS

	DB *sql.DB = nil
)

func InitDB() error {
	dbFilePath := "db/database.db"

	// Check if the file exists
	_, err := os.Stat(dbFilePath)
	if os.IsNotExist(err) {
		file, err := os.Create(dbFilePath)
		if err != nil {
			return fmt.Errorf("failed to create database file: %v", err)
		}
		file.Close() // File must be closed after creation
	}

	DB, err = sql.Open("sqlite3", dbFilePath)
	if err != nil {
		return fmt.Errorf("failed to open database: %v", err)
	}

	err = DB.Ping()
	if err != nil {
		return fmt.Errorf("failed to reach database: %w", err)
	}

	err = runMigrations()
	if err != nil {
		return fmt.Errorf("couldn't migrate database: %w", err)
	}

	return nil
}

func runMigrations() error {
	goose.SetBaseFS(embedMigrations)

	err := goose.SetDialect("sqlite")
	if err != nil {
		return fmt.Errorf("goose couldn't set SQL dialect: %w", err)
	}

	err = goose.Up(DB, "migrations")
	if err != nil {
		return fmt.Errorf("goose up failed: %w", err)
	}

	return nil
}
