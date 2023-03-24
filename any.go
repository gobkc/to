package to

import (
	"fmt"
	"reflect"
)

func Any[ToType int | int64 | int32 | string | float32 | float64](dest any) (t ToType) {
	kind := reflect.TypeOf(t).Kind()
	destKind := reflect.TypeOf(dest).Kind()
	switch kind {
	case reflect.Int, reflect.Int32, reflect.Int64:
		if find, ok := intMap[destKind]; ok {
			v := find(dest)
			return ToType(v)
		}
	case reflect.String:
		reflect.ValueOf(&t).Elem().Set(reflect.ValueOf(fmt.Sprintf(`%v`, dest)))
		return
	case reflect.Float32:
		if find, ok := floatMap[destKind]; ok {
			v := find(dest)
			reflect.ValueOf(&t).Elem().Set(reflect.ValueOf(float32(v)))
			return
		}
	case reflect.Float64:
		if find, ok := floatMap[destKind]; ok {
			v := find(dest)
			reflect.ValueOf(&t).Elem().Set(reflect.ValueOf(v))
			return
		}
	}
	return
}
