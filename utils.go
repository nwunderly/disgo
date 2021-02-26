package disgo

import (
	"fmt"
	"reflect"
	"regexp"
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

func ExecuteSafely(callback func() error) (err error) {
	defer func() {
		if e := recover(); e != nil {
			switch x := e.(type) {
			case error:
				err = x
			default:
				err = fmt.Errorf("%s", e)
			}
		}
	}()
	err = callback()
	return
}

func GetFunctionName(i interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
}

func deleteCommand(s []*Command, i int) []*Command {
	return append(s[:i], s[i+1:]...)
}

func GetArgs(s string) []string {
	pattern, _ := regexp.Compile(`("[\w\s]+"|\w+)`)

	args := pattern.FindAllString(s, -1)

	for i, a := range args {
		if strings.HasPrefix(a, `"`) && strings.HasSuffix(a, `"`) {
			args[i] = strings.TrimSuffix(strings.TrimPrefix(a, `"`), `"`)
		}
	}

	return args
}
