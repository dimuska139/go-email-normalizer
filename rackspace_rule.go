package emailnormalizer

import "strings"

// Rule for Rackspace
type RackspaceRule struct {
}

func (rule *RackspaceRule) processUsername(username string) string {
	return strings.Replace(username, "+", "", -1)
}

func (rule *RackspaceRule) processDomain(domain string) string {
	return domain
}
