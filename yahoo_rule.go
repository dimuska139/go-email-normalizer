package emailnormalizer

import "strings"

// Rule for Yahoo
type YahooRule struct {
}

func (rule *YahooRule) processUsername(username string) string {
	result := strings.Replace(username, ".", "", -1)
	return strings.Replace(result, "-", "", -1)
}

func (rule *YahooRule) processDomain(domain string) string {
	return domain
}
