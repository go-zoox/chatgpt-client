# ChatGPT-Client is a ChatGPT Client with Offical OpenAI API.

[![PkgGoDev](https://pkg.go.dev/badge/github.com/go-zoox/chatgpt-client)](https://pkg.go.dev/github.com/go-zoox/chatgpt-client)
[![Build Status](https://github.com/go-zoox/chatgpt-client/actions/workflows/ci.yml/badge.svg?branch=master)](https://github.com/go-zoox/chatgpt-client/actions/workflows/ci.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/go-zoox/chatgpt-client)](https://goreportcard.com/report/github.com/go-zoox/chatgpt-client)
[![Coverage Status](https://coveralls.io/repos/github/go-zoox/ip/badge.svg?branch=master)](https://coveralls.io/github/go-zoox/ip?branch=master)
[![GitHub issues](https://img.shields.io/github/issues/go-zoox/ip.svg)](https://github.com/go-zoox/chatgpt-client/issues)
[![Release](https://img.shields.io/github/tag/go-zoox/ip.svg?label=Release)](https://github.com/go-zoox/chatgpt-client/tags)

## Installation
To install the package, run:
```bash
go get -u github.com/go-zoox/chatgpt-client
```

## Getting Started

### Simple 

```go
package main

import (
	"fmt"
	"log"
	"os"

	_ "github.com/go-zoox/dotenv"
)


import (
  chatgpt "github.com/go-zoox/chatgpt-client"
)

func main(t *testing.T) {
	client, err := chatgpt.New(&Config{
		APIKey: os.Getenv("API_KEY"),
	})
	if err != nil {
		log.Fatal(err)
	}

	var question []byte
	var answer []byte

	question = []byte("OpenAI 是什么？")
	fmt.Printf("question: %s\n", question)
	answer, err = client.Ask(question)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("answer: %s\n\n", answer)
}
```


### Use Conversation

```go
package main

import (
	"fmt"
	"log"
	"os"

	_ "github.com/go-zoox/dotenv"
	"github.com/go-zoox/uuid"
)


import (
  chatgpt "github.com/go-zoox/chatgpt-client"
)

func main(t *testing.T) {
	client, err :=chatgpt.New(&Config{
		APIKey: os.Getenv("API_KEY"),
	})
	if err != nil {
		log.Fatal(err)
	}

	conversation, _ := client.GetOrCreateConversation(uuid.V4(), &chatgpt.ConversationConfig{})

	var question []byte
	var answer []byte

	question = []byte("OpenAI 是什么？")
	fmt.Printf("question: %s\n", question)
	answer, err = conversation.Ask(question)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("answer: %s\n\n", answer)

	question = []byte("可以展开讲讲吗？")
	fmt.Printf("question: %s\n", question)
	answer, err = conversation.Ask(question)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("answer: %s\n\n", answer)

	question = []byte("我们现在讨论的内容是什么？")
	fmt.Printf("question: %s\n", question)
	answer, err = conversation.Ask(question)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("answer: %s\n\n", answer)

	prompt, _ := conversation.BuildPrompt()
	fmt.Printf("prompt:\n\n%s\n", prompt)
}
```

## License
GoZoox is released under the [MIT License](./LICENSE).
