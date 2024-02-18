package tgbot

import (
	"fmt"
	"github.com/MasonStore/RobotForTelegram/app/backend/chatgpt"
	"github.com/MasonStore/RobotForTelegram/app/backend/store"
	"github.com/MasonStore/RobotForTelegram/app/backend/store/model"
	"github.com/MasonStore/RobotForTelegram/app/backend/utils"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/sashabaranov/go-openai"
	_ "github.com/sashabaranov/go-openai"
	"sync"
)

type Bot struct {
	BotId uint
	API   *tgbotapi.BotAPI

	gameHandlers     map[string]func(query *tgbotapi.CallbackQuery, update tgbotapi.Update)
	gameHandlersLock sync.RWMutex
	releaseChan      chan struct{}

	dialogMessageMap map[string][]openai.ChatCompletionMessage
}

var BotMap = make(map[uint]*Bot)

func NewBot(botId uint) (*Bot, error) {
	account, ok := model.GetAccountById(store.Instance.DB, botId)
	if !ok {
		return nil, ErrNoAccount
	}
	bot, err := tgbotapi.NewBotAPI(account.Token)
	if err != nil {
		return nil, err
	}
	b := &Bot{
		BotId:            botId,
		API:              bot,
		gameHandlers:     make(map[string]func(query *tgbotapi.CallbackQuery, update tgbotapi.Update)),
		releaseChan:      make(chan struct{}),
		dialogMessageMap: make(map[string][]openai.ChatCompletionMessage),
	}
	BotMap[botId] = b
	b.Start()
	return b, nil
}

func CheckRobot() {
	accounts := make([]*model.Account, 0)
	db := store.Instance.DB
	db.Model(&model.Account{}).Find(&accounts)
	for _, account := range accounts {
		chatgpt.ReloadChatGPT(account.ID)
		if _, ok := BotMap[account.ID]; !ok {
			NewBot(account.ID)
		}
	}
	for id, b := range BotMap {
		// 是否存在于数据库
		exist := false
		for _, account := range accounts {
			if account.ID == id {
				exist = true
				break
			}
		}
		if !exist {
			//释放
			b.Release()
			delete(BotMap, id)
		}
	}
}

func safeGo(f func()) {
	go func() {
		defer func() {
			if err := recover(); err != nil {
				utils.LogPrintf("panic: %v\n", err)
			}
		}()
		f()
	}()
}

func (b *Bot) Start() {
	go func() {
		u := tgbotapi.NewUpdate(0)
		u.Timeout = 60
		updates := b.API.GetUpdatesChan(u)
		for {
			select {
			case <-b.releaseChan:
				return
			case update := <-updates:
				// 运行状态设置
				running, _ := model.GetValueByKey(store.Instance.DB, fmt.Sprintf(model.KeyRunning, b.BotId))
				if running != "true" {
					continue
				}
				if update.Message != nil {
					if update.Message.IsCommand() {
						// command
						safeGo(func() {
							b.handleCommand(update.Message)
						})
					} else {
						text := update.Message.Text
						if text != "" {
							// command
							safeGo(func() {
								b.handleCommand(update.Message)
							})
						}
					}
				}

				if update.CallbackQuery != nil {
					message := update.CallbackQuery.Message
					if message != nil {
						game := message.Game
						if game != nil {
							// game
							safeGo(func() {
								b.handleGame(update, game)
							})
						}
					}
				}
			}
		}
	}()
}

func (b *Bot) Release() {
	b.releaseChan <- struct{}{}
}
