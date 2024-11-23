package utils

import "encoding/json"

func Obj2JsonStr(v any) string {
	body, err := json.Marshal(v)
	if err != nil {
		return ""
	}
	return string(body)
}
