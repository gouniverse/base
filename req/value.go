package req

import "net/http"

// Value returns a POST or GET key, or empty string if not exists
//
// Parameters:
//  - r *http.Request: HTTP request
//  - key string: key to get value for
//
// Returns:
//  - string: value for key, or empty string if not exists
func Value(r *http.Request, key string) string {
	postValue := r.FormValue(key)

	if len(postValue) > 0 {
		return postValue
	}

	getValue := r.URL.Query().Get(key)

	if len(getValue) > 0 {
		return getValue
	}

	return ""
}
