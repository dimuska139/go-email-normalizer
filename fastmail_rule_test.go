package emailnormalizer

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFastmailUsername(t *testing.T) {
	rule := FastmailRule{}
	assert.Equal(t, "johnbrown", rule.ProcessUsername("JohnBrown+test"))
}

func TestFastmailDomain(t *testing.T) {
	rule := FastmailRule{}
	assert.Equal(t, "fastmail.com", rule.ProcessDomain("fastmail.com"))
	assert.Equal(t, "messagingengine.com", rule.ProcessDomain("messagingengine.com"))
}
