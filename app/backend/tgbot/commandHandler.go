package tgbot

import (
	"bytes"
	"github.com/MasonStore/RobotForTelegram/app/backend/store"
	"github.com/MasonStore/RobotForTelegram/app/backend/store/model"
	"github.com/MasonStore/RobotForTelegram/app/backend/utils"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"strings"
	"text/template"
)

func (b *Bot) handleCommand(message *tgbotapi.Message) {
	command := message.Command()
	if command == "" {
		command = message.Text
	}
	db := store.Instance.DB
	commands := make([]*model.Command, 0)
	db.Model(&model.Command{}).
		Where("command = ? OR command = ?",
			command, "/"+command,
		).Find(&commands)
	if len(commands) == 0 {
		b.chatGPT(message, command)
		return
	}
	for _, c := range commands {
		c := c
		if c.ID != "" {
			// handle command
			messageModel := &model.Reply{}
			err := db.Model(&model.Reply{}).Where("id = ?", c.ReplyId).First(messageModel).Error
			if err != nil {
				continue
			}
			b.handleCommandMessage(c, messageModel, message)
		}
	}
}

func (b *Bot) handleCommandMessage(commandModel *model.Command, messageModel *model.Reply, message *tgbotapi.Message) {
	// 先hook
	var resp *map[string]any
	if commandModel.HookUrl != "" {
		tmp, err := b.hook(commandModel.HookUrl, message)
		if err != nil {
			utils.LogPrintf("hook error: %v\n", err)
		} else {
			if tmp != nil {
				// log
				utils.LogPrintf("hook resp: %v\n", tmp)
				resp = &tmp
			}
		}
	}
	if messageModel.IsTemplate == "true" {
		if resp == nil {
			utils.LogPrintf("resp is nil\n")
			return
		}
		newResp := make(map[string]string)
		for k, v := range *resp {
			if vStr, ok := v.(string); ok {
				// vStr 可能有\n，改成<br>
				vStr = strings.ReplaceAll(vStr, "\n", `\n`)
				vStr = strings.ReplaceAll(vStr, "\r", `\r`)
				newResp[k] = vStr
			}
		}
		if len(newResp) == 0 {
			utils.LogPrintf("new resp is empty\n")
			return
		}
		contentTmpl := messageModel.GetContentTmpl()
		// apply tmpl
		parse, err := template.New("tmpl").Parse(contentTmpl)
		if err != nil {
			utils.LogPrintf("parse template error: %v\n", err)
			return
		}
		var text = bytes.NewBuffer([]byte{})
		err = parse.Execute(text, newResp)
		if err != nil {
			utils.LogPrintf("execute template error: %v\n", err)
			return
		}
		messageModel.IsTemplate = "false"
		messageModel.SetUpFromContent(text.String())
		b.sendMessage(
			messageModel,
			message,
		)
	} else {
		b.sendMessage(
			messageModel,
			message,
		)
	}
}
