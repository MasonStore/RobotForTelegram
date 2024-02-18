package tgbot

import (
	"fmt"
	"github.com/MasonStore/RobotForTelegram/app/backend/store/model"
	"github.com/MasonStore/RobotForTelegram/app/backend/utils"
	. "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (b *Bot) createGameMessage(messageModel *model.Reply, message *tgbotapi.Message) tgbotapi.Chattable {
	contentGame := messageModel.GetGameContent()
	tmp := GameConfig{
		BaseChat: BaseChat{
			ChatID:           message.Chat.ID,
			ReplyToMessageID: message.MessageID,
			ReplyMarkup:      nil,
		},
		GameShortName: contentGame.GameShortName,
	}
	b.setupGameReplyMarkup(contentGame, &tmp, messageModel)
	return &tmp
}

func (b *Bot) setupGameReplyMarkup(contentGame *model.Game, m *GameConfig, reply *model.Reply) {
	replyMarkup := make([][]InlineKeyboardButton, 0)
	// 第一个必须是游戏
	replyMarkup = append(replyMarkup, []InlineKeyboardButton{{
		Text:                         contentGame.ButtonText,
		URL:                          nil,
		LoginURL:                     nil,
		CallbackData:                 nil,
		SwitchInlineQuery:            nil,
		SwitchInlineQueryCurrentChat: nil,
		CallbackGame:                 &CallbackGame{},
		Pay:                          false,
	}})
	b.registerGameHandler(contentGame, func(query *tgbotapi.CallbackQuery, update tgbotapi.Update) {
		f := "%s?from=telegram&game=%s&userId=%d&userName=%s&userAvatar=%s"
		url := fmt.Sprintf(f, contentGame.Url, contentGame.GameShortName, query.From.ID, query.From.UserName, query.From.FirstName)
		callback := tgbotapi.NewUrlCallback(query.ID, url)
		if _, err := b.API.Request(callback); err != nil {
			panic(err)
		}
		utils.LogPrintf("url: %s\n", url)
	})
	switch reply.ExpandType {
	case "keyboardGroup":
		// 不支持 走下一个case
	case "button":
		buttons := reply.GetButtons()
		inlineButtons := make([][]InlineKeyboardButton, 0)
		for _, button := range buttons {
			url := button.Url
			inlineButtons = append(inlineButtons, []InlineKeyboardButton{{
				Text: button.Title,
				URL:  &url,
			}})
		}
		replyMarkup = append(replyMarkup, inlineButtons...)
	}
	m.ReplyMarkup = NewInlineKeyboardMarkup(replyMarkup...)
}

func (b *Bot) registerGameHandler(game *model.Game, f func(query *tgbotapi.CallbackQuery, update tgbotapi.Update)) {
	b.gameHandlersLock.Lock()
	defer b.gameHandlersLock.Unlock()
	b.gameHandlers[game.GameTitle] = f
}

func (b *Bot) handleGame(update tgbotapi.Update, game *tgbotapi.Game) {
	title := game.Title
	b.gameHandlersLock.RLock()
	defer b.gameHandlersLock.RUnlock()
	f, ok := b.gameHandlers[title]
	if !ok {
		return
	}
	f(update.CallbackQuery, update)
}
