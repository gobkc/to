package to

import (
	"encoding/json"
	"fmt"
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
