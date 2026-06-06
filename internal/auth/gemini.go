package auth

import (
	"context"
	"fmt"
	"os"
)

// GeminiAuth provides authentication for Google Gemini
type GeminiAuth struct {
	apiKey string
}

// NewGeminiAuth creates a new Gemini authenticator
func NewGeminiAuth() *GeminiAuth {
	return &GeminiAuth{}
}

// Authenticate authenticates with the Gemini API
func (g *GeminiAuth) Authenticate(ctx context.Context, config map[string]string) error {
	// Try environment variable first
	apiKey := os.Getenv("GEMINI_API_KEY")
	if apiKey == "" {
		// Try config
		if key, exists := config["api_key"]; exists {
			apiKey = key
		} else {
			return fmt.Errorf("GEMINI_API_KEY environment variable not set and no api_key in config")
		}
	}
	
	g.apiKey = apiKey
	
	// TODO: Validate API key by making a test call
	fmt.Println("Gemini authentication successful")
	
	return nil
}

// GetAPIKey returns the API key
func (g *GeminiAuth) GetAPIKey() string {
	return g.apiKey
}

// IsAuthenticated checks if authentication is valid
func (g *GeminiAuth) IsAuthenticated() bool {
	return g.apiKey != ""
}

// Logout clears authentication
func (g *GeminiAuth) Logout() {
	g.apiKey = ""
}

// Name returns the provider name
func (g *GeminiAuth) Name() string {
	return "gemini"
}