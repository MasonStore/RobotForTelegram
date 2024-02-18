package tgbot

import (
	"github.com/MasonStore/RobotForTelegram/app/backend/store"
	"github.com/MasonStore/RobotForTelegram/app/backend/store/model"
	. "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

func (b *Bot) createTextMessage(messageModel *model.Reply, message *tgbotapi.Message) tgbotapi.Chattable {
	contentText := messageModel.GetTextContent()
	tmp := MessageConfig{
		BaseChat: BaseChat{
			ChatID:           message.Chat.ID,
			ReplyToMessageID: message.MessageID,
			ReplyMarkup:      nil,
		},
		Text:                  contentText.Content,
		DisableWebPagePreview: false,
	}
	b.setupTextReplyMarkup(&tmp, messageModel)
	return &tmp
}

func (b *Bot) setupTextReplyMarkup(m *MessageConfig, reply *model.Reply) {
	db := store.Instance.DB
	switch reply.ExpandType {
	case "keyboardGroup":
		if reply.KeyboardGroupId == "" {
			// 去除键盘
			m.ReplyMarkup = NewRemoveKeyboard(false)
		} else {
			keyboard := &model.KeyboardGroup{}
			err := db.Model(keyboard).Where("id = ?", reply.KeyboardGroupId).First(keyboard).Error
			if err != nil {
				log.Printf("get keyboard error: %v", err)
				return
			}
			keyboardButtons := make([][]KeyboardButton, 0)
			for _, row := range keyboard.GetKeys() {
				rowButtons := make([]KeyboardButton, 0)
				for _, button := range row {
					rowButtons = append(rowButtons, KeyboardButton{
						Text: button,
					})
				}
				keyboardButtons = append(keyboardButtons, rowButtons)
			}
			m.ReplyMarkup = NewReplyKeyboard(keyboardButtons...)
		}
	case "button":
		if reply.Buttons == "" {
			m.ReplyMarkup = nil
		} else {
			buttons := reply.GetButtons()
			inlineButtons := make([][]InlineKeyboardButton, 0)
			for _, button := range buttons {
				url := button.Url
				inlineButtons = append(inlineButtons, []InlineKeyboardButton{{
					Text: button.Title,
					URL:  &url,
				}})
			}
			m.ReplyMarkup = NewInlineKeyboardMarkup(inlineButtons...)
		}
	default:
		return
	}
}
