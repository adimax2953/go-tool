package gotool

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"math"
	"runtime/debug"
	"strings"

	LogTool "github.com/adimax2953/log-tool"
	"github.com/shopspring/decimal"
	"golang.org/x/text/encoding/charmap"
	"golang.org/x/text/encoding/simplifiedchinese"
)

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

func Compress(s string) string {
	//使用GBK字符集encode
	gbk, err := simplifiedchinese.GBK.NewEncoder().Bytes([]byte(s))
	if err != nil {
		//logrus.Error(err)
		return ""
	}

	//轉為ISO8859_1，也就是latin1字串集
	latin1, err := charmap.ISO8859_1.NewDecoder().Bytes(gbk)
	if err != nil {
		return ""
	}

	//使用gzip壓縮
	var buf bytes.Buffer
	zw := gzip.NewWriter(&buf)

	_, err = zw.Write(latin1)
	if err != nil {
	}

	if err := zw.Close(); err != nil {
	}

	//使用base64編碼
	return base64.StdEncoding.EncodeToString(buf.Bytes())
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

// Percent - 取得 百分比
func Percent(value1, value2 int64) string {
	if value1 == 0 || value2 == 0 {
		return "0"
	}

	result := decimal.NewFromInt(value2).Div(decimal.NewFromInt(value1)).Mul(decimal.NewFromInt(100)).StringFixed(2)

	return result
}
