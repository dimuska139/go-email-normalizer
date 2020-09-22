package emailnormalizer

import "strings"

// RackspaceRule : email normalization rule for Rackspace
type RackspaceRule struct {
}

func (rule *RackspaceRule) ProcessUsername(username string) string {
	return strings.Replace(username, "+", "", -1)
}

func (rule *RackspaceRule) ProcessDomain(domain string) string {
	return domain
}
