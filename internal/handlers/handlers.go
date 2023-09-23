package handlers

import (
	"context"
	"fmt"
	"github.com/go-telegram/bot"
	bmodels "github.com/go-telegram/bot/models"
	"github.com/porky256/dnd-tg-bot/internal/database"
	"github.com/porky256/dnd-tg-bot/internal/models"
	"github.com/porky256/dnd-tg-bot/internal/repository"
	"github.com/porky256/dnd-tg-bot/internal/repository/postgres"
	"go.uber.org/zap"
	"strconv"
	"strings"
	"time"
)

type Handlers struct {
	userProvider repository.UserProvider
	charProvider repository.CharacterProvider
	chatProvider repository.ChatProvider
	log          *zap.Logger
}

func NewHandlers(db *database.DB, duration time.Duration, log *zap.Logger) *Handlers {
	return &Handlers{
		userProvider: postgres.NewPGUserProvider(db.DB, duration),
		charProvider: postgres.NewPGCharacterProvider(db.DB, duration),
		chatProvider: postgres.NewPGChatProvider(db.DB, duration),
		log:          log,
	}
}

func (h *Handlers) Register(b *bot.Bot) *bot.Bot {
	b.RegisterHandler(bot.HandlerTypeMessageText, "/start", bot.MatchTypeExact,
		h.ApplyMiddlewares(h.HandleStart, h.StopAddingMiddleware))
	b.RegisterHandler(bot.HandlerTypeMessageText, "/stop", bot.MatchTypeExact, h.StopAddingMiddleware)

	b.RegisterHandler(bot.HandlerTypeCallbackQueryData, "get", bot.MatchTypeExact,
		h.ApplyMiddlewares(h.HandleCharCallback, h.StopAddingMiddleware))
	b.RegisterHandler(bot.HandlerTypeCallbackQueryData, "add", bot.MatchTypeExact,
		h.ApplyMiddlewares(h.HandleAddCallback, h.StopAddingMiddleware))
	b.RegisterHandler(bot.HandlerTypeCallbackQueryData, "get-char", bot.MatchTypePrefix,
		h.ApplyMiddlewares(h.HandleGetCharacterCallback, h.StopAddingMiddleware))
	b.RegisterHandler(bot.HandlerTypeCallbackQueryData, "back-from-get-character", bot.MatchTypeExact,
		h.ApplyMiddlewares(h.HandleCharCallback, h.StopAddingMiddleware))
	b.RegisterHandler(bot.HandlerTypeCallbackQueryData, "back-from-get", bot.MatchTypeExact,
		h.ApplyMiddlewares(h.HandleBackFromGetCallback, h.StopAddingMiddleware))
	return b
}

func (h *Handlers) DefaultHandler(ctx context.Context, b *bot.Bot, update *bmodels.Update) {
	if update.Message != nil {
		chat, err := h.chatProvider.GetChatByChatID(strconv.FormatInt(update.Message.Chat.ID, 10))
		if err != nil {
			switch err {
			case database.ErrNoRows:
				{
					params := bot.SendMessageParams{
						ChatID: chat.ChatID,
						Text:   "use command /start first!",
					}
					_, err = b.SendMessage(ctx, &params)
					if err != nil {
						h.log.Info("error have occurred while sending message: " + err.Error())
					}
				}
			default:
				{
					h.log.Info("error have occurred while searching for chat: " + err.Error())
				}
			}
			return
		}

		switch chat.Status {
		case models.ChatStatusAdding:
			{
				var msg bot.SendMessageParams
				switch chat.AddStatus {
				case models.ChatStatusAddName:
					{
						msg = h.handleDefaultAddName(ctx, b, update)
					}
				case models.ChatStatusAddHP:
					{
						msg = h.handleDefaultAddHP(ctx, b, update)
					}
				case models.ChatStatusAddLevel:
					{
						msg = h.handleDefaultAddLevel(ctx, b, update)
					}
				case models.ChatStatusAddArmor:
					{
						msg = h.handleDefaultAddArmor(ctx, b, update)
					}
				case models.ChatStatusAddAbilities:
					{
						msg = h.handleDefaultAddAbilities(ctx, b, update)
					}
				case models.ChatStatusAddMoney:
					{
						msg = h.handleDefaultAddMoney(ctx, b, update)
					}
				}
				_, err := b.SendMessage(ctx, &msg)
				if err != nil {
					h.log.Info("error while sending message " + err.Error())
				}
			}

		}
	}

}

func (h *Handlers) HandleStart(ctx context.Context, b *bot.Bot, update *bmodels.Update) {
	switch update.Message.Chat.Type {
	case "private":
		{

			user := models.User{
				TgID:     strconv.FormatInt(update.Message.From.ID, 10),
				ChatID:   strconv.FormatInt(update.Message.Chat.ID, 10),
				Username: update.Message.From.Username,
			}
			userFromDB, err := h.userProvider.GetUserByTgID(user.TgID)

			if err != nil {
				if err == database.ErrNoRows {
					h.log.Info(err.Error())
					newID, err := h.userProvider.InsertUser(&user)

					if err != nil {
						h.log.Info("error have occurred while inserting user: " + err.Error())
						return
					}
					user.ID = newID
					h.log.Info(fmt.Sprintf("User with username %s registered successfully!", user.Username))

					chat := models.Chat{
						ChatID:       strconv.FormatInt(update.Message.Chat.ID, 10),
						Status:       models.ChatStatusDefault,
						AddStatus:    models.ChatStatusAddDefault,
						UpdateStatus: models.ChatStatusUpdateDefault,
						CreatedAt:    time.Time{},
						UpdatedAt:    time.Time{},
					}
					newID, err = h.chatProvider.InsertChat(&chat)
					if err != nil {
						h.log.Info("error have occurred while inserting chat: " + err.Error())
					}
					h.log.Info(fmt.Sprintf("Chat with username %s registered successfully!", user.Username))
				} else {
					h.log.Info("error have occurred while searching for user: " + err.Error())
					return
				}
			} else {
				user.ID = userFromDB.ID
				if user.Username != userFromDB.Username ||
					user.ChatID != userFromDB.ChatID {
					err = h.userProvider.UpdateUser(&user)
					if err != nil {
						h.log.Info("error have occurred while updating user: " + err.Error())
						return
					}
				}
			}

			param := bot.SendMessageParams{
				ChatID:      user.ChatID,
				Text:        fmt.Sprintf("Hello %s!", user.Username),
				ReplyMarkup: buildBaseInlineKeyboard(),
			}
			_, err = b.SendMessage(ctx, &param)

			if err != nil {
				h.log.Info("error have occurred while sending message: " + err.Error())
			}
		}
	}
}

func (h *Handlers) HandleCharCallback(ctx context.Context, b *bot.Bot, update *bmodels.Update) {
	tgID := strconv.FormatInt(update.CallbackQuery.Sender.ID, 10)
	characters, err := h.charProvider.GetCharactersByOwnerTgID(tgID)
	if err != nil {
		h.log.Info("error have occurred while searching for all users characters: " + err.Error())
	}

	var keyboard [][]bmodels.InlineKeyboardButton

	for _, x := range characters {
		keyboard = append(keyboard, []bmodels.InlineKeyboardButton{{
			Text:         fmt.Sprintf("%s, Lvl:%d", x.Name, x.Level),
			CallbackData: fmt.Sprintf("get-char-%d", x.ID),
		},
		})
	}
	keyboard = append(keyboard, []bmodels.InlineKeyboardButton{{
		Text:         "back",
		CallbackData: "back-from-get",
	}})

	param := bot.SendMessageParams{
		ChatID: update.CallbackQuery.Message.Chat.ID,
		Text:   "Choose character",
		ReplyMarkup: &bmodels.InlineKeyboardMarkup{
			InlineKeyboard: keyboard,
		},
	}

	_, err = b.SendMessage(ctx, &param)
	if err != nil {
		h.log.Info("error have occurred while sending message: " + err.Error())
	}

	_, err = b.DeleteMessage(ctx, &bot.DeleteMessageParams{
		ChatID:    update.CallbackQuery.Message.Chat.ID,
		MessageID: update.CallbackQuery.Message.ID,
	})
	if err != nil {
		h.log.Info("error have occurred while deleting message: " + err.Error())
		return
	}
}

func (h *Handlers) HandleBackFromGetCallback(ctx context.Context, b *bot.Bot, update *bmodels.Update) {
	param := bot.SendMessageParams{
		ChatID:      update.CallbackQuery.Message.Chat.ID,
		Text:        fmt.Sprintf("Hello %s!", update.CallbackQuery.Sender.Username),
		ReplyMarkup: buildBaseInlineKeyboard(),
	}
	_, err := b.SendMessage(ctx, &param)

	if err != nil {
		h.log.Info("error have occurred while sending message: " + err.Error())
	}

	_, err = b.DeleteMessage(ctx, &bot.DeleteMessageParams{
		ChatID:    update.CallbackQuery.Message.Chat.ID,
		MessageID: update.CallbackQuery.Message.ID,
	})
	if err != nil {
		h.log.Info("error have occurred while deleting message: " + err.Error())
		return
	}
}

func (h *Handlers) HandleGetCharacterCallback(ctx context.Context, b *bot.Bot, update *bmodels.Update) {
	idStr, ok := strings.CutPrefix(update.CallbackQuery.Data, "get-char-")
	if !ok {
		h.log.Info("wrong callback data format: " + update.CallbackQuery.Data)
		return
	}
	id, err := strconv.Atoi(idStr)
	if err != nil {
		h.log.Info("wrong callback data format: " + update.CallbackQuery.Data)
		h.log.Info("error have occurred while parsing character's id: " + err.Error())
		return
	}
	char, err := h.charProvider.GetFullCharacterByID(id)
	if err != nil {
		h.log.Info("error have occurred while searching for character: " + err.Error())
		return
	}

	param := bot.SendMessageParams{
		ChatID:    update.CallbackQuery.Message.Chat.ID,
		Text:      models.FormMessage(char),
		ParseMode: bmodels.ParseModeHTML,
		ReplyMarkup: &bmodels.InlineKeyboardMarkup{
			InlineKeyboard: [][]bmodels.InlineKeyboardButton{
				{{Text: "Back",
					CallbackData: "back-from-get-character"}},
			},
		},
	}

	_, err = b.SendMessage(ctx, &param)
	if err != nil {
		h.log.Info("error have occurred while sending message: " + err.Error())
	}

	_, err = b.DeleteMessage(ctx, &bot.DeleteMessageParams{
		ChatID:    update.CallbackQuery.Message.Chat.ID,
		MessageID: update.CallbackQuery.Message.ID,
	})
	if err != nil {
		h.log.Info("error have occurred while deleting message: " + err.Error())
		return
	}
}

func (h *Handlers) HandleAddCallback(ctx context.Context, b *bot.Bot, update *bmodels.Update) {
	emptyCharacter := models.Character{
		OwnerTgID: strconv.FormatInt(update.CallbackQuery.Sender.ID, 10),
	}
	newCharID, err := h.charProvider.InsertCharacter(&emptyCharacter)
	if err != nil {
		h.log.Info("error have occurred while inserting new character: " + err.Error())
	}

	emptyCharacter.ID = newCharID

	_, err = h.charProvider.InsertCharactersAbilities(&emptyCharacter)

	if err != nil {
		h.log.Info("error have occurred while inserting new character's abilities: " + err.Error())
	}

	_, err = h.charProvider.InsertCharactersSkillInsights(&emptyCharacter)

	if err != nil {
		h.log.Info("error have occurred while inserting new character's skill insights: " + err.Error())
	}

	chat := models.Chat{
		ChatID:             strconv.FormatInt(update.CallbackQuery.Message.Chat.ID, 10),
		CurrentCharacterID: newCharID,
		Status:             models.ChatStatusAdding,
		AddStatus:          models.ChatStatusAddName,
	}
	err = h.chatProvider.UpdateChatByChatID(&chat)
	if err != nil {
		h.log.Info("error have occurred while updating chat: " + err.Error())
	}
	params := bot.SendMessageParams{
		Text:   "Please enter a name for your new character (3-30 characters), Use /stop to stop creating character.",
		ChatID: chat.ChatID,
	}
	_, err = b.SendMessage(ctx, &params)
	if err != nil {
		h.log.Info("error have occurred while sending message: " + err.Error())
	}
}

func buildBaseInlineKeyboard() *bmodels.InlineKeyboardMarkup {
	return &bmodels.InlineKeyboardMarkup{
		InlineKeyboard: [][]bmodels.InlineKeyboardButton{
			{{Text: "get characters", CallbackData: "get"},
				{Text: "add character", CallbackData: "add"},
			},
		},
	}
}
