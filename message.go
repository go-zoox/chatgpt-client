package chatgptclient

import (
	"fmt"
	"time"

	"github.com/go-zoox/core-utils/safe"
	"github.com/go-zoox/core-utils/strings"
	"github.com/go-zoox/datetime"
)

// Message is a message.
type Message struct {
	ID             string    `json:"id"`
	Text           string    `json:"text"`
	IsChatGPT      bool      `json:"is_chatgpt"`
	ConversationID string    `json:"conversation"`
	CreatedAt      time.Time `json:"created_at"`

	// User used for support multiple users, like group chat
	User string `json:"user"`

	//
	Role string
}

func buildMessages(context string, messages *safe.List, maxLength int) (messagesX []*Message, err error) {
	contextMessage, err := strings.Template(context, map[string]interface{}{
		"date": datetime.Now().Format("YYYY-MM-DD"),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to render context message: %v", err)
	}

	charCountRes := len(contextMessage)

	var currentTextLength int
	messages.Reverse().ForEach(func(i interface{}) (done bool) {
		currentMessage := i.(*Message)

		currentTextLength = len(currentMessage.Text)
		if maxLength > 0 && charCountRes+currentTextLength >= maxLength {
			return true
		}

		charCountRes += currentTextLength
		messagesX = append([]*Message{currentMessage}, messagesX...)
		return false
	})

	// messagesX = append(messagesX, &Message{
	// 	Role: "assistant",
	// 	Text: contextMessage,
	// })

	return
}
