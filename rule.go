package emailNormalizer

type NormalizingRule interface {
	processUsername(string) string
	processDomain(string) string
}
