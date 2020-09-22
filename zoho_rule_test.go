package emailnormalizer

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestZohoUsername(t *testing.T) {
	rule := ZohoRule{}
	assert.Equal(t, "t.est", rule.ProcessUsername("t+.est"))
}

func TestZohoDomain(t *testing.T) {
	rule := ZohoRule{}
	assert.Equal(t, "zoho.com", rule.ProcessDomain("zoho.com"))
}
