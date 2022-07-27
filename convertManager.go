package gotool

import (
	"fmt"
	"strconv"

	LogTool "github.com/adimax2953/log-tool"
)

// Int轉str - 將整形轉換成字符串
func IntToStr(n int) string {
	return strconv.FormatInt(int64(n), 10)
}

// Int轉str - 將整形轉換成字符串
func Int32ToStr(n int32) string {
	return strconv.FormatInt(int64(n), 10)
}

// Float轉Str - 浮點數轉換成字符串
func FloatToStr(f float64) string {
	return fmt.Sprintf("%f", f)
}

// RoundingTwo - 四捨五入取小數兩位
func RoundingTwo(value float64) float64 {
	value64, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", value), 64)
	return value64
}

// RoundingFour - 四捨五入取小數四位
func RoundingFour(value float64) float64 {
	value64, _ := strconv.ParseFloat(fmt.Sprintf("%.4f", value), 64)
	return value64
}

// RoundingSeven - 四捨五入取小數七位(後端統一)
func RoundingSeven(value float64) float64 {
	value64, _ := strconv.ParseFloat(fmt.Sprintf("%.7f", value), 64)
	return value64
}

// AbsInt32 -整數轉正
func AbsInt32(n int32) int32 {
	if n < 0 {
		return -n
	}
	return n
}

// Str2Int - 字串轉Int
func Str2Int(str string) int {
	num, err := strconv.Atoi(str)
	if err != nil {
		LogTool.LogDebug("轉int出錯")
	}
	return num
}

// Str2int32 - 字串轉 Int32
func Str2int32(str string) (int32, error) {
	num, err := strconv.ParseInt(str, 10, 32) //轉完可能變int64
	if err != nil {
		LogTool.LogDebug("轉int32出錯")
	}
	return int32(num), err
}

// Str2int64 - 字串轉Int64
func Str2int64(str string) int64 {
	num, err := strconv.ParseInt(str, 10, 64) //轉完可能變int64
	if err != nil {
		LogTool.LogDebug("轉int64出錯%v", err)
	}
	return int64(num)
}

// GetStringEnd -取得字串最後一碼字
func GetStringEnd(str string) string {
	strlen := len(str)
	if strlen == 0 {
		return "14"
	}
	return str[strlen-1 : strlen]
}
