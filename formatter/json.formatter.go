package formatter

import (
	"encoding/json"
)

func FormatJson(text string) (string, error) {
	var obj any
	if err := json.Unmarshal([]byte(text), &obj); err != nil {
		return "", err
	}

	formattedJson, err := json.MarshalIndent(obj, "", "  ")
	if err != nil {
		return "", err
	}

	return string(formattedJson), nil
}
