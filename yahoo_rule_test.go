package emailnormalizer

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestYahooUsername(t *testing.T) {
	rule := YahooRule{}
	assert.Equal(t, "johnbrown", rule.ProcessUsername("JohnBrown"))
	assert.Equal(t, "john+brown", rule.ProcessUsername("John+Brown"))
	assert.Equal(t, "john", rule.ProcessUsername("John-Brown"))
}

func TestYahooDomain(t *testing.T) {
	rule := YahooRule{}
	assert.Equal(t, "yahoodns.net", rule.ProcessDomain("yahoodns.net"))
}
