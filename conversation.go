package chatgptclient

import (
	"fmt"
	"time"

	"github.com/go-zoox/core-utils/safe"
	"github.com/go-zoox/core-utils/strings"
	"github.com/go-zoox/logger"
	openai "github.com/go-zoox/openai-client"
	"github.com/go-zoox/uuid"
)

// Conversation is the conversation interface.
type Conversation interface {
	Ask(question []byte, cfg ...*ConversationAskConfig) (answer []byte, err error)
	IsQuestionAsked(id string) (err error)
	//
	ID() string
	Messages() *safe.List
	//
	BuildPrompt() (prompt []byte, err error)
	//
	SetModel(model string) error
	GetModel() string
}

type conversation struct {
	client      *client
	id          string
	messages    *safe.List
	messagesMap *safe.Map
	//
	cfg *ConversationConfig
}

// ConversationConfig is the configuration for creating a new Conversation.
type ConversationConfig struct {
	ID                       string
	Context                  string
	Language                 string
	MaxMessages              int
	MaxAge                   time.Duration
	MaxRequestResponseTokens int64
	Model                    string `json:"model"`
	ChatGPTName              string `json:"chatgpt_name"`

	//
	MaxRequestTokens  int64
	MaxResponseTokens int64
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

	// ensure language
	if cfg.Language != "" {
		cfg.Language = strings.ToUpper(cfg.Language)
		// @TODO
		if vc, ok := DefaultContextIntl[cfg.Language]; ok {
			DefaultContext = vc
		}
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

	c := &conversation{
		client:      client,
		id:          cfg.ID,
		messages:    safe.NewList(cfg.MaxMessages),
		messagesMap: safe.NewMap(),
		cfg:         cfg,
	}

	if c.cfg.Model == "" {
		// c.SetModel(openai.ModelGPT_4_32K)
		c.SetModel(openai.ModelGPT3_5Turbo)
	} else {
		c.SetModel(c.cfg.Model)
	}

	return c, nil
}

func (c *conversation) IsQuestionAsked(id string) (err error) {
	if c.messagesMap.Has(id) {
		return fmt.Errorf("duplicate message(id: %s) to ask", id)
	}

	return nil
}

func (c *conversation) Ask(question []byte, cfg ...*ConversationAskConfig) (answer []byte, err error) {
	cfgX := &ConversationAskConfig{}
	if len(cfg) > 0 && cfg[0] != nil {
		cfgX = cfg[0]
	}
	if cfgX.ID == "" {
		logger.Warnf("question id is recommand")

		cfgX.ID = uuid.V4()
	}
	if cfgX.CreatedAt.IsZero() {
		cfgX.CreatedAt = time.Now()
	}

	if c.messagesMap.Has(cfgX.ID) {
		return nil, fmt.Errorf("duplicate message(id: %s) to ask", cfgX.ID)
	}

	c.messagesMap.Set(cfgX.ID, true)
	c.messages.Push(&Message{
		ID:             cfgX.ID,
		Text:           string(question),
		IsChatGPT:      false,
		ConversationID: c.id,
		User:           cfgX.User,
		Role:           "user",
		CreatedAt:      cfgX.CreatedAt,
	})

	prompt, err := c.BuildPrompt()
	if err != nil {
		return nil, fmt.Errorf("failed to build prompt: %v", err)
	}

	messages, err := c.BuildMessages()
	if err != nil {
		return nil, fmt.Errorf("failed to build messages: %v", err)
	}

	answer, err = c.client.Ask(&AskConfig{
		Model:    c.cfg.Model,
		Prompt:   string(prompt),
		Messages: messages,
		//
		MaxRequestResponseTokens: int(c.cfg.MaxRequestResponseTokens),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to ask: %v", err)
	}

	c.messages.Push(&Message{
		ID:             uuid.V4(),
		Text:           string(answer),
		IsChatGPT:      true,
		ConversationID: c.id,
		Role:           "assistant",
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
		int(c.cfg.MaxRequestTokens),
		c.cfg.ChatGPTName,
	)
}

func (c *conversation) BuildMessages() (messages []*Message, err error) {
	return buildMessages(
		c.cfg.Context,
		c.messages,
		int(c.cfg.MaxRequestTokens),
	)
}

func (c *conversation) GetModel() string {
	return c.cfg.Model
}

func (c *conversation) SetModel(model string) error {
	c.cfg.Model = model

	// generate MaxRequestResponseTokens
	c.cfg.MaxRequestResponseTokens = openai.GetMaxTokens(c.cfg.Model)
	c.cfg.MaxRequestTokens = c.cfg.MaxRequestResponseTokens - c.cfg.MaxResponseTokens

	return nil
}
