package emailnormalizer

import "strings"

// FastmailRule : email normalization rule for Fastmail
type FastmailRule struct {
}

func (rule *FastmailRule) ProcessUsername(username string) string {
	result := strings.ToLower(username)
	return strings.Replace(result, "+", "", -1)
}

func (rule *FastmailRule) ProcessDomain(domain string) string {
	return domain
}
