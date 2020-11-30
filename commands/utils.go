package commands

import (
	"fmt"
	"reflect"
	"runtime"
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

func GetFunctionName(i interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
}

func deleteCommand(s []*Command, i int) []*Command {
	return append(s[:i], s[i+1:]...)
}