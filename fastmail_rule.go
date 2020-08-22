package email_normalizer

import "strings"

type FastmailRule struct {
}

func (rule *FastmailRule) processUsername(username string) string {
	return strings.Replace(username, "+", "", -1)
}

func (rule *FastmailRule) processDomain(domain string) string {
	return domain
}
