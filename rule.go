package emailnormalizer

// NormalizingRule : interface for all email normalization rules
type NormalizingRule interface {
	processUsername(string) string
	processDomain(string) string
}
