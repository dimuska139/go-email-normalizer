package emailnormalizer

// NormalizingRule
type NormalizingRule interface {
	processUsername(string) string
	processDomain(string) string
}
