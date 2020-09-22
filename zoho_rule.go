package emailnormalizer

import "strings"

// ZohoRule : email normalization rule for Zoho
type ZohoRule struct {
}

func (rule *ZohoRule) ProcessUsername(username string) string {
	return strings.Replace(username, "+", "", -1)
}

func (rule *ZohoRule) ProcessDomain(domain string) string {
	return domain
}
