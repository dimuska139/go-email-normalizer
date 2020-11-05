package emailnormalizer

import "strings"

// AppleRule : email normalization rule for Apple
type AppleRule struct {
}

func (rule *AppleRule) ProcessUsername(username string) string {
	result := strings.ToLower(username)
	return strings.Replace(result, "+", "", -1)
}

func (rule *AppleRule) ProcessDomain(domain string) string {
	return domain
}
