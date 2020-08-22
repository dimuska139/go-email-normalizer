package email_normalizer

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMicrosoftUsername(t *testing.T) {
	rule := MicrosoftRule{}
	assert.Equal(t, "t.est", rule.processUsername("t+.est"))
}

func TestMicrosoftDomain(t *testing.T) {
	rule := MicrosoftRule{}
	assert.Equal(t, "microsoft.com", rule.processDomain("microsoft.com"))
	assert.Equal(t, "live.com", rule.processDomain("live.com"))
}
