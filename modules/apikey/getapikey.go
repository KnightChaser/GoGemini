package apikey

import (
	"fmt"
	"log"
	"os"
)

// Grab the API key for the service usage, reading API key from the designated file.
func GetGoogleGenAIAPIKey(apiKeyFilepath string) string {
	// Setup API key (the given API file will contain text of API key)

	var geminiAPIkeyFilePath string
	if apiKeyFilepath == "default" {
		// default
		geminiAPIkeyFilePath = "modules/apikey/apikey.txt"
	} else {
		geminiAPIkeyFilePath = apiKeyFilepath
	}

	geminiAPIKey, err := os.ReadFile(geminiAPIkeyFilePath)
	if err != nil {
		fmt.Printf("Failed to load the API key (Expected filepath: %v)\n", geminiAPIkeyFilePath)
		log.Fatal(err)
		return "ERR"
	}

	// Convert geminiAPIKey from []byte to string
	apiKeyString := string(geminiAPIKey)

	return apiKeyString
}
