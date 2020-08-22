package emailnormalizer

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGoogleUsername(t *testing.T) {
	rule := GoogleRule{}
	assert.Equal(t, "test", rule.processUsername("t+.est"))
}

func TestGoogleDomain(t *testing.T) {
	rule := GoogleRule{}
	assert.Equal(t, "gmail.com", rule.processDomain("googlemail.com"))
	assert.Equal(t, "gmail.com", rule.processDomain("gmail.com"))
}
