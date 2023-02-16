package chatgptclient

// Message is a message.
type Message struct {
	ID             string `json:"id"`
	Text           string `json:"text"`
	IsChatGPT      bool   `json:"is_chatgpt"`
	ConversationID string `json:"conversation"`
}
