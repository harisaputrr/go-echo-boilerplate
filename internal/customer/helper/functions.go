package helpers

import (
	"encoding/json"
	"strings"
)

func GetListSortMapper(payload string) string {
	var sort map[string]string
	if err := json.Unmarshal([]byte(payload), &sort); err != nil {
		return ""
	}

	mappings := map[string]string{
		"name":  "name",
		"email": "email",
	}

	var result []string
	for key, value := range sort {
		if newKey, ok := mappings[key]; ok {
			if value != "asc" && value != "desc" {
				value = "asc"
			}
			result = append(result, newKey+" "+value)
		}
	}

	return strings.Join(result, ", ")
}
