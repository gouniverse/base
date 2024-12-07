package req

import "net/http"

// ValueOr returns a POST or GET key, or provided default value if not exists
//
// Parameters:
//  - r *http.Request: HTTP request
//  - key string: key to get value for
//  - defaultValue string: default value to return if key does not exist
//
// Returns:
//  - string: value for key, or default value if not exists
func ValueOr(r *http.Request, key string, defaultValue string) string {
	postValue := r.FormValue(key)

	if len(postValue) > 0 {
		return postValue
	}

	getValue := r.URL.Query().Get(key)

	if len(getValue) > 0 {
		return getValue
	}

	return defaultValue
}
