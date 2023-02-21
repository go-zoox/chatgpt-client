package chatgptclient

import (
	"math"

	"github.com/go-zoox/core-utils/strings"
	"github.com/go-zoox/lru"

	openai "github.com/go-zoox/openai-client"
)

// Client is the ChatGPT Client.
type Client interface {
	Ask(question []byte) ([]byte, error)
	//
	GetOrCreateConversation(id string, cfg *ConversationConfig) (Conversation, error)
	//
	ResetConversations() error
	ResetConversation(id string) error
}

type client struct {
	core openai.Client
	cfg  *Config
	//
	conversationsCache *lru.LRU
}

// Config is the configuration for the ChatGPT Client.
type Config struct {
	APIKey                   string `json:"api_key"`
	APIServer                string `json:"api_server"`
	MaxRequestResponseTokens int    `json:"max_request_response_tokens"`
	MaxResponseTokens        int    `json:"max_response_tokens"`
	MaxConversations         int    `json:"max_conversations"`
	ConversationMaxAge       int    `json:"conversation_max_age"`
}

// New creates a new ChatGPT Client.
func New(cfg *Config) (Client, error) {
	if cfg.MaxRequestResponseTokens == 0 {
		cfg.MaxRequestResponseTokens = DefaultMaxRequestResponseTokens
	}

	if cfg.MaxResponseTokens == 0 {
		cfg.MaxResponseTokens = DefaultMaxResponseTokens
	}

	if cfg.MaxConversations == 0 {
		cfg.MaxConversations = DefaultMaxConversations
	}

	core, err := openai.New(&openai.Config{
		APIKey:    cfg.APIKey,
		APIServer: cfg.APIServer,
	})
	if err != nil {
		return nil, err
	}

	return &client{
		core:               core,
		cfg:                cfg,
		conversationsCache: lru.New(cfg.MaxConversations),
	}, nil
}

func (c *client) Ask(question []byte) (answer []byte, err error) {
	// numTokens := float64(len(question))
	// maxTokens := math.Max(float64(c.cfg.MaxResponseTokens), math.Min(openai.MaxTokens-numTokens, float64(c.cfg.MaxResponseTokens)))

	completion, err := c.core.CreateCompletion(&openai.CreateCompletionRequest{
		Model:     openai.ModelTextDavinci003,
		Prompt:    string(question),
		MaxTokens: calculationPromptMaxTokens(len(question), c.cfg.MaxResponseTokens),
	})
	if err != nil {
		return nil, err
	}

	return []byte(strings.TrimSpace(completion.Choices[0].Text)), nil
}

func (c *client) GetOrCreateConversation(id string, cfg *ConversationConfig) (conversation Conversation, err error) {
	if cfg.ID == "" {
		cfg.ID = id
	}
	if cfg.MaxAge == 0 {
		cfg.MaxAge = DefaultConversationMaxAge
	}
	if cfg.MaxRequestTokens == 0 {
		cfg.MaxRequestTokens = c.cfg.MaxRequestResponseTokens - c.cfg.MaxResponseTokens
	}

	if cache, ok := c.conversationsCache.Get(cfg.ID); ok {
		if c, ok := cache.(Conversation); ok {
			conversation = c
			return conversation, nil
		}
	}

	conversation, err = NewConversation(c, cfg)
	if err != nil {
		return nil, err
	}

	c.conversationsCache.Set(id, conversation, cfg.MaxAge)

	return conversation, nil
}

func (c *client) ResetConversations() error {
	c.conversationsCache.Clear()

	return nil
}

func (c *client) ResetConversation(id string) error {
	c.conversationsCache.Delete(id)

	return nil
}

func calculationPromptMaxTokens(questLength, MaxResponseTokens int) int {
	numTokens := float64(questLength)
	maxTokens := math.Max(float64(MaxResponseTokens), math.Min(openai.MaxTokens-numTokens, float64(MaxResponseTokens)))

	return int(maxTokens)
}
