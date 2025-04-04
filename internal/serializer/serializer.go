package serializer

import (
	"encoding/json"
	"github.com/vlbarou/sampleproject/internal/constants"
)

// Marshal marshals any struct
func Marshal(v any) (string, error) {
	jsonData, err := json.Marshal(v)
	if err != nil {
		return "", err
	}

	return string(jsonData), nil
}

// MarshalOrEmpty marshals any struct and returns a json representation of the struct, or empty
func MarshalOrEmpty(v any) (result string) {
	jsonData, err := json.Marshal(v)
	if err != nil {
		result = constants.EmptyString
	}

	result = string(jsonData)
	return
}

// UnmarshalJSON unmarshals JSON into any struct
func UnmarshalJSON(jsonStr string, v any) error {
	return json.Unmarshal([]byte(jsonStr), v)
}
