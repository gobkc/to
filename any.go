package to

import (
	"fmt"
	"reflect"
)

func Any[ToType int | int64 | int32 | string | float32 | float64](dest any) (t ToType) {
	switch kind := reflect.TypeOf(dest).Kind(); kind {
	case reflect.Int, reflect.Int32, reflect.Int64:
		if find, ok := intMap[kind]; ok {
			v := find(dest)
			return ToType(v)
		}
	case reflect.String:
		reflect.ValueOf(t).SetString(fmt.Sprintf(`%v`, dest))
		return
	case reflect.Float32:
		if find, ok := floatMap[kind]; ok {
			v := find(dest)
			reflect.ValueOf(t).Set(reflect.ValueOf(float32(v)))
			return
		}
	case reflect.Float64:
		if find, ok := floatMap[kind]; ok {
			v := find(dest)
			reflect.ValueOf(t).Set(reflect.ValueOf(v))
			return
		}
	}
	return
}
