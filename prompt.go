package chatgptclient

import (
	"fmt"
	"strings"

	"github.com/go-zoox/core-utils/safe"
)

func buildPrompt(context, date string, messages *safe.List) (prompt []byte, err error) {
	textMessages := []string{
		fmt.Sprintf("%s\nCurrent date: %s", context, date),
	}

	// for _, message := range messages {
	// 	if message.IsChatGPT {
	// 		textMessages = append(textMessages, fmt.Sprintf("ChatGPT:\n\n%s", message.Text))
	// 	} else {
	// 		textMessages = append(textMessages, fmt.Sprintf("User:\n\n%s", message.Text))
	// 	}
	// }

	messages.ForEach(func(i interface{}) {
		message := i.(*Message)
		if message.IsChatGPT {
			textMessages = append(textMessages, fmt.Sprintf("ChatGPT:\n\n%s", message.Text))
		} else {
			textMessages = append(textMessages, fmt.Sprintf("User:\n\n%s", message.Text))
		}
	})

	textMessages = append(textMessages, "ChatGPT:")

	return []byte(strings.Join(textMessages, "<|endoftext|>\n\n")), nil
}
