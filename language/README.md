# language

```
import "github.com/gdevillele/gin-middleware/language"
```

This middleware parses the request's `Accept-Language` HTTP header and matches its content against the supported languages you provide. The language that best matches (or the default language if no match is found) is stored in the `gin.Context` structure at the key `"language"`.

```go
import "golang.org/x/text/language"

func routeHandler(c *gin.Context) {
	var lang language.Tag = c.Get("language")
	
	// ...
}
```
