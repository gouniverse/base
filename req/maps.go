package req

import (
	"net/http"
	"net/url"
	"strings"

	"github.com/samber/lo"
)

func Maps(r *http.Request, key string, defaultValue []map[string]string) []map[string]string {
	all := All(r)

	if all == nil {
		return []map[string]string{}
	}

	keyEntries := filterKeyEntries(all, key)

	if keyEntries == nil {
		return defaultValue
	}

	if len(keyEntries) < 1 {
		return defaultValue
	}

	keys := lo.Keys(keyEntries)
	lenValues := len(keyEntries[keys[0]])

	reqArrayOfMaps := []map[string]string{}

	for i := 0; i < lenValues; i++ {
		m := map[string]string{}
		for k, v := range keyEntries {
			m[k] = v[i]
		}

		reqArrayOfMaps = append(reqArrayOfMaps, m)
	}

	return reqArrayOfMaps
}

func filterKeyEntries(all url.Values, key string) map[string][]string {
	filtered := map[string][]string{}

	for k, v := range all {
		if !strings.HasPrefix(k, key+"[") {
			continue
		}

		if !strings.HasSuffix(k, "]") {
			continue
		}

		if !strings.Contains(k, "][") {
			continue
		}

		str := strings.TrimSuffix(strings.TrimPrefix(k, key+"["), "]")

		split := strings.Split(str, "][")

		if len(split) != 2 {
			// Handle invalid format
			continue
		}

		key, _ := split[0], split[1]

		filtered[key] = v
	}

	return filtered
}
