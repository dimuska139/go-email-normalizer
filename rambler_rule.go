package emailnormalizer

import "strings"

// Rule for Rambler
type RamblerRule struct {
}

func (rule *RamblerRule) processUsername(username string) string {
	return strings.Replace(username, "+", "", -1)
}

func (rule *RamblerRule) processDomain(domain string) string {
	return domain
}
