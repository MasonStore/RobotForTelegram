package views

import (
	"github.com/MasonStore/RobotForTelegram/app/backend/utils"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"testing"
)

func TestTgBot(t *testing.T) {
	botAPI, err := tgbotapi.NewBotAPI("")
	if err != nil {
		t.Fatalf("NewBotAPI error: %s", err)
	}
	me, err := botAPI.GetMe()
	if err != nil {
		t.Fatalf("GetMe error: %s", err)
	}
	photos, err := botAPI.GetUserProfilePhotos(tgbotapi.NewUserProfilePhotos(me.ID))
	if err != nil {
		t.Fatalf("GetUserProfilePhotos error: %s", err)
	}
	if len(photos.Photos) != 0 {
		photo := photos.Photos[0][0]
		photoFile, err := botAPI.GetFileDirectURL(photo.FileID)
		if err == nil {
			utils.LogPrintf("photoFile: %s\n", photoFile)
		}
	}
	//botAPI.SetAPIEndpoint("https://api.telegram.org/bot")
}
