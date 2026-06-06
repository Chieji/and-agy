package auth

import (
	"context"
	"fmt"
	"sync"
)

// AuthManager manages multiple AI providers
type AuthManager struct {
	providers map[string]Provider
	configs   map[string]ProviderConfig
	defaultProvider string
	mu        sync.RWMutex
}

// NewAuthManager creates a new auth manager
func NewAuthManager() *AuthManager {
	return &AuthManager{
		providers: make(map[string]Provider),
		configs:   make(map[string]ProviderConfig),
	}
}

// AddProvider adds a provider to the manager
func (am *AuthManager) AddProvider(provider Provider, config ProviderConfig) {
	am.mu.Lock()
	defer am.mu.Unlock()
	
	am.providers[provider.Name()] = provider
	am.configs[provider.Name()] = config
	
	if config.IsDefault {
		am.defaultProvider = provider.Name()
	}
}

// GetProvider returns a provider by name
func (am *AuthManager) GetProvider(name string) (Provider, bool) {
	am.mu.RLock()
	defer am.mu.RUnlock()
	
	provider, exists := am.providers[name]
	return provider, exists
}

// GetDefaultProvider returns the default provider
func (am *AuthManager) GetDefaultProvider() (Provider, bool) {
	am.mu.RLock()
	defer am.mu.RUnlock()
	
	if am.defaultProvider == "" {
		return nil, false
	}
	
	return am.providers[am.defaultProvider], true
}

// ListProviders returns all configured providers
func (am *AuthManager) ListProviders() []string {
	am.mu.RLock()
	defer am.mu.RUnlock()
	
	providers := make([]string, 0, len(am.providers))
	for name := range am.providers {
		providers = append(providers, name)
	}
	
	return providers
}

// Authenticate authenticates with a specific provider
func (am *AuthManager) Authenticate(ctx context.Context, name string, config map[string]string) error {
	am.mu.Lock()
	defer am.mu.Unlock()
	
	provider, exists := am.providers[name]
	if !exists {
		// Create new provider
		provider, err := ProviderFactory(name, config)
		if err != nil {
			return fmt.Errorf("failed to create provider: %w", err)
		}
		am.providers[name] = provider
	}
	
	if err := provider.Authenticate(ctx, config); err != nil {
		return fmt.Errorf("authentication failed: %w", err)
	}
	
	am.configs[name] = ProviderConfig{
		Name:   name,
		Type:   "api_key", // TODO: Detect type
		Config: config,
	}
	
	return nil
}

// SetDefaultProvider sets the default provider
func (am *AuthManager) SetDefaultProvider(name string) error {
	am.mu.Lock()
	defer am.mu.Unlock()
	
	if _, exists := am.providers[name]; !exists {
		return fmt.Errorf("provider not found: %s", name)
	}
	
	am.defaultProvider = name
	
	// Update config
	for k := range am.configs {
		am.configs[k].IsDefault = (k == name)
	}
	
	return nil
}

// Logout logs out from a provider
func (am *AuthManager) Logout(name string) error {
	am.mu.Lock()
	defer am.mu.Unlock()
	
	provider, exists := am.providers[name]
	if !exists {
		return fmt.Errorf("provider not found: %s", name)
	}
	
	provider.Logout()
	
	if am.defaultProvider == name {
		am.defaultProvider = ""
	}
	
	delete(am.providers, name)
	delete(am.configs, name)
	
	return nil
}

// RemoveProvider removes a provider
func (am *AuthManager) RemoveProvider(name string) error {
	return am.Logout(name)
}

// GetConfig returns configuration for a provider
func (am *AuthManager) GetConfig(name string) (ProviderConfig, bool) {
	am.mu.RLock()
	defer am.mu.RUnlock()
	
	config, exists := am.configs[name]
	return config, exists
}

// SaveConfig saves the current configuration
func (am *AuthManager) SaveConfig() (map[string]ProviderConfig, error) {
	am.mu.RLock()
	defer am.mu.RUnlock()
	
	// Return a copy
	configs := make(map[string]ProviderConfig)
	for k, v := range am.configs {
		configs[k] = v
	}
	
	return configs, nil
}

// LoadConfig loads configuration
func (am *AuthManager) LoadConfig(configs map[string]ProviderConfig) {
	am.mu.Lock()
	defer am.mu.Unlock()
	
	for name, config := range configs {
		provider, err := ProviderFactory(config.Type, config.Config)
		if err == nil && provider != nil {
			am.providers[name] = provider
			am.configs[name] = config
			if config.IsDefault {
				am.defaultProvider = name
			}
		}
	}
}