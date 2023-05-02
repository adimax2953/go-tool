package gotool

import (
	jsoniter "github.com/json-iterator/go"
)

// 解Json
func JsonUnmarshal(data []byte, v interface{}) error {
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	if err := json.Unmarshal(data, v); err != nil {
		return err
	}
	return nil
}

// 轉Json
func JsonMarshal(v interface{}) ([]byte, error) {
	var json = jsoniter.ConfigCompatibleWithStandardLibrary

	body, err := json.Marshal(v)
	if err != nil {
		return nil, err
	}
	return body, nil
}
