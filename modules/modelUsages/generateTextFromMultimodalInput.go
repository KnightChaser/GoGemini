package modelUsages

import (
	"context"
	"log"
	"main/modules/apikey"
	"os"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

// Send a single text query to the model and receive the response from the server
func GenerateTextFromMultimodalInput(question string, inputImagePathArray []string) *genai.GenerateContentResponse {

	apiKeyString := apikey.GetGoogleGenAIAPIKey("default")

	// Ready to bring the model
	context := context.Background()
	client, err := genai.NewClient(context, option.WithAPIKey(apiKeyString))
	if err != nil {
		log.Panic(err)
	}
	defer client.Close()

	model := client.GenerativeModel("gemini-pro-vision")
	prompt := make([]genai.Part, len(inputImagePathArray)+1)
	for _, imagePath := range inputImagePathArray {
		imageData, error := os.ReadFile(imagePath)
		if error != nil {
			log.Panic(error)
		}
		prompt = append(prompt, genai.ImageData("jpeg", imageData))
	}
	prompt = append(prompt, genai.Text(question))

	response, error := model.GenerateContent(context, prompt...)
	if error != nil {
		log.Panic(error)
	}

	return response

}
