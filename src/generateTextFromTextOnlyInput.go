package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"main/modules/apikey"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

type GeminiTextResponseSafetyRating struct {
	Probability           string
	BlockedBySafetyPolicy bool
}

type GeminiTextResponseStructure struct {
	Question      string
	ResponseCount uint32
	Response      []string
	FininshReason string
	TokenCount    uint32
	SafetyRating  map[string]GeminiTextResponseSafetyRating
}

func main() {
	generateTextFromTextOnlyInput("gemini-pro", "Create a random quote related to programming or programmers")
}

// getAIModelName: The GenAI model to use like "Gemini-Pro"
// question		 : The input query to the model
func generateTextFromTextOnlyInput(genAIModelName string, question string) {

	apiKeyString := apikey.GetGoogleGenAIAPIKey("default")

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
	model := client.GenerativeModel(genAIModelName)
	modelInput := question
	fmt.Printf("Question: %s\n", modelInput)
	response, err := model.GenerateContent(context, genai.Text(modelInput))
	if err != nil {
		log.Fatal(err)
	}

	// printGenAIResponse(response)
	result, _ := getGeminiAITextOnlyResponseStruct(question, response)
	resultInJSON, _ := json.MarshalIndent(result, "", "	")
	fmt.Println(string(resultInJSON))

}

func getGeminiAITextOnlyResponseStruct(question string, response *genai.GenerateContentResponse) (GeminiTextResponseStructure, error) {
	var responseStructure GeminiTextResponseStructure
	responseStructure.Question = question
	for _, candidates := range response.Candidates {
		if candidates.Content != nil {
			// response text
			for _, part := range candidates.Content.Parts {
				responseStructure.ResponseCount += 1
				responseStructure.Response = append(responseStructure.Response, fmt.Sprintf("%s", part))
			}

			// metadata
			responseStructure.FininshReason = fmt.Sprintf("%s", candidates.FinishReason)
			responseStructure.TokenCount = uint32(candidates.TokenCount)
			responseStructure.SafetyRating = make(map[string]GeminiTextResponseSafetyRating)
			for _, safetyData := range candidates.SafetyRatings {
				fmt.Println(safetyData)
				safetyCategory := fmt.Sprintf("%s", safetyData.Category)

				// Check if the map entry exists, create it if not
				if _, ok := responseStructure.SafetyRating[safetyCategory]; !ok {
					responseStructure.SafetyRating[safetyCategory] = GeminiTextResponseSafetyRating{}
				}

				// Create an instance of GeminiTextResponseSafetyRating
				safetyRating := responseStructure.SafetyRating[safetyCategory]

				// Assign values to the struct fields
				safetyRating.BlockedBySafetyPolicy = safetyData.Blocked
				safetyRating.Probability = fmt.Sprintf("%s", safetyData.Probability)

				// Update the map entry
				responseStructure.SafetyRating[safetyCategory] = safetyRating
			}

		} else {
			return GeminiTextResponseStructure{}, fmt.Errorf("Failed to receive data from Google Gen. API")
		}
	}
	return responseStructure, nil
}
