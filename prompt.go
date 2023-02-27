package chatgptclient

import (
	"fmt"

	"github.com/go-zoox/core-utils/strings"

	"github.com/go-zoox/core-utils/safe"
	"github.com/go-zoox/datetime"
)

func buildPrompt(context string, messages *safe.List, maxLength int, chatGPTName string) (prompt []byte, err error) {
	if chatGPTName == "" {
		chatGPTName = "ChatGPT"
	}

	contextMessage, err := strings.Template(context, map[string]interface{}{
		"date":         datetime.Now().Format("YYYY-MM-DD"),
		"chatgpt_name": chatGPTName,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to render context message: %v", err)
	}

	endMessage := fmt.Sprintf("%s:", chatGPTName)
	endOfText := "<|endoftext|>\n\n"

	charCountRes := len(contextMessage) + len(endMessage)
	coreMessages := ""

	var currentMessage string
	var currentTextLength int
	messages.Reverse().ForEach(func(i interface{}) (done bool) {
		message := i.(*Message)
		if message.IsChatGPT {
			currentMessage = fmt.Sprintf("%s:\n\n%s", chatGPTName, message.Text)
		} else {
			if message.User != "" {
				currentMessage = fmt.Sprintf("%s:\n\n%s", message.User, message.Text)
			} else {
				currentMessage = fmt.Sprintf("User:\n\n%s", message.Text)
			}
		}

		currentTextLength = len(currentMessage) + len(endOfText)
		if maxLength > 0 && charCountRes+currentTextLength >= maxLength {
			return true
		}

		charCountRes += currentTextLength
		coreMessages = fmt.Sprintf("%s%s%s", currentMessage, endOfText, coreMessages)

		return false
	})

	// textMessages := append([]string{contextMessage}, array.Reverse(coreMessages)...)
	// textMessages = append(textMessages, chat)

	// return []byte(strings.Join(textMessages, "<|endoftext|>\n\n")), nil

	message := fmt.Sprintf("%s%s%s%s", contextMessage, endOfText, coreMessages, endMessage)
	return []byte(message), nil
}
