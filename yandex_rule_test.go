package email_normalizer

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestYandexUsername(t *testing.T) {
	rule := YandexRule{}
	assert.Equal(t, "t.est", rule.processUsername("t+.est"))
}

func TestYandexDomain(t *testing.T) {
	rule := YandexRule{}
	assert.Equal(t, "yandex.ru", rule.processDomain("yandex.ru"))
}
