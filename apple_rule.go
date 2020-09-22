package emailnormalizer

import "strings"

// AppleRule : email normalization rule for Apple
type AppleRule struct {
}

func (rule *AppleRule) ProcessUsername(username string) string {
	return strings.Replace(username, "+", "", -1)
}

func (rule *AppleRule) ProcessDomain(domain string) string {
	return domain
}
