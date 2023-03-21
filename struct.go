package to

import (
	"encoding/json"
	"reflect"
)

func Jsonb[T []any | any](list T) string {
	jByte, _ := json.Marshal(list)
	js := string(jByte)
	if js == "" {
		if reflect.TypeOf(list).Kind() == reflect.Slice {
			js = "[]"
		} else {
			js = "{}"
		}
	}
	return js
}

func Object[Source any, Dest []byte | string](dest Dest) *Source {
	t := new(Source)
	_ = json.Unmarshal([]byte(dest), t)
	return t
}
