package emailnormalizer

import "strings"

// ProtonmailRule
type ProtonmailRule struct {
}

func (rule *ProtonmailRule) processUsername(username string) string {
	return strings.Replace(username, "+", "", -1)
}

func (rule *ProtonmailRule) processDomain(domain string) string {
	return domain
}
