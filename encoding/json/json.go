package json

import "encoding/json"

// JSONString converts given structure to JSON string.
func JSONString(params interface{}) (string, error) {
	b, err := json.Marshal(params)
	if err != nil {
		return "", err
	}
	return string(b), nil
}
