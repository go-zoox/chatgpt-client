package chatgptclient

import (
	"testing"

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

func TestMaxTokens(t *testing.T) {
	question := `Instructions:\nYou are ChatGPT, a large language model trained by OpenAI.\nCurrent date: 2023-02-18\u003c|endoftext|\u003e\n\nChatGPT:\n\n在经济学领域，数据可视化往往能够有效帮助经济学家们更好地理解数据。一些有用的案例如下：\n\n• 经济学家可以利用数据可视化来分析和评估市场行为。\n\n• 将复杂的数据抽象化，帮助经济学家们更好地理解经济动态。\n\n• 可以展示国家和地区之间经济活动的比较，以供研究和分析。\n\n• 可以更好地掌握市场风险，显示金融市场的波动及趋势。\n\n• 帮助经济学家预测未来经济状况，并做出更具有洞察力的决策。\u003c|endoftext|\u003e\n\nUser:\n\n扩展第一个案例\u003c|endoftext|\u003e\n\nChatGPT:\n\n通过数据可视化分析市场行为，经济学家可以更好地了解不同市场之间的竞争方式、需求情况以及产品特性，从而制定更有竞争力的市场策略。\u003c|endoftext|\u003e\n\nUser:\n\n具体怎么做的\u003c|endoftext|\u003e\n\nChatGPT:\n\n首先，经济学家可以从历史数据中收集时间序列数据，以了解不同市场之间的竞争状况。然后，可以使用可视化工具对数据进行分析和探索，进一步了解每个市场之间的关联性和特性。最后，利用可视化工具可以成功预测未来市场发展趋势，以帮助经济学家制定更有效的策略。\u003c|endoftext|\u003e\n\nUser:\n\n怎么从历史数据中收集时间序列数据\u003c|endoftext|\u003e\n\nChatGPT:\n\n从历史数据中收集时间序列数据可以通过使用不同的数据挖掘方法来完成，比如统计分析、回归分析、聚类分析等等。另外，可以利用基于Python的可视化库来更好地可视化历史数据，从而发现有趣的现象和趋势。\u003c|endoftext|\u003e\n\nUser:\n\n这些历史数据怎么找\u003c|endoftext|\u003e\n\nChatGPT:\n\n你可以从很多来源获取历史数据，比如政府部门发布的公共数据集、经济机构发布的可用数据集，或者从社交媒体中获取到的历史数据。你也可以编写一些数据爬虫，抓取网络上的历史数据并进行分析处理。\u003c|endoftext|\u003e\n\nUser:\n\n具体需要哪些数据\u003c|endoftext|\u003e\n\nChatGPT:\n\n市场行为中需要分析的数据可能有：行为数据、交易记录、市场行情、宏观经济指标等等。这些数据都可以从上面提到的各种来源中获得。\u003c|endoftext|\u003e\n\nUser:\n\n还有吗\u003c|endoftext|\u003e\n\nChatGPT:\n\n还有很多数据可用于市场研究，例如用户行为数据、产品数据、竞争研究数据等等，你可以根据你的研究需求来考虑哪些数据是必要的。\u003c|endoftext|\u003e\n\nUser:\n\n可以帮我列一些 数据可视化 python想问的论文题目吗\u003c|endoftext|\u003e\n\nChatGPT:\n\n以下是一些Python数据可视化相关的论文题目：\n\n1. 基于Python的经济学数据可视化研究\n2. 利用Python进行数据可视化处理的应用\n3. 基于Python的金融数据可视化技术\n4. 运用Python对数据可视化进行实证分析\n5. 探索式数据可视化的Python工具实现\n6. 基于Python的复杂数据可视化研究\u003c|endoftext|\u003e\n\nUser:\n\n还有吗\u003c|endoftext|\u003e\n\nChatGPT:\n\n其他一些论文题目包括：\n\n7. 基于Python的\u003c|endoftext|\u003e\n\nUser:\n\n还有吗\u003c|endoftext|\u003e\n\nChatGPT:\n\n社会科学数据可视化研究\n8. 基于Python的大数据可视化分析\n9. 利用Python构建复杂数据可视化交互界面\n10. 应用Python可视化算法于经济学研究\n11. 基于Python的演化式数据可视化方法\u003c|endoftext|\u003e\n\nUser:\n\n还有吗\u003c|endoftext|\u003e\n\nChatGPT:\n\n12. 基于Python的动态数据可视化研究\n13. 利用Python进行数据可视化对比分析\n14. 基于Python的高维数据可视化研究\n15. 应用Python进行交互式数\u003c|endoftext|\u003e\n\nUser:\n\n还有吗\u003c|endoftext|\u003e\n\nChatGPT:\n\n数据可视化的研究\n16. 基于Python的网络数据可视化方法\n17. 基于Python的嵌入式数据可视化研究\n18. 基于Python的事实推理式数据可视化技术\u003c|endoftext|\u003e\n\nUser:\n\n还有吗\u003c|endoftext|\u003e\n\nChatGPT:\n\n目前来看，这就是一些关于Python数据可视化的相关论文题目。\u003c|endoftext|\u003e\n\nUser:\n\n以利用Python进行数据可视化对比分析为题列一个详细的硕士论文大纲\u003c|endoftext|\u003e\n\nChatGPT:`
	// numTokens := float64(len(question))
	// maxTokens := math.Max(DefaultMaxResponseTokens, math.Min(openai.MaxTokens-numTokens, float64(DefaultMaxResponseTokens)))

	// fmt.Println("len(question):", len(question))
	// fmt.Println("numTokens:", numTokens)
	// fmt.Println("openai.MaxTokens-numTokens:", openai.MaxTokens-numTokens)
	// fmt.Println("DefaultMaxResponseTokens:", DefaultMaxResponseTokens)
	// fmt.Println("maxTokens:", maxTokens)

	maxTokens := calculationPromptMaxTokens(len(question), 4097, DefaultMaxResponseTokens)
	if maxTokens != DefaultMaxResponseTokens {
		t.Fatalf("expected %d, but got %d", DefaultMaxResponseTokens, maxTokens)
	}
}
