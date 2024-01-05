package modelUsages

import (
	"context"
	"fmt"
	"log"
	"main/modules/apikey"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

// getAIModelName: The GenAI model to use like "Gemini-Pro"
// question		 : The input query to the model
func GenerateTextFromTextOnlyInput(genAIModelName string, question string) *genai.GenerateContentResponse {

	apiKeyString := apikey.GetGoogleGenAIAPIKey("default")

	// Ready to bring the model
	context := context.Background()
	fmt.Printf("Obtained an API KEY: %s\n", apiKeyString)
	client, err := genai.NewClient(context, option.WithAPIKey(apiKeyString))
	if err != nil {
		log.Panic(err)
	}
	defer client.Close()

	// For text-only input, use the gemini-pro model
	model := client.GenerativeModel(genAIModelName)
	modelInput := question
	fmt.Printf("Question: %s\n", modelInput)
	response, err := model.GenerateContent(context, genai.Text(modelInput))
	if err != nil {
		log.Panic(err)
	}

	return response

}
