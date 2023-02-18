package chatgptclient

import (
	"fmt"
	"time"

	"github.com/go-zoox/core-utils/safe"
	"github.com/go-zoox/datetime"
	"github.com/go-zoox/uuid"
)

// Conversation is the conversation interface.
type Conversation interface {
	Ask(question []byte) (answer []byte, err error)
	//
	ID() string
	Messages() *safe.List
	//
	BuildPrompt() (prompt []byte, err error)
}

type conversation struct {
	client   *client
	id       string
	messages *safe.List
	//
	cfg *ConversationConfig
}

// ConversationConfig is the configuration for creating a new Conversation.
type ConversationConfig struct {
	ID               string
	Context          string
	MaxMessages      int
	MaxAge           time.Duration
	MaxRequestTokens int
}

// NewConversation creates a new Conversation.
func NewConversation(client *client, cfg *ConversationConfig) (Conversation, error) {
	if cfg.ID == "" {
		cfg.ID = uuid.V4()
	}
	if cfg.Context == "" {
		cfg.Context = DefaultContext
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

func (c *conversation) Ask(question []byte) (answer []byte, err error) {
	c.messages.Push(&Message{
		ID:             uuid.V4(),
		Text:           string(question),
		IsChatGPT:      false,
		ConversationID: c.id,
	})

	prompt, err := c.BuildPrompt()
	if err != nil {
		return nil, fmt.Errorf("failed to build prompt: %v", err)
	}

	answer, err = c.client.Ask(prompt)
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
		datetime.Now().Format("YYYY-MM-DD"),
		c.messages,
		c.cfg.MaxRequestTokens,
	)
}
