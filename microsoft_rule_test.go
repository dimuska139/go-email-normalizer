package emailnormalizer

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMicrosoftUsername(t *testing.T) {
	rule := MicrosoftRule{}
	assert.Equal(t, "t.est", rule.ProcessUsername("t+.est"))
}

func TestMicrosoftDomain(t *testing.T) {
	rule := MicrosoftRule{}
	assert.Equal(t, "microsoft.com", rule.ProcessDomain("microsoft.com"))
	assert.Equal(t, "live.com", rule.ProcessDomain("live.com"))
}
