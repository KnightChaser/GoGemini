package main

import (
	"context"
	"fmt"
	"log"
	"main/modules/apikey"

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

	// For text-only input, use the gemini-pro model
	model := client.GenerativeModel("gemini-pro")
	modelInput := "Explain the mechanism of LLM(Large Language Models)s briefly."
	fmt.Printf("Question: %s\n", modelInput)
	response, err := model.GenerateContent(context, genai.Text(modelInput))
	if err != nil {
		log.Fatal(err)
	}

	printGenAIResponse(response)

}

func printGenAIResponse(response *genai.GenerateContentResponse) {
	for _, candidates := range response.Candidates {
		if candidates.Content != nil {
			// response text
			for index, part := range candidates.Content.Parts {
				fmt.Printf("Response(#%d): %v\n", index, part)
			}
			fmt.Println("---------------------------------------------------")

			// metadata
			fmt.Printf("FinishReason: %v\n", candidates.FinishReason)
			fmt.Printf("CitationMetadata: %v\n", candidates.CitationMetadata)
			fmt.Printf("TokenCount: %d\n", candidates.TokenCount)
			for _, safetyData := range candidates.SafetyRatings {
				fmt.Printf("SafetyRatings: %v(%v), Blocked: %v\n", safetyData.Category, safetyData.Probability, safetyData.Blocked)
			}
			fmt.Println("===================================================")
		} else {
			fmt.Println("Failed to receive data from Google Gen. API")
			return
		}
	}
}
