package emailnormalizer

import "strings"

// RamblerRule : email normalization rule for Rambler
type RamblerRule struct {
}

func (rule *RamblerRule) ProcessUsername(username string) string {
	result := strings.ToLower(username)
	return strings.Replace(result, "+", "", -1)
}

func (rule *RamblerRule) ProcessDomain(domain string) string {
	return domain
}
