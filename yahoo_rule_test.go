package emailnormalizer

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestYahooUsername(t *testing.T) {
	rule := YahooRule{}
	assert.Equal(t, "t+.est", rule.ProcessUsername("t+.-est"))
}

func TestYahooDomain(t *testing.T) {
	rule := YahooRule{}
	assert.Equal(t, "yahoodns.net", rule.ProcessDomain("yahoodns.net"))
}
