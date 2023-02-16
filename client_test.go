package chatgptclient

import (
	_ "github.com/go-zoox/dotenv"
)

// func TestAsk(t *testing.T) {
// 	client, _ := New(&Config{
// 		APIKey: os.Getenv("API_KEY"),
// 	})

// 	answer, err := client.Ask([]byte("\"你好\"用英语怎么说？"))
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	fmt.PrintJSON(string(answer))
// }

// func TestConversation(t *testing.T) {
// 	client, err := New(&Config{
// 		APIKey: os.Getenv("API_KEY"),
// 	})
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	conversation, _ := client.GetOrCreateConversation(uuid.V4(), &ConversationConfig{})

// 	var question []byte
// 	var answer []byte

// 	question = []byte("OpenAI 是什么？")
// 	fmt.Printf("question: %s\n", question)
// 	answer, err = conversation.Ask(question)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	fmt.Printf("answer: %s\n\n", answer)

// 	question = []byte("可以展开讲讲吗？")
// 	fmt.Printf("question: %s\n", question)
// 	answer, err = conversation.Ask(question)
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
