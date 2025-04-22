package main 
import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/paphachanok/coder-gene-test/api"
)

func main() {
	// Create a unified request
	req := modelapi.APIRequest{
		Model: "gpt-4",
		Messages: []modelapi.Message{
			{Role: "system", Content: "You are a helpful assistant."},
			{Role: "user", Content: "Tell me a joke."},
		},
		Temperature: modelapi.PtrFloat64(0.7),
		MaxTokens:   modelapi.PtrInt(100),
		Stream:      modelapi.PtrBool(false),
	}

	// Convert to OpenAI request
	openaiPayload, err := api.ConvertToOpenAI(req)
	if err != nil {
		log.Fatalf("conversion failed: %v", err)
	}

	// Marshal to JSON to simulate API sending
	jsonData, err := json.MarshalIndent(openaiPayload, "", "  ")
	if err != nil {
		log.Fatalf("json marshal failed: %v", err)
	}

	fmt.Println("🧠 OpenAI Payload:\n", string(jsonData))

	// Next step: use an HTTP client to send `jsonData` to OpenAI's API endpoint.
}
