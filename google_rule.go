package emailnormalizer

import "strings"

// GoogleRule : email normalization rule for Google
type GoogleRule struct {
}

func (rule *GoogleRule) processUsername(username string) string {
	result := strings.Replace(username, ".", "", -1)
	return strings.Replace(result, "+", "", -1)
}

func (rule *GoogleRule) processDomain(domain string) string {
	return "gmail.com" // googlemail.com/gmail.com => gmail.com
}
