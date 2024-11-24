package emailnormalizer

import "strings"

// AppleRule : email normalization rule for Apple
type AppleRule struct {
}

func (rule *AppleRule) ProcessUsername(username string) string {
	result := strings.ToLower(username)
	// Remove sub-addressing part (RFC 5233)
	plusSignIndex := strings.Index(result, "+")
	if plusSignIndex != -1 {
		result = result[0:plusSignIndex]
	}

	return result
}

func (rule *AppleRule) ProcessDomain(domain string) string {
	return "icloud.com"
}
