package runner

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestResolveProviderAPIKey_PrefersAnthropicOAuthToken(t *testing.T) {
	t.Setenv("ANTHROPIC_OAUTH_TOKEN", "oauth-token")
	t.Setenv("ANTHROPIC_API_KEY", "api-token")
	t.Setenv("DSPY_API_KEY", "dspy-token")

	key := ResolveProviderAPIKey("anthropic")
	assert.Equal(t, "oauth-token", key)
}

func TestResolveProviderAPIKey_DefaultIncludesAnthropicOAuthToken(t *testing.T) {
	t.Setenv("ANTHROPIC_OAUTH_TOKEN", "oauth-token")
	t.Setenv("ANTHROPIC_API_KEY", "")
	t.Setenv("DSPY_API_KEY", "")
	t.Setenv("GEMINI_API_KEY", "")
	t.Setenv("GOOGLE_API_KEY", "")
	t.Setenv("OPENAI_API_KEY", "")

	key := ResolveProviderAPIKey("unknown")
	assert.Equal(t, "oauth-token", key)
}
