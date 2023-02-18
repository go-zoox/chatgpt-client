package chatgptclient

import (
	"fmt"
	"strings"

	"github.com/go-zoox/core-utils/array"
	"github.com/go-zoox/core-utils/safe"
)

func buildPrompt(context, date string, messages *safe.List, maxLength int) (prompt []byte, err error) {
	// for _, message := range messages {
	// 	if message.IsChatGPT {
	// 		textMessages = append(textMessages, fmt.Sprintf("ChatGPT:\n\n%s", message.Text))
	// 	} else {
	// 		textMessages = append(textMessages, fmt.Sprintf("User:\n\n%s", message.Text))
	// 	}
	// }

	charCountRes := 0
	coreMessages := []string{}
	messages.Reverse().ForEach(func(i interface{}) (done bool) {
		if maxLength > 0 && charCountRes >= maxLength {
			return true
		}

		message := i.(*Message)
		charCountRes += len(message.Text)

		if message.IsChatGPT {
			coreMessages = append(coreMessages, fmt.Sprintf("ChatGPT:\n\n%s", message.Text))
		} else {
			coreMessages = append(coreMessages, fmt.Sprintf("User:\n\n%s", message.Text))
		}

		return false
	})

	contextMessage := fmt.Sprintf("%s\nCurrent date: %s", context, date)
	textMessages := append([]string{contextMessage}, array.Reverse(coreMessages)...)
	textMessages = append(textMessages, "ChatGPT:")

	return []byte(strings.Join(textMessages, "<|endoftext|>\n\n")), nil
}
