package to

import (
	"reflect"
	"strconv"
)

var floatMap = map[reflect.Kind]func(v any) float64{
	reflect.Int: func(v any) float64 {
		value, _ := v.(int)
		return float64(value)
	},
	reflect.Int32: func(v any) float64 {
		value, _ := v.(int32)
		return float64(value)
	},
	reflect.Int64: func(v any) float64 {
		value, _ := v.(int64)
		return float64(value)
	},
	reflect.String: func(v any) float64 {
		value, _ := strconv.ParseFloat(v.(string), 64)
		return value
	},
	reflect.Float32: func(v any) float64 {
		return float64(v.(float32))
	},
	reflect.Bool: func(v any) float64 {
		value, ok := v.(bool)
		if ok && value {
			return 1.00
		}
		return 0.00
	},
}

func Float[ToType float32 | float64](dest any) ToType {
	kind := reflect.TypeOf(dest).Kind()
	if find, ok := floatMap[kind]; ok {
		v := find(dest)
		return ToType(v)
	}
	return ToType(0)
}
