package emailnormalizer

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestYandexUsername(t *testing.T) {
	rule := YandexRule{}
	assert.Equal(t, "t.est", rule.ProcessUsername("t+.est"))
}

func TestYandexDomain(t *testing.T) {
	rule := YandexRule{}
	assert.Equal(t, "yandex.ru", rule.ProcessDomain("yandex.ru"))
}
