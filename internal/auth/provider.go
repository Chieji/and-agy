package auth

import (
	"context"
)

// Provider interface for all AI providers
type Provider interface {
	// Name returns the provider name
	Name() string
	
	// Authenticate authenticates with the provider
	Authenticate(ctx context.Context, config map[string]string) error
	
	// IsAuthenticated checks if authentication is valid
	IsAuthenticated() bool
	
	// Logout clears authentication
	Logout()
	
	// GetConfig returns the current configuration
	GetConfig() map[string]string
}

// ProviderConfig represents configuration for a provider
type ProviderConfig struct {
	Name     string
	Type     string // "api_key", "oauth", "remote"
	Config   map[string]string
	IsDefault bool
}

// ProviderFactory creates providers based on configuration
func ProviderFactory(providerType string, config map[string]string) (Provider, error) {
	switch providerType {
	case "gemini":
		return NewGeminiAuth(), nil
	// TODO: Add other providers
	case "openai":
		// return NewOpenAIAuth(), nil
		return nil, nil
	case "anthropic":
		// return NewAnthropicAuth(), nil
		return nil, nil
	case "deepseek":
		// return NewDeepSeekAuth(), nil
		return nil, nil
	case "xai":
		// return NewXAIAuth(), nil
		return nil, nil
	case "ollama":
		// return NewOllamaAuth(), nil
		return nil, nil
	case "lmstudio":
		// return NewLMStudioAuth(), nil
		return nil, nil
	case "llamacpp":
		// return NewLlamaCppAuth(), nil
		return nil, nil
	default:
		return nil, nil
	}
}