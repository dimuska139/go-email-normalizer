package email_normalizer

type NormalizingRule interface {
	processUsername(string) string
	processDomain(string) string
}
