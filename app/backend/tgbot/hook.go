package tgbot

import (
	"encoding/json"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"io"
	"log"
	"net/http"
	"strings"
)

type HookRequestFrom struct {
	ID        int64  `json:"id"`
	IsBot     bool   `json:"isBot"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	UserName  string `json:"userName"`
}
type HookRequest struct {
	MessageID int             `json:"messageId"`
	From      HookRequestFrom `json:"from"`
	Date      int             `json:"date"` // 10位时间戳
	Text      string          `json:"text"`
}

func (b *Bot) hook(url string, message *tgbotapi.Message) (resp map[string]any, err error) {
	req := &HookRequest{
		MessageID: message.MessageID,
		From: HookRequestFrom{
			ID:        message.From.ID,
			IsBot:     message.From.IsBot,
			FirstName: message.From.FirstName,
			LastName:  message.From.LastName,
			UserName:  message.From.UserName,
		},
		Date: message.Date,
		Text: message.Text,
	}
	log.Printf("HookUrl: %s, req: %v", url, req)
	reqBytes, _ := json.Marshal(req)
	body := strings.NewReader(string(reqBytes))
	request, err := http.NewRequest("POST", url, body)
	if err != nil {
		log.Printf("new request error: %v", err)
		return
	}
	request.Header.Set("Content-Type", "application/json")
	client := http.DefaultClient
	response, err := client.Do(request)
	if err != nil {
		log.Printf("do request error: %v", err)
		return
	}
	if response.StatusCode != 200 {
		log.Printf("response status code: %d", response.StatusCode)
		return
	}
	resp = make(map[string]any)
	bytes, err := io.ReadAll(response.Body)
	if err != nil {
		log.Printf("read response body error: %v", err)
		return
	}
	err = json.Unmarshal(bytes, &resp)
	if err != nil {
		log.Printf("unmarshal response body error: %v", err)
		return
	}
	return
}
