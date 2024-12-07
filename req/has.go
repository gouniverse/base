package req

import "net/http"

// Has returns true if GET or POST key exists
//
// Parameters:
//  - r *http.Request: HTTP request
//  - key string: key to check if exists
//
// Returns:
//  - bool: true if key exists
func Has(r *http.Request, key string) bool {
	if HasGet(r, key) {
		return true
	}

	return HasPost(r, key)
}

// HasPost returns true if POST key exists
//
// Parameters:
//  - r *http.Request: HTTP request
//  - key string: key to check if exists
//
// Returns:
//  - bool: true if key exists
func HasPost(r *http.Request, key string) bool {
	err := r.ParseForm()

	if err != nil {
		return false
	}

	_, exists := r.Form[key]
	return exists
}

// HasGet returns true if GET key exists
//
// Parameters:
//  - r *http.Request: HTTP request
//  - key string: key to check if exists
//
// Returns:
//  - bool: true if key exists
func HasGet(r *http.Request, key string) bool {
	_, exists := r.URL.Query()[key]
	return exists
}
