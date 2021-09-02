package emailnormalizer

import (
	"strings"
)

// Normalizer : main library object for normalization emails
type Normalizer struct {
	rules map[string]NormalizingRule
}

// NewNormalizer : creates Normalizer instance
func NewNormalizer() *Normalizer {
	s := map[string]NormalizingRule{
		"fastmail.com":        &FastmailRule{},
		"messagingengine.com": &FastmailRule{},
		"fastmail.fm":         &FastmailRule{},
		"gmail.com":           &GoogleRule{},
		"googlemail.com":      &GoogleRule{},
		"google.com":          &GoogleRule{},
		"hotmail.com":         &MicrosoftRule{},
		"live.com":            &MicrosoftRule{},
		"outlook.com":         &MicrosoftRule{},
		"rambler.ru":          &RamblerRule{},
		"lenta.ru":            &RamblerRule{},
		"autorambler.ru":      &RamblerRule{},
		"myrambler.ru":        &RamblerRule{},
		"ro.ru":               &RamblerRule{},
		"icloud.com":          &AppleRule{},
		"protonmail.ch":       &ProtonmailRule{},
		"emailsrvr.com":       &RackspaceRule{},
		"yahoodns.net":        &YahooRule{},
		"yahoo.com":           &YahooRule{},
		"ymail.com":           &YahooRule{},
		"yandex.ru":           &YandexRule{},
		"ya.ru":               &YandexRule{},
		"narod.ru":            &YandexRule{},
		"yandex.com":          &YandexRule{},
		"yandex.kz":           &YandexRule{},
		"yandex.ua":           &YandexRule{},
		"yandex.by":           &YandexRule{},
		"zoho.com":            &ZohoRule{},
	}
	return &Normalizer{rules: s}
}

// AddRule : appends custom normalization rule
func (n *Normalizer) AddRule(domain string, strategy NormalizingRule) {
	n.rules[domain] = strategy
}

// Normalize : converts email to canonical form
func (n *Normalizer) Normalize(email string) string {
	prepared := strings.TrimSpace(email)
	prepared = strings.TrimRight(prepared, ".")

	parts := strings.Split(prepared, "@")
	if len(parts) != 2 {
		return prepared
	}

	username := parts[0] // The first part of the address may be case sensitive (RFC 5336)
	domain := strings.ToLower(parts[1]) // Domain names are case-insensitive (RFC 4343)

	if rule, ok := n.rules[domain]; ok {
		return rule.ProcessUsername(username) + "@" + rule.ProcessDomain(domain)
	}

	return username + "@" + domain
}
