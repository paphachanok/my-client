package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"time"

	anthropic_sdk "github.com/anthropics/anthropic-sdk-go"
	"github.com/connectedtechco/modelgene/pkg/client"
	"github.com/connectedtechco/modelgene/pkg/types"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()

	// Step 1: Prepare configuration
	cfg := &types.Config{
		OllamaConfig: &types.OllamaConfig{
			BaseURL: os.Getenv("OLLAMA_BASE_URL"),
			HTTPClient: &http.Client{
				Timeout: 3000 * time.Second,
			},
		},
		AnthropicConfig: &types.AnthropicConfig{
			APIKey: os.Getenv("ANTHROPIC_API_KEY"),
		},
	}

	// Step 2: Initialize modelgene client
	c, err := client.NewClient(cfg)
	if err != nil {
		panic(fmt.Errorf("failed to initialize client: %w", err))
	}

	// Step 3: Create APIRequest to Ollama
	ollamaReq := types.APIRequest{
		Model: "gemma3:12b",
		Messages: []types.Message{
			{Role: "user", Content: "What's color is banana"},
		},
	}

	// Step 4: Create APIRequest to Anthropic
	anthropicReq := types.APIRequest{
		Model: anthropic_sdk.ModelClaude3_7SonnetLatest,
		Messages: []types.Message{
			{Role: "user", Content: "What is collection in non-relational database"},
		},
	}

	// Step 5: Create embed request to Ollama
	embedReq := types.APIRequest{
		Model: "snowflake-arctic-embed2:568m",
		Input: "The quick brown fox jumps over the lazy dog",
	}

	// Step 6: Call Chat for Ollama
	ollamaResp, err := c.Chat(context.Background(), types.ProviderOllama, ollamaReq)
	if err != nil {
		fmt.Println("Ollama chat error:", err)
	} else {
		fmt.Println("🤖 Ollama:", ollamaResp.Choices[0].Message.Content)
	}

	// Step 7: Call Chat for Anthropic
	anthropicResp, err := c.Chat(context.Background(), types.ProviderAnthropic, anthropicReq)
	if err != nil {
		fmt.Println("Anthropic chat error:", err)
	} else {
		fmt.Println("🤖 Anthropic:", anthropicResp.Choices[0].Message.Content)
	}

	// Step 8: Call Embed function
	embedResp, err := c.Embed(context.Background(), types.ProviderOllama, embedReq)
	if err != nil {
		fmt.Println("🔴 Ollama embed error:", err)
		return
	}

	// Step 9: Print result (the embedding vector as a comma-separated string)
	fmt.Println("🧠 Ollama embedding result:")
	for _, choice := range embedResp.Choices {
		fmt.Printf("🔹 Index %d:\n", choice.Index)
		fmt.Printf("   Vector (stringified): %s\n", choice.Message.Content[:80]+"...") // truncated for display
	}
}
