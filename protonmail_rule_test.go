package emailnormalizer

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestProtonmailUsername(t *testing.T) {
	rule := ProtonmailRule{}
	assert.Equal(t, "t", rule.ProcessUsername("t+.est"))
}

func TestProtonmailDomain(t *testing.T) {
	rule := ProtonmailRule{}
	assert.Equal(t, "protonmail.ch", rule.ProcessDomain("protonmail.ch"))
}
