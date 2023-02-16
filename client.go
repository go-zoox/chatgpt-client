package chatgptclient

import (
	"math"

	"github.com/go-zoox/core-utils/strings"

	openai "github.com/go-zoox/openai-client"
)

// Client is the ChatGPT Client.
type Client interface {
	Ask(question []byte) ([]byte, error)
	//
	GetOrCreateConversation(id string, cfg *ConversationConfig) (Conversation, error)
}

type client struct {
	core openai.Client
	cfg  *Config
}

// Config is the configuration for the ChatGPT Client.
type Config struct {
	APIKey            string `json:"api_key"`
	APIServer         string `json:"api_server"`
	MaxResponseTokens int    `json:"max_response_tokens"`
}

// New creates a new ChatGPT Client.
func New(cfg *Config) (Client, error) {
	if cfg.MaxResponseTokens == 0 {
		cfg.MaxResponseTokens = DefaultMaxResponseTokens
	}

	core, err := openai.New(&openai.Config{
		APIKey:    cfg.APIKey,
		APIServer: cfg.APIServer,
	})
	if err != nil {
		return nil, err
	}

	return &client{
		core: core,
		cfg:  cfg,
	}, nil
}

func (c *client) Ask(question []byte) (answer []byte, err error) {
	numTokens := float64(len(question))
	maxTokens := math.Max(1, math.Min(openai.MaxTokens-numTokens, float64(c.cfg.MaxResponseTokens)))

	completion, err := c.core.CreateCompletion(&openai.CreateCompletionRequest{
		Model:       openai.ModelTextDavinci003,
		Prompt:      string(question),
		MaxTokens:   int(maxTokens),
		Temperature: 1,
	})
	if err != nil {
		return nil, err
	}

	return []byte(strings.TrimSpace(completion.Choices[0].Text)), nil
}

func (c *client) GetOrCreateConversation(id string, cfg *ConversationConfig) (Conversation, error) {
	if cfg.ID == "" {
		cfg.ID = id
	}

	return NewConversation(c, cfg)
}
