package chatgptclient

// func TestConversation(t *testing.T) {
// 	cfg := &Config{
// 		APIKey: os.Getenv("CHATGPT_API_KEY"),
// 	}

// 	core, err := openai.New(&openai.Config{
// 		APIKey: cfg.APIKey,
// 	})
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	client := &client{
// 		core: core,
// 		cfg:  cfg,
// 	}

// 	if cfg.MaxRequestResponseTokens == 0 {
// 		cfg.MaxRequestResponseTokens = DefaultMaxRequestResponseTokens
// 	}

// 	if cfg.MaxResponseTokens == 0 {
// 		cfg.MaxResponseTokens = DefaultMaxResponseTokens
// 	}

// 	// fmt.Println("MaxResponseTokens:", client.cfg.MaxRequestResponseTokens)
// 	conversation, _ := NewConversation(client, &ConversationConfig{})

// 	var question []byte
// 	var answer []byte

// 	question = []byte("OpenAI 是什么？")
// 	fmt.Printf("question: %s\n", question)
// 	answer, err = conversation.Ask(question, &ConversationAskConfig{
// 		ID: "1",
// 	})
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	fmt.Printf("answer: %s\n\n", answer)

// 	question = []byte("可以展开讲讲吗？")
// 	fmt.Printf("question: %s\n", question)
// 	answer, err = conversation.Ask(question, &ConversationAskConfig{
// 		ID: "1",
// 	})
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	fmt.Printf("answer: %s\n\n", answer)

// 	question = []byte("我们现在讨论的内容是什么？")
// 	fmt.Printf("question: %s\n", question)
// 	answer, err = conversation.Ask(question)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	fmt.Printf("answer: %s\n\n", answer)

// 	prompt, _ := conversation.BuildPrompt()
// 	fmt.Printf("prompt:\n\n%s\n", prompt)
// }
