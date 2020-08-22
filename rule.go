package emailnormalizer

// Interface for normalization rules
type NormalizingRule interface {
	processUsername(string) string
	processDomain(string) string
}
