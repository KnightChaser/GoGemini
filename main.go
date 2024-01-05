package main

import (
	"encoding/json"
	"fmt"
	"log"
	"main/modules/modelUsages"
	"main/modules/responseProcess"
)

// main testing
func main() {
	// Prepare the model
	modelName := "gemini-pro"
	if _, error := modelUsages.ModelConnectionTest(modelName); error != nil {
		log.Fatal(fmt.Sprintf("Failed to bring the model \"%v\" from the server\n", modelName))
		return
	}

	// ask a question and receive an answer of it
	question := "Give me any famous quote from popular people"
	response := modelUsages.GenerateTextFromTextOnlyInput(modelName, question)

	// Procces the data to beautify UwU
	result, _ := responseProcess.GetGeminiAITextOnlyResponseStruct(question, response)
	resultInJSON, _ := json.MarshalIndent(result, "", "	")
	fmt.Println(string(resultInJSON))
}
