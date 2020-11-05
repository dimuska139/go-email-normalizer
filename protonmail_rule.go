package emailnormalizer

import "strings"

// ProtonmailRule : email normalization rule for Protonmail
type ProtonmailRule struct {
}

func (rule *ProtonmailRule) ProcessUsername(username string) string {
	result := strings.ToLower(username)
	return strings.Replace(result, "+", "", -1)
}

func (rule *ProtonmailRule) ProcessDomain(domain string) string {
	return domain
}
