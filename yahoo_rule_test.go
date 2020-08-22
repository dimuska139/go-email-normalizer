package emailNormalizer

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestYahooUsername(t *testing.T) {
	rule := YahooRule{}
	assert.Equal(t, "t+est", rule.processUsername("t+.-est"))
}

func TestYahooDomain(t *testing.T) {
	rule := YahooRule{}
	assert.Equal(t, "yahoodns.net", rule.processDomain("yahoodns.net"))
}
