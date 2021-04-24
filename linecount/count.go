package linecount

import (
	"strings"
)

func Count(file []byte) int {
	if len(file) == 0 {
		return 0
	}
	str := string(file)
	str = strings.TrimSuffix(str, "\n")
	result := strings.Split(str, "\n")
	return len(result)
}
