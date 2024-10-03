```go
package router

import (
	"net/http"
)

// Router holds the routes for the application
type Router struct {
	mux *http.ServeMux
}

// NewRouter creates a new Router instance
func NewRouter() *Router {
	return &Router{
		mux: http.NewServeMux(),
	}
}

// HandleFunc registers a new route with a matcher for the URL path and a handler function
func (r *Router) HandleFunc(pattern string, handler func(http.ResponseWriter, *http.Request)) {
	r.mux.HandleFunc(pattern, handler)
}

// ServeHTTP dispatches the request to the handler whose pattern most closely matches the request URL
func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	r.mux.ServeHTTP(w, req)
}
```