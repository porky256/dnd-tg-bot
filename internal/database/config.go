package database

import (
	"database/sql"
	"errors"
)

var ErrNoRowsDeleted = errors.New("no rows was deleted")

var ErrNoRowsUpdated = errors.New("no rows was updated")

var ErrNoRowsInserted = errors.New("no rows was inserted")

var ErrNoRows = sql.ErrNoRows

var ErrViolatesUnique = errors.New("trying to insert value that already exists")

// DBConfig config to connect to database
type DBConfig struct {
	DriverName    string
	MaxOpenDBConn int
	MaxIdleDBConn int
	Host          string
	Port          string
	Name          string
	User          string
	Password      string
	SSLMode       string
}
