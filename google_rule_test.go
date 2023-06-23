package emailnormalizer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGoogleUsername(t *testing.T) {
	rule := GoogleRule{}
	assert.Equal(t, "t", rule.ProcessUsername("t+.est"))
}

func TestGoogleDomain(t *testing.T) {
	rule := GoogleRule{}
	assert.Equal(t, "gmail.com", rule.ProcessDomain("googlemail.com"))
	assert.Equal(t, "gmail.com", rule.ProcessDomain("gmail.com"))
	assert.Equal(t, "google.com", rule.ProcessDomain("google.com"))
}
