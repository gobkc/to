package to

import (
	"encoding/json"
	"fmt"
)

func String(dest any) string {
	return fmt.Sprintf(`%v`, dest)
}

func Json(dest any) string {
	b, _ := json.MarshalIndent(dest, "", "\t")
	return string(b)
}
