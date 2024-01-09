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
	// modelTextInputTextOutput()
	// modelMultimoalInputTextOutput()
	modelUsages.GenerateTextChatSession("gemini-pro")
	// modelUsages.GenerateTextChatSessionStreaming("gemini-pro")
}

// An example usage of
// - input: text only query
// - output: text only response
func modelTextInputTextOutput() {
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

// An example usage of
// - input: multimodal input(text + image(s))
// - output: text only response
func modelMultimoalInputTextOutput() {
	// model set is fixed to "gemini-pro-vision" automatically
	question := "Describe about the first and second images related to desert"
	imagePathSlice := []string{
		"modules/modelUsages/sampleImages/cheesecake.jpeg",
		"modules/modelUsages/sampleImages/coffee.png",
	}
	response := modelUsages.GenerateTextFromMultimodalInput(question, imagePathSlice)

	// Procces the data to beautify UwU
	result, _ := responseProcess.GetGeminiAITextOnlyResponseStruct(question, response)
	resultInJSON, _ := json.MarshalIndent(result, "", "	")
	fmt.Println(string(resultInJSON))
}
