package handlers

import (
	"context"
	"github.com/go-telegram/bot"
	bmodels "github.com/go-telegram/bot/models"
	"github.com/porky256/dnd-tg-bot/internal/database"
	"github.com/porky256/dnd-tg-bot/internal/models"
	"strconv"
	"strings"
)

type updateParams struct {
	ChatID             string
	NewAddStatus       models.ChatStatusAdd
	NewStatus          models.ChatStatus
	CharacterName      *string
	CharacterHP        *int
	CharacterLevel     *int
	CharacterArmor     *int
	CharacterAbilities *models.Abilities
	CharacterGold      *int
	CharacterSilver    *int
	CharacterCopper    *int
}

func (h *Handlers) handleDefaultAddName(ctx context.Context, b *bot.Bot, update *bmodels.Update) bot.SendMessageParams {
	answer := bot.SendMessageParams{
		Text:   "Sorry, something went wrong. Try again!",
		ChatID: update.Message.Chat.ID,
	}
	name := update.Message.Text
	if len(name) < 3 || len(name) > 30 {
		answer.Text = "Choose another name. It should be from 3 to 30 characters long"
		return answer
	}
	chatID := strconv.FormatInt(update.Message.Chat.ID, 10)

	err := h.addCharacterUpdate(updateParams{
		ChatID:        chatID,
		NewAddStatus:  models.ChatStatusAddHP,
		NewStatus:     models.ChatStatusAdding,
		CharacterName: &name,
	})

	if err == nil {
		answer.Text = "Now add max HP of your character:"
	}

	return answer
}

func (h *Handlers) handleDefaultAddHP(ctx context.Context, b *bot.Bot, update *bmodels.Update) bot.SendMessageParams {
	answer := bot.SendMessageParams{
		Text:   "Sorry, something went wrong. Try again!",
		ChatID: update.Message.Chat.ID,
	}
	hpStr := update.Message.Text
	hp, err := strconv.Atoi(hpStr)
	if err != nil {
		answer.Text = "Typo in hp. It should be number."

		return answer
	}
	chatID := strconv.FormatInt(update.Message.Chat.ID, 10)
	err = h.addCharacterUpdate(updateParams{
		ChatID:       chatID,
		NewAddStatus: models.ChatStatusAddLevel,
		NewStatus:    models.ChatStatusAdding,
		CharacterHP:  &hp,
	})

	if err == nil {

		answer.Text = "Now add level of your character:"
	}

	return answer
}

func (h *Handlers) handleDefaultAddLevel(ctx context.Context, b *bot.Bot, update *bmodels.Update) bot.SendMessageParams {
	answer := bot.SendMessageParams{
		Text:   "Sorry, something went wrong. Try again!",
		ChatID: update.Message.Chat.ID,
	}
	levelStr := update.Message.Text
	level, err := strconv.Atoi(levelStr)
	if err != nil {
		answer.Text = "Typo in level. It should be number."
		return answer
	}
	chatID := strconv.FormatInt(update.Message.Chat.ID, 10)
	err = h.addCharacterUpdate(updateParams{
		ChatID:         chatID,
		NewAddStatus:   models.ChatStatusAddArmor,
		NewStatus:      models.ChatStatusAdding,
		CharacterLevel: &level,
	})

	if err == nil {
		answer.Text = "Now add armor of your character:"
	}

	return answer
}

func (h *Handlers) handleDefaultAddArmor(ctx context.Context, b *bot.Bot, update *bmodels.Update) bot.SendMessageParams {
	answer := bot.SendMessageParams{
		Text:   "Sorry, something went wrong. Try again!",
		ChatID: update.Message.Chat.ID,
	}
	armorStr := update.Message.Text
	armor, err := strconv.Atoi(armorStr)
	if err != nil {
		answer.Text = "Typo in armor. It should be number."

		return answer
	}
	chatID := strconv.FormatInt(update.Message.Chat.ID, 10)
	err = h.addCharacterUpdate(updateParams{
		ChatID:         chatID,
		NewAddStatus:   models.ChatStatusAddAbilities,
		NewStatus:      models.ChatStatusAdding,
		CharacterArmor: &armor,
	})

	if err == nil {
		answer.Text = "Now add your abilities in format:\n" +
			"STR DEX CON INT WIS CHA\n" +
			"separated by spaces"
	}

	return answer
}

func (h *Handlers) handleDefaultAddAbilities(ctx context.Context, b *bot.Bot, update *bmodels.Update) bot.SendMessageParams {
	answer := bot.SendMessageParams{
		Text:   "Sorry, something went wrong. Try again!",
		ChatID: update.Message.Chat.ID,
	}
	abilitiesStr := update.Message.Text
	abilitiesListStr := strings.Split(abilitiesStr, " ")
	var err error
	var abilities []int
	for _, x := range abilitiesListStr {
		if len(x) > 0 {
			var val int
			val, err = strconv.Atoi(x)
			if err != nil {
				break
			}
			abilities = append(abilities, val)
		}
	}

	if len(abilitiesListStr) != 6 || err != nil {
		answer.Text = "Typo in abilities. They should follow format:\n" +
			"STR DEX CON INT WIS CHA\n" +
			"separated by spaces"

		return answer
	}
	chatID := strconv.FormatInt(update.Message.Chat.ID, 10)

	err = h.addCharacterUpdate(updateParams{
		ChatID:       chatID,
		NewAddStatus: models.ChatStatusAddMoney,
		NewStatus:    models.ChatStatusAdding,
		CharacterAbilities: &models.Abilities{
			Strength:     abilities[0],
			Dexterity:    abilities[1],
			Constitution: abilities[2],
			Intelligence: abilities[3],
			Wisdom:       abilities[4],
			Charisma:     abilities[5],
		},
	})

	if err == nil {
		answer.Text = "Now add your money in format:\n" +
			"Gold Silver Copper\n" +
			"separated by spaces"
	}

	return answer
}

func (h *Handlers) handleDefaultAddMoney(ctx context.Context, b *bot.Bot, update *bmodels.Update) bot.SendMessageParams {
	answer := bot.SendMessageParams{
		Text:   "Sorry, something went wrong. Try again!",
		ChatID: update.Message.Chat.ID,
	}
	moneyStr := update.Message.Text
	moneyListStr := strings.Split(moneyStr, " ")
	var err error
	var money []int
	for _, x := range moneyListStr {
		if len(x) > 0 {
			var val int
			val, err = strconv.Atoi(x)
			if err != nil {
				break
			}
			money = append(money, val)
		}
	}

	if len(moneyListStr) != 3 || err != nil {
		answer.Text = "Typo in coins. They should follow format:\n" +
			"Gold Silver Copper\n" +
			"separated by spaces"

		return answer
	}
	chatID := strconv.FormatInt(update.Message.Chat.ID, 10)

	err = h.addCharacterUpdate(updateParams{
		ChatID:          chatID,
		NewAddStatus:    models.ChatStatusAddDefault,
		NewStatus:       models.ChatStatusDefault,
		CharacterGold:   &money[0],
		CharacterSilver: &money[1],
		CharacterCopper: &money[2],
	})

	if err == nil {
		answer.Text = "Great! your character has been created!"
	}

	return answer
}

func (h *Handlers) addCharacterUpdate(params updateParams) error {
	chat, err := h.chatProvider.GetChatByChatID(params.ChatID)
	if err != nil {
		h.log.Info("error while searching for chat: " + err.Error())
		return err
	}
	char, err := h.charProvider.GetCharacterByID(chat.CurrentCharacterID)
	if err != nil {
		h.log.Info("error while searching for character: " + err.Error())
		return err
	}

	char = updateCharacter(params, char)

	if params.CharacterAbilities != nil {
		err := h.charProvider.UpdateCharactersAbilities(char)
		if err != nil {
			switch err {
			case database.ErrNoRowsUpdated:
				{
					_, err := h.charProvider.InsertCharactersAbilities(char)
					if err != nil {
						h.log.Info("error while inserting abilities: " + err.Error())
						return err
					}
				}
			default:
				{
					h.log.Info("error while updating abilities: " + err.Error())
				}
			}
		}
	}

	err = h.charProvider.UpdateCharacter(char)
	if err != nil {
		h.log.Info("error while updating character: " + err.Error())
		return err
	}
	chat.AddStatus = params.NewAddStatus
	chat.Status = params.NewStatus
	err = h.chatProvider.UpdateChatByChatID(chat)
	if err != nil {
		h.log.Info("error while updating chat: " + err.Error())
		return err
	}
	return nil
}

func updateCharacter(params updateParams, char *models.Character) *models.Character {
	if params.CharacterName != nil {
		char.Name = *params.CharacterName
	}
	if params.CharacterHP != nil {
		char.MaxHP = *params.CharacterHP
		char.CurrentHP = char.MaxHP
	}
	if params.CharacterLevel != nil {
		char.Level = *params.CharacterLevel
	}
	if params.CharacterArmor != nil {
		char.Armor = *params.CharacterArmor
	}
	if params.CharacterAbilities != nil {
		char.Abilities = *params.CharacterAbilities
	}
	if params.CharacterGold != nil {
		char.GoldCoins = *params.CharacterGold
	}
	if params.CharacterSilver != nil {
		char.SilverCoins = *params.CharacterSilver
	}
	if params.CharacterCopper != nil {
		char.CopperCoins = *params.CharacterCopper
	}
	return char
}
