package emailnormalizer

import "strings"

// FastmailRule : email normalization rule for Fastmail
type FastmailRule struct {
}

func (rule *FastmailRule) ProcessUsername(username string) string {
	result := strings.ToLower(username)

	// Remove sub-addressing part (RFC 5233)
	plusSignIndex := strings.Index(result, "+")
	if plusSignIndex != -1 {
		result = result[0:plusSignIndex]
	}

	return result
}

func (rule *FastmailRule) ProcessDomain(domain string) string {
	return domain
}
