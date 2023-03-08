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

// DefaultContext is the default context for instructions to ChatGPT.
const DefaultContext = "Instructions:\nYou are {{.chatgpt_name}}, a large language model trained by OpenAI.\nCurrent date: {{.date}}"
