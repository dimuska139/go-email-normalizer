package emailnormalizer

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRackspaceUsername(t *testing.T) {
	rule := RackspaceRule{}
	assert.Equal(t, "t.est", rule.ProcessUsername("t+.est"))
}

func TestRackspaceDomain(t *testing.T) {
	rule := RackspaceRule{}
	assert.Equal(t, "emailsrvr.com", rule.ProcessDomain("emailsrvr.com"))
}
