package emailNormalizer

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNormalizer_NormalizeSpaces(t *testing.T) {
	normalizer := NewNormalizer()
	assert.Equal(t, "email@gmail.com", normalizer.Normalize(" email@gmail.com "))
}

func TestNormalizer_NormalizeRightDot(t *testing.T) {
	normalizer := NewNormalizer()
	assert.Equal(t, "email@gmail.com", normalizer.Normalize("email@gmail.com."))
}

func TestNormalizer_NormalizeLower(t *testing.T) {
	normalizer := NewNormalizer()
	assert.Equal(t, "email@gmail.com", normalizer.Normalize("emAil@gMail.cOm."))
}

func TestNormalizer_InvalidEmail(t *testing.T) {
	normalizer := NewNormalizer()
	assert.Equal(t, "emailgmail.com", normalizer.Normalize("emAilgmail.cOm."))
}

func TestNormalizer_UnknownStrategy(t *testing.T) {
	normalizer := NewNormalizer()
	assert.Equal(t, "email@test.com", normalizer.Normalize(" emAil@tEsT.cOm. "))
}

type fakeRule struct {

}

func (rule *fakeRule) processUsername(username string) string {
	return "test"
}

func (rule *fakeRule) processDomain(domain string) string {
	return "email.com"
}

func TestNormalizer_CustomStrategy(t *testing.T) {
	normalizer := NewNormalizer()
	normalizer.AddRule("email.com", &fakeRule{})
	assert.Equal(t, "test@email.com", normalizer.Normalize("user@email.com"))
}