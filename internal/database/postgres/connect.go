package postgres

import (
	"database/sql"
	"fmt"
	"github.com/porky256/dnd-tg-bot/internal/database"

	// driver for postgresql
	_ "github.com/lib/pq"
)

// ConnectPGSQL establishes connection to userProvider
func ConnectPGSQL(config database.DBConfig) (*database.DB, error) {
	db, err := sql.Open("postgres", buildPGConnString(config))

	if err != nil {
		return nil, fmt.Errorf("error occurred while connecting to userProvider: %w", err)
	}

	if err = db.Ping(); err != nil {
		return nil, fmt.Errorf("error occurred while accessing userProvider: %w", err)
	}

	db.SetMaxIdleConns(config.MaxIdleDBConn)
	db.SetMaxOpenConns(config.MaxOpenDBConn)

	return &database.DB{DB: db}, nil
}

// buildPGConnString forms conn string from config
func buildPGConnString(c database.DBConfig) string {
	return fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=%s",
		c.User, c.Password, c.Name, c.Host, c.Port, c.SSLMode)
}
