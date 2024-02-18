package model

import (
	"encoding/json"
	"gorm.io/gorm"
)

const (
	KeySelectedAccountID       = "selected_account_id"
	KeyRunning                 = "running:%d"
	KeyChatGPTEnabled          = "chat_gpt_enabled:%d"
	KeyChatGPTApiKey           = "chat_gpt_api_key:%d"
	KeyChatGPTHttpProxy        = "chat_gpt_http_proxy:%d"
	KeyChatGPTModel            = "chat_gpt_model:%d"
	KeyChatGPTSystemPrompt     = "chat_gpt_system_prompt:%d"
	KeyChatGPTBaseUrl          = "chat_gpt_base_url:%d"
	KeyChatGPTTemperature      = "chat_gpt_temperature:%d"
	KeyChatGPTMaxTokens        = "chat_gpt_max_tokens:%d"
	KeyChatGPTTopP             = "chat_gpt_top_p:%d"
	KeyChatGPTPresencePenalty  = "chat_gpt_presence_penalty:%d"
	KeyChatGPTFrequencyPenalty = "chat_gpt_frequency_penalty:%d"
	KeyChatGPTHistoryCount     = "chat_gpt_history_count:%d"
	KeyChatGPTTimeout          = "chat_gpt_timeout:%d"
)

type KV struct {
	K string `gorm:"column:k;type:text;not null;primaryKey" json:"k"`
	V string `gorm:"column:v;type:text;not null;" json:"v"`
}

func (m *KV) TableName() string {
	return "kv"
}

// GetValueByKey 根据key获取value
func GetValueByKey(db *gorm.DB, key string) (string, bool) {
	var kv KV
	db.Where("k = ?", key).First(&kv)
	if kv.K == "" {
		return "", false
	}
	return kv.V, true
}

// SetValueByKey 根据key设置value
func SetValueByKey(db *gorm.DB, key, value string) {
	db.Where("k = ?", key).Assign(KV{V: value}).FirstOrCreate(&KV{K: key, V: value})
}

func JSONMarshalToString(v any) (string, error) {
	if v == nil {
		return "null", nil
	}
	b, err := json.Marshal(v)
	if err != nil {
		return "null", err
	}
	return string(b), nil
}
