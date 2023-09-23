package models

import "time"

type User struct {
	ID        int
	TgID      string
	ChatID    string
	Username  string
	CreatedAt time.Time
	UpdatedAt time.Time
}
