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
	rules := make(map[string]NormalizingRule)

	microsoftRule := &MicrosoftRule{}
	for _, domain := range microsoftDomains {
		rules[domain] = microsoftRule
	}

	yahooRule := &YahooRule{}
	for _, domain := range yahooDomains {
		rules[domain] = yahooRule
	}

	googleRule := &GoogleRule{}
	for _, domain := range googleDomains {
		rules[domain] = googleRule
	}

	fastmailRule := &FastmailRule{}
	for _, domain := range fastmailDomains {
		rules[domain] = fastmailRule
	}

	ramblerRule := &RamblerRule{}
	for _, domain := range ramblerDomains {
		rules[domain] = ramblerRule
	}

	yandexRule := &YandexRule{}
	for _, domain := range yandexDomains {
		rules[domain] = yandexRule
	}

	protonmailRule := &ProtonmailRule{}
	for _, domain := range protonmailDomains {
		rules[domain] = protonmailRule
	}

	appleRule := &AppleRule{}
	for _, domain := range appleDomains {
		rules[domain] = appleRule
	}

	rules["emailsrvr.com"] = &RackspaceRule{}
	rules["zoho.com"] = &ZohoRule{}
	return &Normalizer{rules: rules}
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

	username := parts[0]                // The first part of the address may be case sensitive (RFC 5336)
	domain := strings.ToLower(parts[1]) // Domain names are case-insensitive (RFC 4343)

	if rule, ok := n.rules[domain]; ok {
		return rule.ProcessUsername(username) + "@" + rule.ProcessDomain(domain)
	}

	return username + "@" + domain
}
