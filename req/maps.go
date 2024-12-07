package req

import (
	"net/http"
	"strings"

	"github.com/samber/lo"
)

func Maps(r *http.Request, key string, defaultValue []map[string]string) []map[string]string {
	all := All(r)

	reqArrayOfMaps := []map[string]string{}

	if all == nil {
		return reqArrayOfMaps
	}

	mapIndexMap := map[string]map[string]string{}

	// Iterate through all the parameters
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
		mapValue := v[0]

		str := strings.TrimSuffix(strings.TrimPrefix(k, key+"["), "]")
		split := strings.Split(str, "][")
		if len(split) != 2 {
			// Handle invalid format
			continue
		}

		index, key := split[0], split[1]

		if lo.HasKey(mapIndexMap, index) {
			mapIndexMap[index][key] = mapValue
		} else {
			mapIndexMap[index] = map[string]string{
				key: mapValue,
			}
		}
	}

	for _, v := range mapIndexMap {
		if v == nil {
			continue
		}
		reqArrayOfMaps = append(reqArrayOfMaps, v)
	}

	return reqArrayOfMaps
}
