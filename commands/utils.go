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

func RecoverAndLog() {
	err := recover()
	if err != nil {
		fmt.Println("RecoverAndLog found error:", err)
	}
}

func ExecuteSafely(callback func() error) {
	defer RecoverAndLog()
	err := callback()

	if err != nil {
		fmt.Println("ExecuteSafely found error:", err)
	}
}
