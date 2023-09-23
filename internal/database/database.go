package database

import "database/sql"

// DB wrapper for sql.DB
type DB struct {
	DB *sql.DB
}
