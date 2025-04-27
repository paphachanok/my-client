package main

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/paphachanok/modelgene/pkg/client"
	"github.com/paphachanok/modelgene/pkg/types"
)

func main() {
	// Step 1: Prepare configuration
	cfg := &types.Config{
		OllamaConfig: &types.OllamaConfig{
			BaseURL: "https://ollama.env.connectedtech.dev",
			HTTPClient: &http.Client{
				Timeout: 60 * time.Second,
			},
		},
	}

	// Step 2: Initialize modelgene client
	c, err := client.NewClient(cfg)
	if err != nil {
		panic(fmt.Errorf("failed to initialize client: %w", err))
	}

	// Step 3: Create API request
	req := types.APIRequest{
		Model: "qwq:32b",
		Messages: []types.Message{
			{
				Role:    "user",
				Content: "Hello, can you tell me a joke?",
			},
		},
	}

	// Step 4: Call Chat
	resp, err := c.Chat(context.Background(), types.ProviderOllama, req)
	if err != nil {
		panic(fmt.Errorf("failed to chat: %w", err))
	}

	// Step 5: Print result
	fmt.Println("🤖 Assistant:", resp.Choices[0].Message.Content)
}
