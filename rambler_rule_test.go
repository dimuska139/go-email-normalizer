package emailNormalizer

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRamblerUsername(t *testing.T) {
	rule := RamblerRule{}
	assert.Equal(t, "t.est", rule.processUsername("t+.est"))
}

func TestRamblerDomain(t *testing.T) {
	rule := RamblerRule{}
	assert.Equal(t, "lenta.ru", rule.processDomain("lenta.ru"))
	assert.Equal(t, "rambler.ru", rule.processDomain("rambler.ru"))
}
