package api

import "encoding/json"

func EncodeJSONString(v any) (string, error) {
	data, err := json.Marshal(v)
	if err != nil {
		return "", err
	}

	return string(data), nil
}

func EncodeJSON(v any) ([]byte, error) {
	data, err := json.Marshal(v)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func ParseJSON(data []byte, v any) error {
	err := json.Unmarshal(data, v)
	if err != nil {
		return err
	}

	return nil
}
