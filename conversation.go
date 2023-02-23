package chatgptclient

import (
	"fmt"
	"time"

	"github.com/go-zoox/core-utils/safe"
	"github.com/go-zoox/logger"
	"github.com/go-zoox/uuid"
)

// Conversation is the conversation interface.
type Conversation interface {
	Ask(question []byte, cfg ...*ConversationAskConfig) (answer []byte, err error)
	//
	ID() string
	Messages() *safe.List
	//
	BuildPrompt() (prompt []byte, err error)
}

type conversation struct {
	client      *client
	id          string
	messages    *safe.List
	messagesMap map[string]bool
	//
	cfg *ConversationConfig
}

// ConversationConfig is the configuration for creating a new Conversation.
type ConversationConfig struct {
	ID               string
	Context          string
	Language         string
	MaxMessages      int
	MaxAge           time.Duration
	MaxRequestTokens int
	Model            string `json:"model"`
}

// ConversationAskConfig is the configuration for ask question.
type ConversationAskConfig struct {
	ID        string    `json:"id"`
	User      string    `json:"user"`
	CreatedAt time.Time `json:"created_at"`
}

// NewConversation creates a new Conversation.
func NewConversation(client *client, cfg *ConversationConfig) (Conversation, error) {
	if cfg.ID == "" {
		cfg.ID = uuid.V4()
	}
	if cfg.Context == "" {
		cfg.Context = DefaultContext
		if cfg.Language != "" {
			cfg.Context = fmt.Sprintf("%s\nLanuage: %s", cfg.Context, cfg.Language)
		}
	}
	if cfg.MaxMessages == 0 {
		cfg.MaxMessages = 100
	}

	return &conversation{
		client:   client,
		id:       cfg.ID,
		messages: safe.NewList(cfg.MaxMessages),
		cfg:      cfg,
	}, nil
}

func (c *conversation) Ask(question []byte, cfg ...*ConversationAskConfig) (answer []byte, err error) {
	cfgX := &ConversationAskConfig{}
	if len(cfg) > 0 && cfg[0] != nil {
		cfgX = cfg[0]
	}
	if cfgX.ID == "" {
		logger.Warnf("ask question id is recommand")

		cfgX.ID = uuid.V4()
	}
	if cfgX.CreatedAt.IsZero() {
		cfgX.CreatedAt = time.Now()
	}

	if _, ok := c.messagesMap[cfgX.ID]; ok {
		return nil, fmt.Errorf("duplicate message(id: %s) to ask", cfgX.ID)
	}

	c.messagesMap[cfgX.ID] = true
	c.messages.Push(&Message{
		ID:             cfgX.ID,
		Text:           string(question),
		IsChatGPT:      false,
		ConversationID: c.id,
		User:           cfgX.User,
		CreatedAt:      cfgX.CreatedAt,
	})

	prompt, err := c.BuildPrompt()
	if err != nil {
		return nil, fmt.Errorf("failed to build prompt: %v", err)
	}

	answer, err = c.client.Ask(prompt, &AskConfig{
		Model: c.cfg.Model,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to ask: %v", err)
	}

	c.messages.Push(&Message{
		ID:             uuid.V4(),
		Text:           string(answer),
		IsChatGPT:      true,
		ConversationID: c.id,
	})

	return answer, nil
}

func (c *conversation) ID() string {
	return c.id
}

func (c *conversation) Messages() *safe.List {
	return c.messages
}

func (c *conversation) BuildPrompt() (prompt []byte, err error) {
	return buildPrompt(
		c.cfg.Context,
		c.messages,
		c.cfg.MaxRequestTokens,
	)
}
