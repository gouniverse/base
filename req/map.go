package req

import (
	"net/http"
	"strings"
)

// Map returns a map from the request
//
// Parameters:
//   - r *http.Request: HTTP request
//   - key string: key to get map for
//
// Returns:
//   - map[string]string: map for key
func Map(r *http.Request, key string) map[string]string {
	all := All(r)

	reqMap := map[string]string{}

	if all == nil {
		return reqMap
	}

	for k, v := range all {
		if strings.HasPrefix(k, key+"[") && strings.HasSuffix(k, "]") {
			if len(v) < 1 {
				reqMap[strings.TrimSuffix(strings.TrimPrefix(k, key+"["), "]")] = ""
				continue
			}

			reqMap[strings.TrimSuffix(strings.TrimPrefix(k, key+"["), "]")] = v[0]
		}
	}

	return reqMap
}
