package linecount

import (
	"strings"
)

func Count(file []byte) int {
	str := string(file)
	result := strings.Split(str, "\n")
	return len(result)
}
