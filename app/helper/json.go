package helper

import "encoding/json"

func JsonDecoded(jsontext string) interface{} {
	var out interface{}
	err := json.Unmarshal([]byte(jsontext), &out)
	if err != nil {
		return nil
	}
	return out
}

func JsonEncoded(data interface{}) string {
	jsontext, err := json.Marshal(data)
	if err != nil {
		return ""
	}
	return string(jsontext)
}
