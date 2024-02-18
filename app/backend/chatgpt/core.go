package chatgpt

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/MasonStore/RobotForTelegram/app/backend/store"
	"github.com/MasonStore/RobotForTelegram/app/backend/store/model"
	"github.com/MasonStore/RobotForTelegram/app/backend/utils"
	"github.com/sashabaranov/go-openai"
	"net/http"
	"net/url"
	"strconv"
	"sync"
)

type Config struct {
	Token            string
	HttpProxy        string
	BaseUrl          string
	Model            string
	SystemPrompt     string
	Temperature      float64 // 0.0 - 1.0 default 0.5 随机性 值越大越随机
	MaxTokens        int     // 1 - 2048 default 2000 最大生成长度
	TopP             float64 // 0.0 - 1.0 default 1.0 选择概率
	PresencePenalty  float64 // 0.0 - 1.0 default 0.0 重复惩罚
	FrequencyPenalty float64 // 0.0 - 1.0 default 0.0 词频惩罚
	HistoryCount     int     // 0 - 20 default 4 历史记录数量
	Timeout          int     // 0 - 600 default 10 超时时间
}

func (c Config) String() string {
	b, _ := json.Marshal(c)
	return string(b)
}

func DefaultConfig() Config {
	return Config{
		Token:            "",
		HttpProxy:        "",
		BaseUrl:          "",
		Model:            openai.GPT3Dot5Turbo,
		SystemPrompt:     "You are a friendly AI created by OpenAI. You are very helpful and will answer any questions you ask.",
		Temperature:      0.5,
		MaxTokens:        2000,
		TopP:             1.0,
		PresencePenalty:  0,
		FrequencyPenalty: 0,
		HistoryCount:     4,
		Timeout:          60,
	}
}

type ChatGPT struct {
	BotId  uint
	Config Config
	API    *openai.Client
}

func MarshalAnyToString(v any) string {
	b, _ := json.Marshal(v)
	return string(b)
}

func (g *ChatGPT) CreateChatCompletion(ctx context.Context, messages []openai.ChatCompletionMessage) (string, error) {
	client := g.API
	request := openai.ChatCompletionRequest{
		Model:            g.Config.Model,
		Messages:         messages,
		MaxTokens:        g.Config.MaxTokens,
		Temperature:      float32(g.Config.Temperature),
		TopP:             float32(g.Config.TopP),
		PresencePenalty:  float32(g.Config.PresencePenalty),
		FrequencyPenalty: float32(g.Config.FrequencyPenalty),
	}
	utils.LogPrintf("openAIClient.CreateChatCompletion request=%s\n", MarshalAnyToString(request))
	resp, err := client.CreateChatCompletion(ctx, request)
	if err != nil {
		utils.LogPrintf("openAIClient.CreateChatCompletion err=%+v\n"+
			"config: %s\n", err, g.Config.String())
		return "", err
	}
	if len(resp.Choices) == 0 {
		utils.LogPrintf("resp is err, resp=%s\n", MarshalAnyToString(resp))
		return "", errors.New("resp is err")
	}
	return resp.Choices[0].Message.Content, nil
}

var chatGPTBotMap = make(map[uint]*ChatGPT)
var chatGPTBotMapLock sync.Mutex

func GetChatGPT(botId uint) (*ChatGPT, error) {
	chatGPTBotMapLock.Lock()
	api, ok := chatGPTBotMap[botId]
	chatGPTBotMapLock.Unlock()
	if ok {
		return api, nil
	}
	return ReloadChatGPT(botId)
}

func ReloadChatGPT(
	botId uint,
) (*ChatGPT, error) {
	chatGPTBotMapLock.Lock()
	defer chatGPTBotMapLock.Unlock()
	delete(chatGPTBotMap, botId)
	config := DefaultConfig()
	db := store.Instance.DB
	enabled, _ := model.GetValueByKey(db, fmt.Sprintf(model.KeyChatGPTEnabled, botId))
	if enabled != "true" {
		return nil, errors.New("ChatGPT is not enabled")
	}
	apiKey, _ := model.GetValueByKey(db, fmt.Sprintf(model.KeyChatGPTApiKey, botId))
	if apiKey == "" {
		return nil, errors.New("ChatGPT apiKey is empty")
	}
	config.Token = apiKey
	defaultConfig := openai.DefaultConfig(config.Token)
	config.HttpProxy, _ = model.GetValueByKey(db, fmt.Sprintf(model.KeyChatGPTHttpProxy, botId))
	config.Model, _ = model.GetValueByKey(db, fmt.Sprintf(model.KeyChatGPTModel, botId))
	if config.Model == "" {
		config.Model = openai.GPT3Dot5Turbo
	}
	config.SystemPrompt, _ = model.GetValueByKey(db, fmt.Sprintf(model.KeyChatGPTSystemPrompt, botId))
	if config.SystemPrompt == "" {
		config.SystemPrompt = "You are a friendly AI created by OpenAI. You are very helpful and will answer any questions you ask."
	}
	config.BaseUrl, _ = model.GetValueByKey(db, fmt.Sprintf(model.KeyChatGPTBaseUrl, botId))
	temperature, _ := model.GetValueByKey(db, fmt.Sprintf(model.KeyChatGPTTemperature, botId))
	if temperature != "" {
		config.Temperature, _ = strconv.ParseFloat(temperature, 64)
	}
	maxTokens, _ := model.GetValueByKey(db, fmt.Sprintf(model.KeyChatGPTMaxTokens, botId))
	if maxTokens != "" {
		config.MaxTokens, _ = strconv.Atoi(maxTokens)
	}
	topP, _ := model.GetValueByKey(db, fmt.Sprintf(model.KeyChatGPTTopP, botId))
	if topP != "" {
		config.TopP, _ = strconv.ParseFloat(topP, 64)
	}
	prePenalty, _ := model.GetValueByKey(db, fmt.Sprintf(model.KeyChatGPTPresencePenalty, botId))
	if prePenalty != "" {
		config.PresencePenalty, _ = strconv.ParseFloat(prePenalty, 64)
	}
	freePenalty, _ := model.GetValueByKey(db, fmt.Sprintf(model.KeyChatGPTFrequencyPenalty, botId))
	if freePenalty != "" {
		config.FrequencyPenalty, _ = strconv.ParseFloat(freePenalty, 64)
	}
	historyCount, _ := model.GetValueByKey(db, fmt.Sprintf(model.KeyChatGPTHistoryCount, botId))
	if historyCount != "" {
		config.HistoryCount, _ = strconv.Atoi(historyCount)
	}
	timeout, _ := model.GetValueByKey(db, fmt.Sprintf(model.KeyChatGPTTimeout, botId))
	if timeout != "" {
		config.Timeout, _ = strconv.Atoi(timeout)
	}
	if config.HttpProxy != "" {
		proxyUrl, err := url.Parse(config.HttpProxy)
		if err != nil {
			return nil, err
		}
		transport := &http.Transport{
			Proxy: http.ProxyURL(proxyUrl),
		}
		defaultConfig.HTTPClient = &http.Client{
			Transport: transport,
		}
	}
	if config.BaseUrl != "" {
		defaultConfig.BaseURL = config.BaseUrl
	}

	c := openai.NewClientWithConfig(defaultConfig)
	cg := &ChatGPT{
		BotId:  botId,
		Config: config,
		API:    c,
	}
	chatGPTBotMap[botId] = cg
	return cg, nil
}
