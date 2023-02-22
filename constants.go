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
const DefaultMaxConversations = 1000

// DefaultConversationMaxAge is timeout for each conversation.
const DefaultConversationMaxAge = 2 * time.Hour

// DefaultContext is the default context for instructions to ChatGPT.
const DefaultContext = "Instructions:\nYou are ChatGPT, a large language model trained by OpenAI.\nCurrent date: {{.date}}"
