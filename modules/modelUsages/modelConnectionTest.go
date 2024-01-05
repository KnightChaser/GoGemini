package modelUsages

import (
	"context"
	"log"
	"main/modules/apikey"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

// Let's try to select the model! Check the existence and availability of the model provided in Google Gen. AI by the given name
func ModelConnectionTest(modelName string) (bool, error) {

	apiKeyString := apikey.GetGoogleGenAIAPIKey("default")

	// Ready to bring the model
	context := context.Background()
	// fmt.Printf("Obtained an API KEY: %s\n", apiKeyString)
	client, err := genai.NewClient(context, option.WithAPIKey(apiKeyString))
	if err != nil {
		return false, err
	}
	defer client.Close()

	model := client.GenerativeModel(modelName)

	log.Printf("Model obtained: %v\n", model)
	return true, nil

}
