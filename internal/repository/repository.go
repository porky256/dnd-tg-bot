package repository

import "github.com/porky256/dnd-tg-bot/internal/models"

type UserProvider interface {
	InsertUser(user *models.User) (int, error)

	GetUserByID(id int) (*models.User, error)
	GetUserByTgID(tgID string) (*models.User, error)

	UpdateUser(user *models.User) error

	DeleteUserByID(id int) error
}

type CharacterProvider interface {
	InsertCharacter(char *models.Character) (int, error)
	InsertCharactersAbilities(char *models.Character) (int, error)
	InsertCharactersSkillInsights(char *models.Character) (int, error)

	GetFullCharacterByID(charID int) (*models.Character, error)
	GetCharacterByID(charID int) (*models.Character, error)
	GetCharactersByOwnerTgID(ownerTgID string) ([]models.Character, error)
	GetCharactersAbilitiesByCharacterID(charID int) (*models.Abilities, error)
	GetCharactersSkillsInsightsByCharacterID(charID int) (*models.Skills, error)

	UpdateCharacter(char *models.Character) error
	UpdateCharactersAbilities(char *models.Character) error
	UpdateCharactersSkillInsights(char *models.Character) error

	DeleteCharacterByID(id int) error
	DeleteAbilityByCharacterID(id int) error
	DeleteSkillInsightsByCharacterID(id int) error
}

type ChatProvider interface {
	InsertChat(chat *models.Chat) (int, error)

	GetChatByChatID(id string) (*models.Chat, error)

	UpdateChatByChatID(chat *models.Chat) error

	DeleteChatByID(id int) error
}
