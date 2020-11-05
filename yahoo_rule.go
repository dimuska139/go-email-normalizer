package emailnormalizer

import "strings"

// YahooRule : email normalization rule for Yahoo
type YahooRule struct {
}

func (rule *YahooRule) ProcessUsername(username string) string {
	result := strings.ToLower(username)
	result = strings.Replace(result, ".", "", -1)
	return strings.Replace(result, "-", "", -1)
}

func (rule *YahooRule) ProcessDomain(domain string) string {
	return domain
}
