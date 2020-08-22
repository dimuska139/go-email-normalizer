package email_normalizer

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestProtonmailUsername(t *testing.T) {
	rule := ProtonmailRule{}
	assert.Equal(t, "t.est", rule.processUsername("t+.est"))
}

func TestProtonmailDomain(t *testing.T) {
	rule := ProtonmailRule{}
	assert.Equal(t, "protonmail.ch", rule.processDomain("protonmail.ch"))
}
