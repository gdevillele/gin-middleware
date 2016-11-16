package language

import (
	"net/http"

	"golang.org/x/text/language"

	"github.com/gin-gonic/gin"
)

// AcceptLanguages creates a middleware which matches the client languages
// against the supported languages provided. Please note that the first language
// provided is considered the default language.
func AcceptLanguages(supportedLanguages []language.Tag) gin.HandlerFunc {
	// matcher is a language.Matcher configured for all supported languages.
	matcher := language.NewMatcher(supportedLanguages)
	// returns middleware function
	return func(c *gin.Context) {
		t, _, _ := language.ParseAcceptLanguage(c.Request.Header.Get("Accept-Language"))
		tag, _, _ := matcher.Match(t...)
		c.Set("language", tag)
	}
}
