package emailnormalizer

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRamblerUsername(t *testing.T) {
	rule := RamblerRule{}
	assert.Equal(t, "t.est", rule.ProcessUsername("t+.est"))
}

func TestRamblerDomain(t *testing.T) {
	rule := RamblerRule{}
	assert.Equal(t, "lenta.ru", rule.ProcessDomain("lenta.ru"))
	assert.Equal(t, "rambler.ru", rule.ProcessDomain("rambler.ru"))
}
