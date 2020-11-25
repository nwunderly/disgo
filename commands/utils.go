package commands

import (
	"fmt"
	"strings"
)

func hasPrefix(str string, prefix string, caseInsensitive bool) (bool, string) {
	if caseInsensitive {
		if strings.HasPrefix(strings.ToLower(str), strings.ToLower(prefix)) {
			return true, str[len(prefix):]
		}
	}
	if strings.HasPrefix(str, prefix) {
		return true, str[len(prefix):]
	}
	return false, str
}

func ExecuteSafely() {
	err := recover()
	if err != nil {
		fmt.Println("recovered in ExecuteSafely, found error:", err)
	}
}
