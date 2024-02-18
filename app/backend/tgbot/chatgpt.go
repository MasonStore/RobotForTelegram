package tgbot

import (
	"context"
	"fmt"
	"github.com/MasonStore/RobotForTelegram/app/backend/chatgpt"
	"github.com/MasonStore/RobotForTelegram/app/backend/store/model"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/sashabaranov/go-openai"
	"time"
)

func (b *Bot) chatGPT(message *tgbotapi.Message, text string) {
	api, err := chatgpt.GetChatGPT(b.BotId)
	if err != nil {
		return
	}
	if api == nil {
		return
	}
	dialogId := fmt.Sprintf("%s:%d", message.Chat.Type, message.Chat.ID)
	b.newUserChatMessage(api, dialogId, openai.ChatCompletionMessage{
		Role:    openai.ChatMessageRoleUser,
		Content: text,
	})
	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Second*time.Duration(api.Config.Timeout))
	defer cancelFunc()
	reply, err := api.CreateChatCompletion(ctx, b.buildChatGPTMessage(dialogId, api.Config.SystemPrompt))
	if err != nil {
		b.sendMessage(&model.Reply{
			Model:       &model.Model{},
			RobotId:     b.BotId,
			IsTemplate:  "false",
			MessageType: "text",
			TextContent: (&model.Text{Content: err.Error()}).String(),
			ExpandType:  "none",
		}, message)
		return
	}
	b.newUserChatMessage(api, dialogId, openai.ChatCompletionMessage{
		Role:    openai.ChatMessageRoleAssistant,
		Content: reply,
	})
	b.sendMessage(&model.Reply{
		Model:       &model.Model{},
		RobotId:     b.BotId,
		IsTemplate:  "false",
		MessageType: "text",
		TextContent: (&model.Text{Content: reply}).String(),
		ExpandType:  "none",
	}, message)
}

func (b *Bot) buildChatGPTMessage(
	dialogId string,
	systemPrompt string,
) []openai.ChatCompletionMessage {
	if _, ok := b.dialogMessageMap[dialogId]; !ok {
		b.dialogMessageMap[dialogId] = make([]openai.ChatCompletionMessage, 0)
	}
	res := make([]openai.ChatCompletionMessage, 0)
	//第一个是系统提示
	res = append(res, openai.ChatCompletionMessage{
		Role:    openai.ChatMessageRoleSystem,
		Content: systemPrompt,
	})
	//历史记录
	res = append(res, b.dialogMessageMap[dialogId]...)
	return res
}

func (b *Bot) newUserChatMessage(
	api *chatgpt.ChatGPT,
	dialogId string,
	message openai.ChatCompletionMessage,
) {
	if _, ok := b.dialogMessageMap[dialogId]; !ok {
		b.dialogMessageMap[dialogId] = make([]openai.ChatCompletionMessage, 0)
	}
	b.dialogMessageMap[dialogId] = append(b.dialogMessageMap[dialogId], message)
	// 限制历史记录数量
	historyCount := api.Config.HistoryCount
	// 删除多余的历史记录
	if len(b.dialogMessageMap[dialogId]) > historyCount {
		b.dialogMessageMap[dialogId] = b.dialogMessageMap[dialogId][len(b.dialogMessageMap[dialogId])-historyCount:]
	}
}
