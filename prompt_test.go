package chatgptclient

import (
	"fmt"
	"testing"

	"github.com/go-zoox/core-utils/safe"
	"github.com/go-zoox/datetime"
	"github.com/go-zoox/testify"
)

func TestBuildPrompt(t *testing.T) {
	messages := safe.NewList()
	messages.Push(&Message{
		ID:        "1",
		Text:      "What is OpenAI?",
		IsChatGPT: false,
	})
	messages.Push(&Message{
		ID:        "2",
		Text:      "OpenAI is an artificial intelligence research laboratory that focuses on creating an open source artificial general intelligence. Founded in 2015 by Elon Musk, Sam Altman, Greg Brockman, and Ilya Sutskever, OpenAI's mission is to \"advance digital intelligence in the way that is most likely to benefit humanity as a whole.\" They develop advanced AI algorithms and apply them to problems like natural language processing, robotics, computer vision, and more.",
		IsChatGPT: true,
	})
	messages.Push(&Message{
		ID:        "3",
		Text:      "Can you expand on that?",
		IsChatGPT: false,
	})

	prompt, err := buildPrompt(DefaultContext, messages, 0)
	if err != nil {
		t.Fatal(err)
	}

	expected := fmt.Sprintf(`Instructions:
You are ChatGPT, a large language model trained by OpenAI.
Current date: %s<|endoftext|>

User:

What is OpenAI?<|endoftext|>

ChatGPT:

OpenAI is an artificial intelligence research laboratory that focuses on creating an open source artificial general intelligence. Founded in 2015 by Elon Musk, Sam Altman, Greg Brockman, and Ilya Sutskever, OpenAI's mission is to "advance digital intelligence in the way that is most likely to benefit humanity as a whole." They develop advanced AI algorithms and apply them to problems like natural language processing, robotics, computer vision, and more.<|endoftext|>

User:

Can you expand on that?<|endoftext|>

ChatGPT:`, datetime.Now().Format("YYYY-MM-DD"))

	testify.Equal(t, expected, string(prompt))
}

func TestBuildPromptMultiUsers(t *testing.T) {
	messages := safe.NewList()
	messages.Push(&Message{
		ID:        "1",
		Text:      "What is OpenAI?",
		IsChatGPT: false,
		User:      "Zero",
	})
	messages.Push(&Message{
		ID:        "2",
		Text:      "OpenAI is an artificial intelligence research laboratory that focuses on creating an open source artificial general intelligence. Founded in 2015 by Elon Musk, Sam Altman, Greg Brockman, and Ilya Sutskever, OpenAI's mission is to \"advance digital intelligence in the way that is most likely to benefit humanity as a whole.\" They develop advanced AI algorithms and apply them to problems like natural language processing, robotics, computer vision, and more.",
		IsChatGPT: true,
	})
	messages.Push(&Message{
		ID:        "3",
		Text:      "Can you expand on that?",
		IsChatGPT: false,
		User:      "Amy",
	})

	prompt, err := buildPrompt(DefaultContext, messages, 0)
	if err != nil {
		t.Fatal(err)
	}

	expected := fmt.Sprintf(`Instructions:
You are ChatGPT, a large language model trained by OpenAI.
Current date: %s<|endoftext|>

Zero:

What is OpenAI?<|endoftext|>

ChatGPT:

OpenAI is an artificial intelligence research laboratory that focuses on creating an open source artificial general intelligence. Founded in 2015 by Elon Musk, Sam Altman, Greg Brockman, and Ilya Sutskever, OpenAI's mission is to "advance digital intelligence in the way that is most likely to benefit humanity as a whole." They develop advanced AI algorithms and apply them to problems like natural language processing, robotics, computer vision, and more.<|endoftext|>

Amy:

Can you expand on that?<|endoftext|>

ChatGPT:`, datetime.Now().Format("YYYY-MM-DD"))

	testify.Equal(t, expected, string(prompt))
}

func TestBuildPromptMaxLength(t *testing.T) {
	messages := safe.NewList()
	messages.Push(&Message{
		ID:        "1",
		Text:      "What is OpenAI?",
		IsChatGPT: false,
	})
	messages.Push(&Message{
		ID:        "2",
		Text:      "OpenAI is an artificial intelligence research laboratory that focuses on creating an open source artificial general intelligence. Founded in 2015 by Elon Musk, Sam Altman, Greg Brockman, and Ilya Sutskever, OpenAI's mission is to \"advance digital intelligence in the way that is most likely to benefit humanity as a whole.\" They develop advanced AI algorithms and apply them to problems like natural language processing, robotics, computer vision, and more.",
		IsChatGPT: true,
	})
	messages.Push(&Message{
		ID:        "3",
		Text:      "Can you expand on that?",
		IsChatGPT: false,
	})

	prompt, err := buildPrompt(DefaultContext, messages, 300)
	if err != nil {
		t.Fatal(err)
	}

	expected := fmt.Sprintf(`Instructions:
You are ChatGPT, a large language model trained by OpenAI.
Current date: %s<|endoftext|>

User:

Can you expand on that?<|endoftext|>

ChatGPT:`, datetime.Now().Format("YYYY-MM-DD"))

	testify.Equal(t, expected, string(prompt))
}
