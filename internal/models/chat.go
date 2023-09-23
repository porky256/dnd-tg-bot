package models

import "time"

type Chat struct {
	ID                 int
	ChatID             string
	CurrentCharacterID int
	Status             ChatStatus
	AddStatus          ChatStatusAdd
	UpdateStatus       ChatStatusUpdate
	CreatedAt          time.Time
	UpdatedAt          time.Time
}

type ChatStatus int

const (
	ChatStatusDefault ChatStatus = iota
	ChatStatusAdding
)

type ChatStatusAdd int

const (
	ChatStatusAddDefault ChatStatusAdd = iota
	ChatStatusAddName
	ChatStatusAddHP
	ChatStatusAddLevel
	ChatStatusAddArmor
	ChatStatusAddMoney
	ChatStatusAddAbilities
)

type ChatStatusUpdate int

const (
	ChatStatusUpdateDefault ChatStatusUpdate = iota
)
