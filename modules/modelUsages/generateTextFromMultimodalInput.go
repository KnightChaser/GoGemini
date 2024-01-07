package modelUsages

import (
	"context"
	"fmt"
	"log"
	"main/modules/apikey"
	"os"
	"path/filepath"
	"strings"

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
	prompt := []genai.Part{}
	for _, imagePath := range inputImagePathArray {
		// Trying to open images
		imageData, error := os.ReadFile(imagePath)
		if error != nil {
			log.Panic(error)
		}

		// Valid extension check. "/path/fo/file.png" => "png" (Even remove dot(.))
		imageFileExtension := strings.ToLower(filepath.Ext(imagePath)[1:])
		if !(imageFileExtension == "png" ||
			imageFileExtension == "jpg" ||
			imageFileExtension == "jpeg" ||
			imageFileExtension == "webp" ||
			imageFileExtension == "heic" ||
			imageFileExtension == "heif") {
			errorMessage := fmt.Sprintf("The allowed file formats are only png, jp(e)g, webp, heic, heif, but received %s extension\n", imageFileExtension)
			// Actually, JPG is the same with JPEG. In official docs of Gemini API,
			// jpeg is only supported. (However, JPEG is also acceptable, so just consider JPG as JPEG)
			if imageFileExtension == "jpg" {
				imageFileExtension = "jpeg"
			}
			log.Panic(errorMessage)
		}
		imageDataPart := genai.ImageData(imageFileExtension, imageData)
		prompt = append(prompt, imageDataPart)
	}
	prompt = append(prompt, genai.Text(question))

	response, error := model.GenerateContent(context, prompt...)
	if error != nil {
		log.Panic(error)
	}

	return response

}
