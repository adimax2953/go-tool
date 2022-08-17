package gotool

import (
	"math/rand"
	"runtime/debug"
	"strings"

	LogTool "github.com/adimax2953/log-tool"
)

// 取亂數 1~num
func RanInt(num int) int {
	if num < 0 {
		LogTool.LogDebug("傳入了負數 %d", num)
		return 0
	}

	if num == 0 {
		num++
	}
	rndInt := rand.Intn(num) + 1
	return rndInt
}

func RecoverPanic() {
	e := recover()
	if e != nil {
		if err, ok := e.(error); ok {
			LogTool.LogError(err.Error())
		} else {
			LogTool.LogError("", e, debug.Stack())
		}
		debug.PrintStack()
		return
	}
}

// 遞增字符串ByBase62
func Base62Increment(s string) string {
	defer RecoverPanic()
	chars := "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

	var firstChar = s[0]
	if strings.Index(chars, string(firstChar)) == 61 {
		return Base62Increment("0" + s)
	}

	var lastChar = s[len(s)-1]
	fragment := s[0 : len(s)-1]
	if strings.Index(chars, string(lastChar)) < 61 {
		lastChar = chars[strings.Index(chars, string(lastChar))+1]
		return fragment + string(lastChar)
	}

	return Base62Increment(fragment) + "0"
}
