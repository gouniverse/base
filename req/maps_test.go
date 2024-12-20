package req

import (
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

func TestMap_AutoNumbered(t *testing.T) {
	formData := url.Values{
		"map[key1][]": []string{"value10", "value20"},
		"map[key2][]": []string{"value11", "value21"},
	}

	req := httptest.NewRequest("POST", "http://example.com", strings.NewReader(formData.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	result := Maps(req, "map", []map[string]string{})

	if len(result) < 2 {
		t.Error("Array did not return the expected values. Got:", len(result))
	}

	if result[0]["key1"] != "value10" {
		t.Error("Array expected value10. Got:", result[0]["key1"])
	}

	if result[0]["key2"] != "value11" {
		t.Error("Array expected value11. Got:", result[0]["key2"])
	}

	if result[1]["key1"] != "value20" {
		t.Error("Array expected value20. Got:", result[1]["key1"])
	}

	if result[1]["key2"] != "value21" {
		t.Error("Array expected value21. Got:", result[1]["key2"])
	}

}
