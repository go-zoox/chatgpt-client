package chatgptclient

import (
	"bytes"
	"fmt"
	"text/template"

	"github.com/go-zoox/core-utils/safe"
	"github.com/go-zoox/datetime"
)

func buildPrompt(context string, messages *safe.List, maxLength int) (prompt []byte, err error) {
	contextTmpl := template.New("context")

	if contextTmpl, err = contextTmpl.Parse(context); err != nil {
		return nil, fmt.Errorf("failed to parse context message: %v", err)
	}
	// 使用 buffer 来格式化字符串，再把 buffer 写入 result
	buffer := &bytes.Buffer{}
	err = contextTmpl.Execute(buffer, map[string]interface{}{
		"date": datetime.Now().Format("YYYY-MM-DD"),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to generate context message from template: %v", err)
	}

	contextMessage := buffer.String()
	endMessage := "ChatGPT:"
	endOfText := "<|endoftext|>\n\n"

	charCountRes := len(contextMessage) + len(endMessage)
	coreMessages := ""

	var currentMessage string
	var currentTextLength int
	messages.Reverse().ForEach(func(i interface{}) (done bool) {
		message := i.(*Message)
		if message.IsChatGPT {
			currentMessage = fmt.Sprintf("ChatGPT:\n\n%s", message.Text)
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
