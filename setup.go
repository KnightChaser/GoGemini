package main

import (
	"context"
	"fmt"
	"log"
	"main/apikey"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

func main() {

	apiKeyString := apikey.GetGoogleGenAIAPIKey()

	// Ready to bring the model
	context := context.Background()
	fmt.Printf("Obtained an API KEY: %s\n", apiKeyString)
	client, err := genai.NewClient(context, option.WithAPIKey(apiKeyString))
	if err != nil {
		log.Fatal(err)
		return
	}
	defer client.Close()

	model := client.GenerativeModel("models/gemini-pro")

	fmt.Printf("Model obtained: %v\n", model)

}
