package chatgptclient

import "time"

// Message is a message.
type Message struct {
	ID             string    `json:"id"`
	Text           string    `json:"text"`
	IsChatGPT      bool      `json:"is_chatgpt"`
	ConversationID string    `json:"conversation"`
	CreatedAt      time.Time `json:"created_at"`

	// User used for support multiple users, like group chat
	User string `json:"user"`
}
