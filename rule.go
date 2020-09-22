package emailnormalizer

// NormalizingRule : interface for all email normalization rules
type NormalizingRule interface {
	ProcessUsername(string) string
	ProcessDomain(string) string
}
