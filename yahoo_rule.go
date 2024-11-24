package emailnormalizer

import "strings"

// YahooRule : email normalization rule for Yahoo
type YahooRule struct {
}

func (rule *YahooRule) ProcessUsername(username string) string {
	return strings.ToLower(username)
}

func (rule *YahooRule) ProcessDomain(domain string) string {
	return domain
}
