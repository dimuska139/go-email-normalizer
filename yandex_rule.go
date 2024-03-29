package emailnormalizer

import "strings"

// YandexRule : email normalization rule for Yandex
type YandexRule struct {
}

func (rule *YandexRule) ProcessUsername(username string) string {
	result := strings.ToLower(username)
	result = strings.Replace(result, "+", "", -1)
	return strings.Replace(result, "-", ".", -1)
}

func (rule *YandexRule) ProcessDomain(domain string) string {
	return "yandex.ru"
}
