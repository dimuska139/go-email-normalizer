package emailNormalizer

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFastmailUsername(t *testing.T) {
	rule := FastmailRule{}
	assert.Equal(t, "t.est", rule.processUsername("t+.est"))
}

func TestFastmailDomain(t *testing.T) {
	rule := FastmailRule{}
	assert.Equal(t, "fastmail.com", rule.processDomain("fastmail.com"))
	assert.Equal(t, "messagingengine.com", rule.processDomain("messagingengine.com"))
}
