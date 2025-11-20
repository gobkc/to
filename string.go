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

// CamelToSnake input a string and convert it to snake_case
// simpleTest=>simple_test
// HTTPRequest=>http_request
func CamelToSnake(s string) string {
	rs := []rune(s)
	n := len(rs)
	var b strings.Builder
	b.Grow(n * 2)

	isAsciiLetter := func(r rune) bool {
		return (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z')
	}

	for i, r := range rs {
		if !isAsciiLetter(r) && r != '_' {
			continue
		}

		if r == '_' {
			b.WriteRune('_')
			continue
		}

		isUpper := (r >= 'A' && r <= 'Z')
		needSep := false
		if isUpper {
			if b.Len() > 0 {
				last := b.String()[b.Len()-1]
				if last == '_' {
					needSep = false
				} else {
					if last >= 'a' && last <= 'z' {
						needSep = true
					} else {
					}
				}
			}
			if !needSep {
				prevIdx := -1
				for j := i - 1; j >= 0; j-- {
					if isAsciiLetter(rs[j]) || rs[j] == '_' {
						prevIdx = j
						break
					}
				}
				nextIdx := -1
				for j := i + 1; j < n; j++ {
					if isAsciiLetter(rs[j]) || rs[j] == '_' {
						nextIdx = j
						break
					}
				}
				if prevIdx != -1 && nextIdx != -1 {
					prev := rs[prevIdx]
					next := rs[nextIdx]
					if isAsciiLetter(prev) && isAsciiLetter(next) {
						if (prev >= 'A' && prev <= 'Z') && (next >= 'a' && next <= 'z') {
							if b.Len() > 0 && b.String()[b.Len()-1] != '_' {
								needSep = true
							}
						}
					}
				}
			}
		}

		if needSep && b.Len() > 0 && b.String()[b.Len()-1] != '_' {
			b.WriteByte('_')
		}

		lower := r
		if isUpper {
			lower = r - 'A' + 'a'
		}
		b.WriteRune(lower)
	}
	return b.String()
}
