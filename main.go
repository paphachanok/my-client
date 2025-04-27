package main

import (
	"context"
	"fmt"
	"net/http"
	"time"

	anthropic_sdk "github.com/anthropics/anthropic-sdk-go"
	"github.com/paphachanok/modelgene/pkg/client"
	"github.com/paphachanok/modelgene/pkg/types"
)

func main() {
	// Step 1: Prepare configuration
	cfg := &types.Config{
		OllamaConfig: &types.OllamaConfig{
			BaseURL: "https://ollama.env.connectedtech.dev/",
			HTTPClient: &http.Client{
				Timeout: 60 * time.Second,
			},
		},
		AnthropicConfig: &types.AnthropicConfig{
			APIKey: "your-anthropic-api-key",
		},
	}

	// Step 2: Initialize modelgene client
	c, err := client.NewClient(cfg)
	if err != nil {
		panic(fmt.Errorf("failed to initialize client: %w", err))
	}

	// Step 3: Create APIRequest to Ollama
	ollamaReq := types.APIRequest{
		Model: "qwq:32b",
		Messages: []types.Message{
			{Role: "user", Content: "Tell me a joke about computers"},
		},
	}

	// Step 4: Create APIRequest to Anthropic
	anthropicReq := types.APIRequest{
		Model: anthropic_sdk.ModelClaude3_7SonnetLatest,
		Messages: []types.Message{
			{Role: "user", Content: "Tell me a joke about computers"},
		},
	}

	// Step 5: Call Chat for Ollama
	ollamaResp, err := c.Chat(context.Background(), types.ProviderOllama, ollamaReq)
	if err != nil {
		fmt.Println("Ollama chat error:", err)
	} else {
		fmt.Println("🤖 Ollama:", ollamaResp.Choices[0].Message.Content)
	}

	// Step 6: Call Chat for Anthropic
	anthropicResp, err := c.Chat(context.Background(), types.ProviderAnthropic, anthropicReq)
	if err != nil {
		fmt.Println("Anthropic chat error:", err)
	} else {
		fmt.Println("🤖 Anthropic:", anthropicResp.Choices[0].Message.Content)
	}
}
