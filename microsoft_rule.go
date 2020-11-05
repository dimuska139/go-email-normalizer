package emailnormalizer

import "strings"

// MicrosoftRule : email normalization rule for Microsoft
type MicrosoftRule struct {
}

func (rule *MicrosoftRule) ProcessUsername(username string) string {
	result := strings.ToLower(username)
	return strings.Replace(result, "+", "", -1)
}

func (rule *MicrosoftRule) ProcessDomain(domain string) string {
	return domain
}
