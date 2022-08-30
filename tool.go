package gotool

import (
	"math"
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

var chars string = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

// 遞增字符串ByBase62
func Base62Increment(s string) string {
	defer RecoverPanic()

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

// Encode10To62 - 10進制轉 62
func Encode10To62(num int64) string {
	bytes := []byte{}
	for num > 0 {
		bytes = append(bytes, chars[num%62])
		num = num / 62
	}
	reverse(bytes)
	return string(bytes)
}

// Decode62To10 - 62進制轉 10
func Decode62To10(str string) int64 {
	var num int64
	n := len(str)
	for i := 0; i < n; i++ {
		pos := strings.IndexByte(chars, str[i])
		num += int64(math.Pow(62, float64(n-i-1)) * float64(pos))
	}
	return num
}

func reverse(a []byte) {
	for left, right := 0, len(a)-1; left < right; left, right = left+1, right-1 {
		a[left], a[right] = a[right], a[left]
	}
}
