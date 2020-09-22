package emailnormalizer

import "strings"

// GoogleRule : email normalization rule for Google
type GoogleRule struct {
}

func (rule *GoogleRule) ProcessUsername(username string) string {
	result := strings.Replace(username, ".", "", -1)

	plusSignIndex := strings.Index(result, "+")
	if plusSignIndex != -1 {
		result = result[0:plusSignIndex]
	}

	return result
}

func (rule *GoogleRule) ProcessDomain(domain string) string {
	return "gmail.com" // googlemail.com/gmail.com => gmail.com
}
