package emailnormalizer

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestYahooUsername(t *testing.T) {
	rule := YahooRule{}
	assert.Equal(t, "johnbrown", rule.ProcessUsername("JohnBrown"))
}

func TestYahooDomain(t *testing.T) {
	rule := YahooRule{}
	assert.Equal(t, "yahoodns.net", rule.ProcessDomain("yahoodns.net"))
}
