package chatgptclient

import (
	"time"

	openai "github.com/go-zoox/openai-client"
)

// DefaultMaxRequestResponseTokens ...
const DefaultMaxRequestResponseTokens = openai.MaxTokens

// DefaultMaxResponseTokens is the default response text max tokens.
const DefaultMaxResponseTokens = 1000

// DefaultMaxConversations is the default max conversation cache count.
const DefaultMaxConversations = 10000

// DefaultConversationMaxAge is timeout for each conversation.
//
//	default: 30 day
const DefaultConversationMaxAge = 30 * 24 * time.Hour

// DefaultContextIntl is the default context for instructions to ChatGPT.
var DefaultContextIntl = map[string]string{
	"EN_US": "Instructions:\nYou are {{.chatgpt_name}}, a large language model trained by OpenAI.\nCurrent date: {{.date}}",
	"ZH_CN": "引导:\n你是 {{.chatgpt_name}}，一个由 OpenAI 训练的大语言模型。\n当前日期: {{.date}}",
}

// DefaultContext is the default context for instructions to ChatGPT.
var DefaultContext = DefaultContextIntl["EN_US"]
