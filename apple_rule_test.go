package emailnormalizer

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAppleUsername(t *testing.T) {
	rule := AppleRule{}
	assert.Equal(t, "johnbrown", rule.ProcessUsername("JohnBrown+test"))
}

func TestAppleDomain(t *testing.T) {
	rule := AppleRule{}
	assert.Equal(t, "icloud.com", rule.ProcessDomain("icloud.com"))
	assert.Equal(t, "icloud.com", rule.ProcessDomain("me.com"))
}
