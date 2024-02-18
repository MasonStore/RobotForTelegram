package views

import (
	"fmt"
	"github.com/MasonStore/RobotForTelegram/app/backend/chatgpt"
	"github.com/MasonStore/RobotForTelegram/app/backend/store"
	"github.com/MasonStore/RobotForTelegram/app/backend/store/model"
	"strconv"
)

type HomeView struct {
}

type GetHomeRespKeyboardGroup struct {
	Id   string     `json:"id"`
	Name string     `json:"name"`
	Keys [][]string `json:"keys"`
}
type GetHomeRespButton struct {
	Title string `json:"Title"`
	Url   string `json:"Url"`
}
type GetHomeRespGame struct {
	GameShortName string `json:"GameShortName"`
	GameTitle     string `json:"GameTitle"`
	Url           string `json:"Url"`
	ButtonText    string `json:"ButtonText"`
}
type GetHomeRespPic struct {
	Base64 string `json:"Base64"`
}
type GetHomeRespText struct {
	Content string `json:"Content"`
}
type GetHomeRespReply struct {
	Id                string              `json:"id"`
	Name              string              `json:"name"`
	IsTemplate        string              `json:"isTemplate"`
	MessageType       string              `json:"messageType"`
	GameContent       GetHomeRespGame     `json:"gameContent"`
	PicContent        GetHomeRespPic      `json:"picContent"`
	TextContent       GetHomeRespText     `json:"textContent"`
	ExpandType        string              `json:"expandType"`
	KeyboardGroupId   string              `json:"keyboardGroupId"`
	KeyboardGroupName string              `json:"keyboardGroupName"`
	Buttons           []GetHomeRespButton `json:"buttons"`
}
type GetHomeRespCommand struct {
	Id              string `json:"id"`
	Command         string `json:"command"`
	ReplyId         string `json:"replyId"`
	ReplyName       string `json:"replyName"`
	ReplyIsTemplate string `json:"replyIsTemplate"`
	HookUrl         string `json:"hookUrl"`
}
type GetHomeRespAdvancedConfig struct {
	ChatGptEnabled          bool   `json:"ChatGptEnabled"`
	ChatGptApiKey           string `json:"ChatGptApiKey"`
	ChatGptHttpProxy        string `json:"ChatGptHttpProxy"`
	ChatGptModel            string `json:"ChatGptModel"`
	ChatGptSystemPrompt     string `json:"ChatGptSystemPrompt"`
	ChatGptBaseUrl          string `json:"ChatGptBaseUrl"`
	ChatGptTemperature      string `json:"ChatGptTemperature"`
	ChatGptMaxTokens        string `json:"ChatGptMaxTokens"`
	ChatGptTopP             string `json:"ChatGptTopP"`
	ChatGptPresencePenalty  string `json:"ChatGptPresencePenalty"`
	ChatGptFrequencyPenalty string `json:"ChatGptFrequencyPenalty"`
	ChatGptHistoryCount     string `json:"ChatGptHistoryCount"`
	ChatGptTimeout          string `json:"ChatGptTimeout"`
}
type GetHomeRespRobotInfo struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Username  string `json:"username"`
	Token     string `json:"token"`
}
type GetHomeRespData struct {
	Running        bool                       `json:"running"`
	RobotInfo      GetHomeRespRobotInfo       `json:"robotInfo"`
	KeyboardGroups []GetHomeRespKeyboardGroup `json:"keyboardGroups"`
	Replies        []GetHomeRespReply         `json:"replies"`
	Commands       []GetHomeRespCommand       `json:"commands"`
	AdvancedConfig GetHomeRespAdvancedConfig  `json:"advancedConfig"`
}
type GetHomeResp struct {
	Ok    bool            `json:"ok"`
	Error string          `json:"error"`
	Data  GetHomeRespData `json:"data"`
}

func (v *HomeView) GetHome() *GetHomeResp {
	account, ok := NowSelectedAccount()
	if !ok {
		return &GetHomeResp{
			Ok:    false,
			Error: "NO_SELECTED_ACCOUNT",
		}
	}
	robotId := account.ID
	var keyboardGroupsResp = make([]GetHomeRespKeyboardGroup, 0)
	var repliesResp = make([]GetHomeRespReply, 0)
	var commandsResp = make([]GetHomeRespCommand, 0)
	var advancedConfigResp = &GetHomeRespAdvancedConfig{}
	{
		keyboardGroupsTmp := model.LoadKeyboardGroups(store.Instance.DB, robotId)
		for _, keyboardGroup := range keyboardGroupsTmp {
			keys := make([][]string, 0)
			for _, key := range keyboardGroup.GetKeys() {
				keys = append(keys, key)
			}
			keyboardGroupsResp = append(keyboardGroupsResp, GetHomeRespKeyboardGroup{
				Id:   keyboardGroup.ID,
				Name: keyboardGroup.Name,
				Keys: keys,
			})
		}
	}
	{
		repliesTmp := model.LoadReplies(store.Instance.DB, robotId)
		for _, reply := range repliesTmp {
			var buttons []GetHomeRespButton
			for _, button := range reply.GetButtons() {
				buttons = append(buttons, GetHomeRespButton{
					Title: button.Title,
					Url:   button.Url,
				})
			}
			gameContent := reply.GetGameContent()
			repliesResp = append(repliesResp, GetHomeRespReply{
				Id:          reply.ID,
				Name:        reply.Name,
				IsTemplate:  reply.IsTemplate,
				MessageType: reply.MessageType,
				GameContent: GetHomeRespGame{
					GameShortName: gameContent.GameShortName,
					GameTitle:     gameContent.GameTitle,
					Url:           gameContent.Url,
					ButtonText:    gameContent.ButtonText,
				},
				PicContent: GetHomeRespPic{
					Base64: reply.GetPicContent().Base64,
				},
				TextContent: GetHomeRespText{
					Content: reply.GetTextContent().Content,
				},
				ExpandType:        reply.ExpandType,
				KeyboardGroupId:   reply.KeyboardGroupId,
				KeyboardGroupName: reply.KeyboardGroupName,
				Buttons:           buttons,
			})
		}
	}
	{
		commandsTmp := model.LoadCommands(store.Instance.DB, robotId)
		for _, command := range commandsTmp {
			commandsResp = append(commandsResp, GetHomeRespCommand{
				Id:              command.ID,
				Command:         command.Command,
				ReplyId:         command.ReplyId,
				ReplyName:       command.ReplyName,
				ReplyIsTemplate: command.ReplyIsTemplate,
				HookUrl:         command.HookUrl,
			})
		}
	}
	{
		value, ok := model.GetValueByKey(store.Instance.DB, fmt.Sprintf(model.KeyChatGPTEnabled, robotId))
		if ok {
			advancedConfigResp.ChatGptEnabled = value == "true"
		}
		value, ok = model.GetValueByKey(store.Instance.DB, fmt.Sprintf(model.KeyChatGPTApiKey, robotId))
		if ok {
			advancedConfigResp.ChatGptApiKey = value
		}
		value, ok = model.GetValueByKey(store.Instance.DB, fmt.Sprintf(model.KeyChatGPTHttpProxy, robotId))
		if ok {
			advancedConfigResp.ChatGptHttpProxy = value
		}
		value, ok = model.GetValueByKey(store.Instance.DB, fmt.Sprintf(model.KeyChatGPTModel, robotId))
		if ok {
			advancedConfigResp.ChatGptModel = value
		}
		value, ok = model.GetValueByKey(store.Instance.DB, fmt.Sprintf(model.KeyChatGPTSystemPrompt, robotId))
		if ok {
			advancedConfigResp.ChatGptSystemPrompt = value
		}
		value, ok = model.GetValueByKey(store.Instance.DB, fmt.Sprintf(model.KeyChatGPTBaseUrl, robotId))
		if ok {
			advancedConfigResp.ChatGptBaseUrl = value
		}
		value, ok = model.GetValueByKey(store.Instance.DB, fmt.Sprintf(model.KeyChatGPTTemperature, robotId))
		if ok {
			advancedConfigResp.ChatGptTemperature = value
		}
		value, ok = model.GetValueByKey(store.Instance.DB, fmt.Sprintf(model.KeyChatGPTMaxTokens, robotId))
		if ok {
			advancedConfigResp.ChatGptMaxTokens = value
		}
		value, ok = model.GetValueByKey(store.Instance.DB, fmt.Sprintf(model.KeyChatGPTTopP, robotId))
		if ok {
			advancedConfigResp.ChatGptTopP = value
		}
		value, ok = model.GetValueByKey(store.Instance.DB, fmt.Sprintf(model.KeyChatGPTPresencePenalty, robotId))
		if ok {
			advancedConfigResp.ChatGptPresencePenalty = value
		}
		value, ok = model.GetValueByKey(store.Instance.DB, fmt.Sprintf(model.KeyChatGPTFrequencyPenalty, robotId))
		if ok {
			advancedConfigResp.ChatGptFrequencyPenalty = value
		}
		value, ok = model.GetValueByKey(store.Instance.DB, fmt.Sprintf(model.KeyChatGPTHistoryCount, robotId))
		if ok {
			advancedConfigResp.ChatGptHistoryCount = value
		}
		value, ok = model.GetValueByKey(store.Instance.DB, fmt.Sprintf(model.KeyChatGPTTimeout, robotId))
		if ok {
			advancedConfigResp.ChatGptTimeout = value
		}
	}
	running, _ := model.GetValueByKey(store.Instance.DB, fmt.Sprintf(model.KeyRunning, account.ID))
	return &GetHomeResp{
		Ok:    true,
		Error: "",
		Data: GetHomeRespData{
			Running: running == "true",
			RobotInfo: GetHomeRespRobotInfo{
				FirstName: account.FirstName,
				LastName:  account.LastName,
				Username:  account.UserName,
				Token:     account.Token,
			},
			KeyboardGroups: keyboardGroupsResp,
			Replies:        repliesResp,
			Commands:       commandsResp,
			AdvancedConfig: *advancedConfigResp,
		},
	}
}

func (v *HomeView) SaveHome(req *GetHomeRespData) *GetHomeResp {
	account, ok := NowSelectedAccount()
	if !ok {
		return &GetHomeResp{
			Ok:    false,
			Error: "NO_SELECTED_ACCOUNT",
		}
	}
	robotId := account.ID
	{
		var keyboardGroups []model.KeyboardGroup
		for _, keyboardGroup := range req.KeyboardGroups {
			var keys = model.T2Keys(keyboardGroup.Keys)
			keyboardGroups = append(keyboardGroups, model.KeyboardGroup{
				Model: &model.Model{
					ID: keyboardGroup.Id,
				},
				RobotId: robotId,
				Name:    keyboardGroup.Name,
				Keys:    keys,
			})
		}
		model.SaveKeyboardGroups(store.Instance.DB, robotId, keyboardGroups)
	}
	{
		var replies []model.Reply
		for _, reply := range req.Replies {
			var buttons = make(model.Buttons, 0)
			for _, button := range reply.Buttons {
				buttons = append(buttons, model.Button{
					Title: button.Title,
					Url:   button.Url,
				})
			}
			game := model.Game{
				GameShortName: reply.GameContent.GameShortName,
				GameTitle:     reply.GameContent.GameTitle,
				Url:           reply.GameContent.Url,
				ButtonText:    reply.GameContent.ButtonText,
			}
			pic := model.Pic{
				Base64: reply.PicContent.Base64,
			}
			text := model.Text{
				Content: reply.TextContent.Content,
			}
			replies = append(replies, model.Reply{
				Model: &model.Model{
					ID: reply.Id,
				},
				RobotId:           robotId,
				Name:              reply.Name,
				IsTemplate:        reply.IsTemplate,
				MessageType:       reply.MessageType,
				GameContent:       game.String(),
				PicContent:        pic.String(),
				TextContent:       text.String(),
				ExpandType:        reply.ExpandType,
				KeyboardGroupId:   reply.KeyboardGroupId,
				KeyboardGroupName: reply.KeyboardGroupName,
				Buttons:           buttons.String(),
			})
		}
		model.SaveReplies(store.Instance.DB, robotId, replies)
	}
	{
		var commands []model.Command
		for _, command := range req.Commands {
			commands = append(commands, model.Command{
				Model: &model.Model{
					ID: command.Id,
				},
				RobotId:         robotId,
				Command:         command.Command,
				ReplyId:         command.ReplyId,
				ReplyName:       command.ReplyName,
				ReplyIsTemplate: command.ReplyIsTemplate,
				HookUrl:         command.HookUrl,
			})
		}
		model.SaveCommands(store.Instance.DB, robotId, commands)
	}
	{
		model.SetValueByKey(store.Instance.DB, fmt.Sprintf(model.KeyChatGPTEnabled, robotId), strconv.FormatBool(req.AdvancedConfig.ChatGptEnabled))
		model.SetValueByKey(store.Instance.DB, fmt.Sprintf(model.KeyChatGPTApiKey, robotId), req.AdvancedConfig.ChatGptApiKey)
		model.SetValueByKey(store.Instance.DB, fmt.Sprintf(model.KeyChatGPTHttpProxy, robotId), req.AdvancedConfig.ChatGptHttpProxy)
		model.SetValueByKey(store.Instance.DB, fmt.Sprintf(model.KeyChatGPTModel, robotId), req.AdvancedConfig.ChatGptModel)
		model.SetValueByKey(store.Instance.DB, fmt.Sprintf(model.KeyChatGPTSystemPrompt, robotId), req.AdvancedConfig.ChatGptSystemPrompt)
		model.SetValueByKey(store.Instance.DB, fmt.Sprintf(model.KeyChatGPTBaseUrl, robotId), req.AdvancedConfig.ChatGptBaseUrl)
		model.SetValueByKey(store.Instance.DB, fmt.Sprintf(model.KeyChatGPTTemperature, robotId), req.AdvancedConfig.ChatGptTemperature)
		model.SetValueByKey(store.Instance.DB, fmt.Sprintf(model.KeyChatGPTMaxTokens, robotId), req.AdvancedConfig.ChatGptMaxTokens)
		model.SetValueByKey(store.Instance.DB, fmt.Sprintf(model.KeyChatGPTTopP, robotId), req.AdvancedConfig.ChatGptTopP)
		model.SetValueByKey(store.Instance.DB, fmt.Sprintf(model.KeyChatGPTPresencePenalty, robotId), req.AdvancedConfig.ChatGptPresencePenalty)
		model.SetValueByKey(store.Instance.DB, fmt.Sprintf(model.KeyChatGPTFrequencyPenalty, robotId), req.AdvancedConfig.ChatGptFrequencyPenalty)
		model.SetValueByKey(store.Instance.DB, fmt.Sprintf(model.KeyChatGPTHistoryCount, robotId), req.AdvancedConfig.ChatGptHistoryCount)
		model.SetValueByKey(store.Instance.DB, fmt.Sprintf(model.KeyChatGPTTimeout, robotId), req.AdvancedConfig.ChatGptTimeout)
	}
	chatgpt.ReloadChatGPT(robotId)
	return v.GetHome()
}

func (v *HomeView) SetRunning(running bool) *GetHomeResp {
	account, ok := NowSelectedAccount()
	if !ok {
		return &GetHomeResp{
			Ok:    false,
			Error: "NO_SELECTED_ACCOUNT",
		}
	}
	robotId := account.ID
	if running {
		model.SetValueByKey(store.Instance.DB, fmt.Sprintf(model.KeyRunning, robotId), "true")
	} else {
		model.SetValueByKey(store.Instance.DB, fmt.Sprintf(model.KeyRunning, robotId), "false")
	}
	return v.GetHome()
}
