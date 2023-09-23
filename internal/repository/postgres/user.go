package postgres

import (
	"context"
	"database/sql"
	"github.com/porky256/dnd-tg-bot/internal/database"
	"github.com/porky256/dnd-tg-bot/internal/models"
	"time"
)

// PGUserProvider implements GlobalDatabaseProvider
type PGUserProvider struct {
	DB           *sql.DB
	QueryTimeout time.Duration
}

// NewPGUserProvider creates a new postgres userProvider entity
func NewPGUserProvider(db *sql.DB, timeout time.Duration) *PGUserProvider {
	return &PGUserProvider{
		DB:           db,
		QueryTimeout: timeout,
	}
}

func (db *PGUserProvider) InsertUser(user *models.User) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), db.QueryTimeout)
	defer cancel()
	stmt := `INSERT INTO  users (telegram_id, chat_id, username) 
			 VALUES ($1, $2, $3) RETURNING id`

	var newID int
	err := db.DB.QueryRowContext(ctx, stmt,
		user.TgID,
		user.ChatID,
		user.Username,
	).Scan(&newID)

	if err != nil {
		return 0, ParsePqError(err)
	}
	return newID, nil
}

func (db *PGUserProvider) GetUserByID(id int) (*models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), db.QueryTimeout)

	defer cancel()

	stmt := `SELECT id, telegram_id, chat_id, username, created_at, updated_at FROM users WHERE id=$1`
	user := new(models.User)
	row := db.DB.QueryRowContext(ctx, stmt, id)
	err := row.Scan(
		&user.ID,
		&user.TgID,
		&user.ChatID,
		&user.Username,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		return nil, ParsePqError(err)
	}

	return user, nil
}

func (db *PGUserProvider) GetUserByTgID(tgID string) (*models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), db.QueryTimeout)

	defer cancel()

	stmt := `SELECT id, telegram_id, chat_id, username, created_at, updated_at FROM users WHERE telegram_id=$1`
	user := new(models.User)
	row := db.DB.QueryRowContext(ctx, stmt, tgID)
	err := row.Scan(
		&user.ID,
		&user.TgID,
		&user.ChatID,
		&user.Username,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		return nil, ParsePqError(err)
	}

	return user, nil
}

func (db *PGUserProvider) UpdateUser(user *models.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), db.QueryTimeout)
	defer cancel()

	stmt := `UPDATE users SET telegram_id = $1, chat_id = $2, username = $3  WHERE id=$4`
	res, err := db.DB.ExecContext(ctx, stmt,
		user.TgID,
		user.ChatID,
		user.Username,
		user.ID,
	)

	if err != nil {
		return ParsePqError(err)
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return ParsePqError(err)
	}

	if rowsAffected == 0 {
		return database.ErrNoRowsUpdated
	}

	return nil
}

func (db *PGUserProvider) DeleteUserByID(id int) error {
	ctx, cancel := context.WithTimeout(context.Background(), db.QueryTimeout)

	defer cancel()

	res, err := db.DB.ExecContext(ctx, "DELETE FROM users WHERE id=$1", id)

	if err != nil {
		return ParsePqError(err)
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return ParsePqError(err)
	}

	if rowsAffected == 0 {
		return database.ErrNoRowsDeleted
	}

	return nil
}
