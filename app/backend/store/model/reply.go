package model

import (
	"encoding/json"
	"gorm.io/gorm"
)

type Game struct {
	GameShortName string `json:"GameShortName"`
	GameTitle     string `json:"GameTitle"`
	Url           string `json:"Url"`
	ButtonText    string `json:"ButtonText"`
}
type Pic struct {
	Base64 string `json:"Base64"`
}
type Text struct {
	Content string `json:"Content"`
}
type Button struct {
	Title string `json:"Title"`
	Url   string `json:"Url"`
}
type Buttons []Button

type Reply struct {
	*Model
	RobotId           uint   `json:"robotId" gorm:"column:robot_id;type:integer;not null"`
	Name              string `json:"name" gorm:"column:name;type:text;not null"`
	IsTemplate        string `json:"isTemplate" gorm:"column:is_template;type:text;not null"`
	MessageType       string `json:"messageType" gorm:"column:message_type;type:text;not null"`
	GameContent       string `json:"gameContent" gorm:"column:game_content;type:text;not null"`
	PicContent        string `json:"picContent" gorm:"column:pic_content;type:text;not null"`
	TextContent       string `json:"textContent" gorm:"column:text_content;type:text;not null"`
	ExpandType        string `json:"expandType" gorm:"column:expand_type;type:text;not null"`
	KeyboardGroupId   string `json:"keyboardGroupId" gorm:"column:keyboard_group_id;type:integer;not null"`
	KeyboardGroupName string `json:"keyboardGroupName" gorm:"column:keyboard_group_name;type:text;not null"`
	Buttons           string `json:"buttons" gorm:"column:buttons;type:text;not null"`
}

func (m *Reply) TableName() string {
	return "reply"
}

func SaveReplies(tx *gorm.DB, robotId uint, replies []Reply) {
	oldReplies := LoadReplies(tx, robotId)
	deleteReplies := make([]Reply, 0)
	//如果oldReplies中的命令不在replies中，表示需要删除
	for _, oldReply := range oldReplies {
		found := false
		for _, reply := range replies {
			if oldReply.ID == reply.ID {
				found = true
				break
			}
		}
		if !found {
			deleteReplies = append(deleteReplies, oldReply)
		}
	}
	_ = tx.Transaction(func(tx *gorm.DB) error {
		for _, reply := range replies {
			tx.Save(&reply)
		}
		for _, reply := range deleteReplies {
			tx.Unscoped().Delete(&reply)
		}
		return nil
	})
	return
}

func LoadReplies(tx *gorm.DB, robotId uint) []Reply {
	var replies []Reply
	tx.Where("robot_id = ?", robotId).Order("id ASC").Find(&replies)
	return replies
}

func (m *Reply) GetGameContent() *Game {
	var gameContent Game
	_ = json.Unmarshal([]byte(m.GameContent), &gameContent)
	return &gameContent
}

func (m *Reply) GetPicContent() *Pic {
	var picContent Pic
	_ = json.Unmarshal([]byte(m.PicContent), &picContent)
	return &picContent
}

func (m *Reply) GetTextContent() *Text {
	var textContent Text
	_ = json.Unmarshal([]byte(m.TextContent), &textContent)
	return &textContent
}

func (m *Reply) GetButtons() Buttons {
	var buttons Buttons
	_ = json.Unmarshal([]byte(m.Buttons), &buttons)
	return buttons
}

func (m *Reply) GetContentTmpl() string {
	switch m.MessageType {
	case "game":
		return `{"GameShortName":"{{.GameShortName}}","GameTitle":"{{.GameTitle}}","Url":"{{.Url}}","ButtonText":"{{.ButtonText}}"}`
	case "pic":
		return `{"Base64":"{{.Base64}}"}`
	case "text":
		return `{"Content":"{{.Content}}"}`
	default:
		return ""
	}
}

func (m *Reply) SetUpFromContent(str string) {
	switch m.MessageType {
	case "game":
		var gameContent Game
		_ = json.Unmarshal([]byte(str), &gameContent)
		m.GameContent = gameContent.String()
	case "pic":
		var picContent Pic
		_ = json.Unmarshal([]byte(str), &picContent)
		m.PicContent = picContent.String()
	case "text":
		var textContent Text
		_ = json.Unmarshal([]byte(str), &textContent)
		m.TextContent = textContent.String()
	}
}

func (t *Game) String() string {
	b, _ := json.Marshal(t)
	return string(b)
}

func (t *Pic) String() string {
	b, _ := json.Marshal(t)
	return string(b)
}

func (t *Text) String() string {
	b, _ := json.Marshal(t)
	return string(b)
}

func (t *Buttons) String() string {
	b, _ := json.Marshal(t)
	return string(b)
}
