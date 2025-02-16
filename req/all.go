package req

import (
	"net/http"
	"net/url"
)

// All returns all request variables as a url.Values object
//
// Parameters:
//   - r *http.Request: HTTP request
//
// Returns:
//   - url.Values: all request variables
func All(r *http.Request) url.Values {
	r.ParseForm()
	return r.Form
}
