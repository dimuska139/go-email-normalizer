package emailnormalizer

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestYandexUsername(t *testing.T) {
	rule := YandexRule{}
	assert.Equal(t, "t.est", rule.ProcessUsername("t+.est"))
	assert.Equal(t, "t.est", rule.ProcessUsername("t-est"))
}

func TestYandexDomain(t *testing.T) {
	rule := YandexRule{}
	assert.Equal(t, "yandex.ru", rule.ProcessDomain("yandex.ru"))
	assert.Equal(t, "yandex.ru", rule.ProcessDomain("ya.ru"))
	assert.Equal(t, "yandex.ru", rule.ProcessDomain("narod.ru"))
	assert.Equal(t, "yandex.ru", rule.ProcessDomain("yandex.com"))
	assert.Equal(t, "yandex.ru", rule.ProcessDomain("yandex.by"))
	assert.Equal(t, "yandex.ru", rule.ProcessDomain("yandex.ua"))
	assert.Equal(t, "yandex.ru", rule.ProcessDomain("yandex.kz"))
}
