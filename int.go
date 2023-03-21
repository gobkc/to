package to

import (
	"fmt"
	"reflect"
	"strconv"
)

var intMap = map[reflect.Kind]func(v any) int64{
	reflect.Int: func(v any) int64 {
		value, ok := v.(int)
		if ok {
			return int64(value)
		}
		return 0
	},
	reflect.Int32: func(v any) int64 {
		value, ok := v.(int32)
		if ok {
			return int64(value)
		}
		return 0
	},
	reflect.Int64: func(v any) int64 {
		value, ok := v.(int64)
		if ok {
			return value
		}
		return 0
	},
	reflect.String: func(v any) int64 {
		value, ok := v.(string)
		if ok {
			i, _ := strconv.ParseInt(value, 10, 64)
			return i
		}
		return 0
	},
	reflect.Float64: func(v any) int64 {
		i, _ := strconv.ParseInt(fmt.Sprintf(`%.0f`, v), 10, 64)
		return i
	},
	reflect.Bool: func(v any) int64 {
		value, ok := v.(bool)
		if ok && value {
			return 1
		}
		return 0
	},
}

func Int[ToType int | int64 | int32](dest any) ToType {
	kind := reflect.TypeOf(dest).Kind()
	if find, ok := intMap[kind]; ok {
		v := find(dest)
		return ToType(v)
	}
	return ToType(0)
}
