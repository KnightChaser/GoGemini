package modelUsages

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"main/modules/apikey"
	"main/modules/responseProcess"
	"os"

	"github.com/fatih/color"
	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
)

func GenerateTextChatSessionStreaming(genAIModelName string) {

	apiKeyString := apikey.GetGoogleGenAIAPIKey("default")
	// Ready to bring the model
	context := context.Background()
	// fmt.Printf("Obtained an API KEY: %s\n", apiKeyString)
	client, err := genai.NewClient(context, option.WithAPIKey(apiKeyString))
	if err != nil {
		log.Panic(err)
	}
	defer client.Close()

	model := client.GenerativeModel("gemini-pro")
	chatSession := model.StartChat()
	// Disable not funny safety setting by disabling annotation
	// model.SafetySettings = []*genai.SafetySetting{
	// 	{
	// 		Category:  genai.HarmCategoryDangerousContent,
	// 		Threshold: genai.HarmBlockNone,
	// 	},
	// 	{
	// 		Category:  genai.HarmCategoryHarassment,
	// 		Threshold: genai.HarmBlockNone,
	// 	},
	// 	{
	// 		Category:  genai.HarmCategoryHateSpeech,
	// 		Threshold: genai.HarmBlockNone,
	// 	},
	// 	{
	// 		Category:  genai.HarmCategorySexuallyExplicit,
	// 		Threshold: genai.HarmBlockNone,
	// 	},
	// }
	chatSession.History = []*genai.Content{}

	fmt.Printf("Chat session with Gemini AI Model(%s) started.\n", genAIModelName)
	yellowColorBoldPrint := color.New(color.FgYellow, color.Bold)
	cyanColorBoldPrint := color.New(color.FgCyan, color.Bold)
	whiteColorItalicPrint := color.New(color.FgWhite, color.Italic)
	// Start chat session endlessly. User -> Model -> User -> Model
	for {
		// Receive user
		yellowColorBoldPrint.Println(" - user")
		fmt.Print("> ")
		var question string
		reader := bufio.NewReader(os.Stdin)
		question, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		// Counting question token
		tokenQtyResponse, error := model.CountTokens(context, genai.Text(question))
		if error != nil {
			log.Fatal(err)
		}
		whiteColorItalicPrint.Printf("...Prompt length: %d tokens\n", tokenQtyResponse.TotalTokens)

		responseIterative := chatSession.SendMessageStream(context, genai.Text(question))
		var answer string

		// Print response from the model via streaming
		cyanColorBoldPrint.Printf(" - %s\n", genAIModelName)
		for {
			response, error := responseIterative.Next()
			if error == iterator.Done {
				break
			} else if error != nil {
				log.Fatal(error)
			}

			result, _ := responseProcess.GetGeminiAITextOnlyResponseStruct(question, response)
			answerIterationFragment := result.Response[0]
			fmt.Print(responseProcess.BoldifyTextInMarkdownRule(answerIterationFragment))
			answer = answer + answerIterationFragment
		}

		fmt.Println()
		responseProcess.AddMessageToChatSessionHistory(chatSession, "user", question)
		responseProcess.AddMessageToChatSessionHistory(chatSession, "model", answer)
	}

}
