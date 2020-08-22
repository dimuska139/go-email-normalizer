package email_normalizer

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAppleUsername(t *testing.T) {
	rule := AppleRule{}
	assert.Equal(t, "t.est", rule.processUsername("t+.est"))
}

func TestAppleDomain(t *testing.T) {
	rule := AppleRule{}
	assert.Equal(t, "icloud.com", rule.processDomain("icloud.com"))
}
