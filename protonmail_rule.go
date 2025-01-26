package emailnormalizer

import "strings"

// ProtonmailRule : email normalization rule for Protonmail
type ProtonmailRule struct {
}

func (rule *ProtonmailRule) ProcessUsername(username string) string {
	result := strings.ToLower(username)

	charsToReplace := []string{
		".",
		"_",
		"-",
	}

	for _, char := range charsToReplace {
		result = strings.Replace(result, char, "", -1)
	}

	plusSignIndex := strings.Index(result, "+")
	if plusSignIndex != -1 {
		result = result[0:plusSignIndex]
	}

	return result
}

func (rule *ProtonmailRule) ProcessDomain(domain string) string {
	return domain
}
