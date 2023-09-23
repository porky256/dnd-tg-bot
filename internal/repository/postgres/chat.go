package postgres

import (
	"context"
	"database/sql"
	"github.com/porky256/dnd-tg-bot/internal/database"
	"github.com/porky256/dnd-tg-bot/internal/models"
	"time"
)

// PGChatProvider implements GlobalDatabaseProvider
type PGChatProvider struct {
	DB           *sql.DB
	QueryTimeout time.Duration
}

// NewPGChatProvider creates a new postgres userProvider entity
func NewPGChatProvider(db *sql.DB, timeout time.Duration) *PGChatProvider {
	return &PGChatProvider{
		DB:           db,
		QueryTimeout: timeout,
	}
}

func (db *PGChatProvider) InsertChat(chat *models.Chat) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), db.QueryTimeout)
	defer cancel()
	stmt := `INSERT INTO  chats (chat_id, current_character_id, status, add_substatus, update_substatus) 
			 VALUES ($1, $2, $3, $4, $5) RETURNING id`

	var newID int
	err := db.DB.QueryRowContext(ctx, stmt,
		chat.ChatID,
		chat.CurrentCharacterID,
		chat.Status,
		chat.AddStatus,
		chat.UpdateStatus,
	).Scan(&newID)

	if err != nil {
		return 0, ParsePqError(err)
	}
	return newID, nil
}

func (db *PGChatProvider) GetChatByID(id int) (*models.Chat, error) {
	ctx, cancel := context.WithTimeout(context.Background(), db.QueryTimeout)

	defer cancel()

	stmt := `SELECT id, chat_id, current_character_id, status, add_substatus, update_substatus, created_at, updated_at FROM chats WHERE id=$1`
	chat := new(models.Chat)
	row := db.DB.QueryRowContext(ctx, stmt, id)
	err := row.Scan(
		&chat.ID,
		&chat.ChatID,
		&chat.CurrentCharacterID,
		&chat.Status,
		&chat.AddStatus,
		&chat.UpdateStatus,
		&chat.CreatedAt,
		&chat.UpdatedAt,
	)

	if err != nil {
		return nil, ParsePqError(err)
	}

	return chat, nil
}

func (db *PGChatProvider) GetChatByChatID(id string) (*models.Chat, error) {
	ctx, cancel := context.WithTimeout(context.Background(), db.QueryTimeout)

	defer cancel()

	stmt := `SELECT id, chat_id, current_character_id, status, add_substatus, update_substatus, created_at, updated_at FROM chats WHERE chat_id=$1`
	chat := new(models.Chat)
	row := db.DB.QueryRowContext(ctx, stmt, id)
	err := row.Scan(
		&chat.ID,
		&chat.ChatID,
		&chat.CurrentCharacterID,
		&chat.Status,
		&chat.AddStatus,
		&chat.UpdateStatus,
		&chat.CreatedAt,
		&chat.UpdatedAt,
	)

	if err != nil {
		return nil, ParsePqError(err)
	}

	return chat, nil
}

func (db *PGChatProvider) UpdateChatByChatID(chat *models.Chat) error {
	ctx, cancel := context.WithTimeout(context.Background(), db.QueryTimeout)
	defer cancel()

	stmt := `UPDATE chats SET current_character_id = $1, status = $2, add_substatus = $3, update_substatus = $4 WHERE chat_id=$5`
	res, err := db.DB.ExecContext(ctx, stmt,
		chat.CurrentCharacterID,
		chat.Status,
		chat.AddStatus,
		chat.UpdateStatus,
		chat.ChatID,
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

func (db *PGChatProvider) DeleteChatByID(id int) error {
	ctx, cancel := context.WithTimeout(context.Background(), db.QueryTimeout)

	defer cancel()

	res, err := db.DB.ExecContext(ctx, "DELETE FROM chats WHERE id=$1", id)

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
