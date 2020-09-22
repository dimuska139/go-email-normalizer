package emailnormalizer

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAppleUsername(t *testing.T) {
	rule := AppleRule{}
	assert.Equal(t, "t.est", rule.ProcessUsername("t+.est"))
}

func TestAppleDomain(t *testing.T) {
	rule := AppleRule{}
	assert.Equal(t, "icloud.com", rule.ProcessDomain("icloud.com"))
}
