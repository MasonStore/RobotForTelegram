package tgbot

import (
	"github.com/MasonStore/RobotForTelegram/app/backend/store/model"
	. "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

func (b *Bot) sendMessage(messageModel *model.Reply, message *Message) {
	var chattable Chattable
	switch messageModel.MessageType {
	case "text":
		chattable = b.createTextMessage(messageModel, message)
	case "game":
		chattable = b.createGameMessage(messageModel, message)
	case "pic":
		chattable = b.createPicMessage(messageModel, message)
	}
	if chattable != nil {
		resp, err := b.API.Send(chattable)
		if err != nil {
			log.Printf("send message error: %v", err)
		} else {
			log.Printf("send message resp: %v", resp)
		}
	}
}
