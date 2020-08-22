package email_normalizer

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestZohoUsername(t *testing.T) {
	rule := ZohoRule{}
	assert.Equal(t, "t.est", rule.processUsername("t+.est"))
}

func TestZohoDomain(t *testing.T) {
	rule := ZohoRule{}
	assert.Equal(t, "zoho.com", rule.processDomain("zoho.com"))
}
