package tgbot

import (
	"bytes"
	"encoding/base64"
	"github.com/MasonStore/RobotForTelegram/app/backend/store"
	"github.com/MasonStore/RobotForTelegram/app/backend/store/model"
	"github.com/MasonStore/RobotForTelegram/app/backend/utils"
	. "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"image"
	"image/draw"
	_ "image/gif"
	"image/jpeg"
	_ "image/jpeg"
	_ "image/png"
	"strings"
)

func base64ToBytes(str string) ([]byte, error) {
	//把图片的base64转成[]byte，前面的data:image/png;base64,去掉
	imgBase64 := strings.Replace(str, "data:image/png;base64,", "", 1)
	imgBase64 = strings.Replace(imgBase64, "data:image/jpeg;base64,", "", 1)
	imgBase64 = strings.Replace(imgBase64, "data:image/jpg;base64,", "", 1)
	imgBase64 = strings.Replace(imgBase64, "data:image/gif;base64,", "", 1)
	imgBase64 = strings.Replace(imgBase64, "data:image/webp;base64,", "", 1)
	imgBase64 = strings.Replace(imgBase64, "data:image/svg+xml;base64,", "", 1)
	imgBase64 = strings.Replace(imgBase64, "data:image/bmp;base64,", "", 1)
	imgBase64 = strings.Replace(imgBase64, "data:image/x-icon;base64,", "", 1)
	//解码
	imgBytes, err := base64.StdEncoding.DecodeString(imgBase64)
	if err != nil {
		return nil, err
	}
	return imgBytes, nil
}

// 图片缩略操作
func PicThumb(picBytes []byte) ([]byte, error) {
	// 先转换成go image.Image
	img, _, err := image.Decode(bytes.NewReader(picBytes))
	if err != nil {
		return nil, err
	}
	// 缩略图
	thumb := image.NewRGBA(image.Rect(0, 0, 100, 100))
	// 上面这行表示创建一个100*100的缩略图
	// 缩略图绘制
	draw.Draw(thumb, thumb.Bounds(), img, img.Bounds().Min, draw.Src)
	// 缩略图转换成[]byte
	buf := new(bytes.Buffer)
	err = jpeg.Encode(buf, thumb, nil)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func (b *Bot) createPicMessage(messageModel *model.Reply, message *tgbotapi.Message) tgbotapi.Chattable {
	contentPhoto := messageModel.GetPicContent()
	imageBytes, err := base64ToBytes(contentPhoto.Base64)
	if err != nil {
		utils.LogPrintf("base64 to bytes error: %v\n", err)
		return nil
	}
	file := FileBytes{
		Name:  "image.png",
		Bytes: imageBytes,
	}
	// 缩略图
	thumb, err := PicThumb(imageBytes)
	if err != nil {
		utils.LogPrintf("pic thumb error: %v\n", err)
		//return nil
		thumb = imageBytes
	}
	fileThumb := FileBytes{
		Name:  "thumb.png",
		Bytes: thumb,
	}
	tmp := PhotoConfig{
		BaseFile: BaseFile{
			BaseChat: BaseChat{
				ChatID:           message.Chat.ID,
				ReplyToMessageID: message.MessageID,
				ReplyMarkup:      nil,
			},
			File: file,
		},
		Thumb:           fileThumb,
		Caption:         "",
		ParseMode:       "",
		CaptionEntities: nil,
	}
	b.setupPhotoReplyMarkup(contentPhoto, &tmp, messageModel)
	return &tmp
}

func (b *Bot) setupPhotoReplyMarkup(photoContent *model.Pic, m *tgbotapi.PhotoConfig, messageModel *model.Reply) {
	db := store.Instance.DB
	switch messageModel.ExpandType {
	case "keyboardGroup":
		if messageModel.KeyboardGroupId == "" {
			// 去除键盘
			m.ReplyMarkup = NewRemoveKeyboard(false)
		} else {
			keyboard := &model.KeyboardGroup{}
			err := db.Model(keyboard).Where("id = ?", messageModel.KeyboardGroupId).First(keyboard).Error
			if err != nil {
				utils.LogPrintf("get keyboard error: %v\n", err)
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
		if messageModel.Buttons == "" {
			m.ReplyMarkup = nil
		} else {
			buttons := messageModel.GetButtons()
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
