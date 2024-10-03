```go
package router

import (
	"net/http"
)

// InitRouter initializes and returns a new HTTP ServeMux
func InitRouter() *http.ServeMux {
	r := http.NewServeMux()

	// Here you can initialize middlewares if needed
	// Example: r = loggingMiddleware(r)

	return r
}
```