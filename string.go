package to

import (
	"encoding/json"
	"fmt"
	"math"
	"regexp"
	"strings"
)

func String(dest any) string {
	return fmt.Sprintf(`%v`, dest)
}

func Json(dest any) string {
	switch v := dest.(type) {
	case string:
		var jsonObj map[string]interface{}
		err := json.Unmarshal([]byte(v), &jsonObj)
		if err != nil {
			return fmt.Sprintf("Error unmarshalling JSON string: %v", err)
		}

		b, err := json.MarshalIndent(jsonObj, "", "\t")
		if err != nil {
			return fmt.Sprintf("Error marshalling JSON: %v", err)
		}
		return string(b)

	case []byte:
		return Json(string(v))

	case nil:
		return ""

	default:
		b, err := json.MarshalIndent(dest, "", "\t")
		if err != nil {
			return fmt.Sprintf("Error marshalling JSON: %v", err)
		}
		return string(b)
	}
}

// FloatDecimals Format Float with decimals
// s := FloatDecimals(3,2); output: 3
// s := FloatDecimals(3.1415,2); output: 3.14
func FloatDecimals[T float64 | float32](f T, dec int) string {
	format := "%.0f"
	if math.Mod(float64(f), 1.0) != 0.0 {
		format = fmt.Sprintf(`%s%d%s`, `%.`, dec, `f`)
	}
	return fmt.Sprintf(format, f)
}

// KebabCase "Hello App" -> "hello-app"
func KebabCase(input string) string {
	re := regexp.MustCompile(`[^a-zA-Z0-9]+`)
	cleaned := re.ReplaceAllString(input, "")
	if len(cleaned) > 0 {
		cleaned = strings.ToLower(string(cleaned[0])) + cleaned[1:]
	}
	re = regexp.MustCompile(`[A-Z]`)
	result := re.ReplaceAllStringFunc(cleaned, func(m string) string {
		return "-" + strings.ToLower(m)
	})

	return result
}
