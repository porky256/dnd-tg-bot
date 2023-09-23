package handlers

import (
	"context"
	"github.com/go-telegram/bot"
	bmodels "github.com/go-telegram/bot/models"
	"github.com/porky256/dnd-tg-bot/internal/models"
	"strconv"
)

func (h *Handlers) StopAddingMiddleware(ctx context.Context, b *bot.Bot, update *bmodels.Update) {
	var chatID string
	if update.Message != nil {
		chatID = strconv.FormatInt(update.Message.Chat.ID, 10)
	} else if update.CallbackQuery != nil {
		chatID = strconv.FormatInt(update.CallbackQuery.Message.Chat.ID, 10)
	} else {
		return
	}
	chat, err := h.chatProvider.GetChatByChatID(chatID)
	if err != nil {
		h.log.Info("error have occurred while searching for chat: " + err.Error())
		return
	}
	if chat.Status == models.ChatStatusAdding {
		err := h.charProvider.DeleteCharacterByID(chat.CurrentCharacterID)
		if err != nil {
			h.log.Info("error have occurred while deleting character: " + err.Error())
			return
		}
		chat.Status = models.ChatStatusDefault
		chat.AddStatus = models.ChatStatusAddDefault
		chat.CurrentCharacterID = 0
		err = h.chatProvider.UpdateChatByChatID(chat)
		if err != nil {
			h.log.Info("error have occurred while updating chat: " + err.Error())
		}
	}
}

func (h *Handlers) ApplyMiddlewares(handler bot.HandlerFunc, middlewares ...bot.HandlerFunc) bot.HandlerFunc {

	return func(ctx context.Context, b *bot.Bot, update *bmodels.Update) {
		for _, x := range middlewares {
			x(ctx, b, update)
		}
		handler(ctx, b, update)
	}
}
