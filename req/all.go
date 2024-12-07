package req

import (
	"net/http"
	"net/url"
)

// All returns all request variables
func All(r *http.Request) url.Values {
	r.ParseForm()
	return r.Form
}
