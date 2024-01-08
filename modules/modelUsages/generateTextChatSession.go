package modelUsages

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"main/modules/apikey"
	"main/modules/responseProcess"
	"os"
	"regexp"
	"strings"

	"github.com/fatih/color"
	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

func GenerateTextChatSession(genAIModelName string) {

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
	// Disable not funny safety setting
	model.SafetySettings = []*genai.SafetySetting{
		{
			Category:  genai.HarmCategoryDangerousContent,
			Threshold: genai.HarmBlockNone,
		},
		{
			Category:  genai.HarmCategoryHarassment,
			Threshold: genai.HarmBlockNone,
		},
		{
			Category:  genai.HarmCategoryHateSpeech,
			Threshold: genai.HarmBlockNone,
		},
		{
			Category:  genai.HarmCategorySexuallyExplicit,
			Threshold: genai.HarmBlockNone,
		},
	}
	chatSession.History = []*genai.Content{}

	fmt.Printf("Chat session with Gemini AI Model(%s) started.\n", genAIModelName)
	yellowColorBoldPrint := color.New(color.FgYellow, color.Bold)
	cyanColorBoldPrint := color.New(color.FgCyan, color.Bold)
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

		response, err := chatSession.SendMessage(context, genai.Text(question))
		if err != nil {
			log.Fatal(err)
		}

		result, _ := responseProcess.GetGeminiAITextOnlyResponseStruct(question, response)
		answer := result.Response[0]
		cyanColorBoldPrint.Printf(" - %s\n", genAIModelName)
		fmt.Println(boldifyTextInMarkdownRule(answer))
		addMessageToChatSessionHistory(chatSession, "user", question)
		addMessageToChatSessionHistory(chatSession, "model", answer)
	}

}

func addMessageToChatSessionHistory(chatSession *genai.ChatSession, role string, text string) {
	message := &genai.Content{
		Parts: []genai.Part{
			genai.Text(text),
		},
		Role: role,
	}
	chatSession.History = append(chatSession.History, message)
}

func boldifyTextInMarkdownRule(input string) string {
	// Define a regular expression to match double asterisks within words
	re := regexp.MustCompile(`\*\*(.+?)\*\*`)

	// Replace double asterisks with Markdown bold syntax
	modifiedInput := re.ReplaceAllStringFunc(input, func(match string) string {
		word := strings.Trim(match, "**")
		boldPrint := color.New(color.Bold)
		boldWord := boldPrint.Sprintf("%s", word)
		return boldWord
	})

	return modifiedInput
}
