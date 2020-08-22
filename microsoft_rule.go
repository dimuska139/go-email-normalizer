package emailnormalizer

import "strings"

// MicrosoftRule
type MicrosoftRule struct {
}

func (rule *MicrosoftRule) processUsername(username string) string {
	return strings.Replace(username, "+", "", -1)
}

func (rule *MicrosoftRule) processDomain(domain string) string {
	return domain
}
