package main

import (
	"encoding/json"
	"fmt"
	"main/modules/modelUsages"
	"main/modules/responseProcess"
)

func main() {
	// ask a question and receive an answer of it
	modelName := "gemini-pro"
	question := "Give me any famous quote from popular people"
	response := modelUsages.GenerateTextFromTextOnlyInput(modelName, question)

	// Procces the data to beautify UwU
	result, _ := responseProcess.GetGeminiAITextOnlyResponseStruct(question, response)
	resultInJSON, _ := json.MarshalIndent(result, "", "	")
	fmt.Println(string(resultInJSON))
}
