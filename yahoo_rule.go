package emailnormalizer

import "strings"

// YahooRule : email normalization rule for Yahoo
type YahooRule struct {
}

func (rule *YahooRule) ProcessUsername(username string) string {
	result := strings.ToLower(username)

	subaddressingIndex := strings.Index(result, "-")
	if subaddressingIndex != -1 {
		result = result[0:subaddressingIndex]
	}

	return result
}

func (rule *YahooRule) ProcessDomain(domain string) string {
	return domain
}
