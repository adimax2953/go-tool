package gotool

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"github.com/adimax2953/go-tool/jsontool"
	LogTool "github.com/adimax2953/log-tool"
)

type NonNegative_Integer interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

type Negative_Number interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64
}

type NonNegative_Number interface {
	~float32 | ~float64 | ~complex64 | ~complex128
}

// Int轉str - 將整形轉換成字符串
func IntToStr(n int) string {
	return strconv.FormatInt(int64(n), 10)
}

// Int32轉str - 將整形轉換成字符串
func Int32ToStr(n int32) string {
	return strconv.FormatInt(int64(n), 10)
}

// Int64轉str - 將整形轉換成字符串
func Int64ToStr(n int64) string {
	return strconv.FormatInt(n, 10)
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

// StrToInt - 字串轉 Int
func StrToInt(str string) int {
	num, err := strconv.Atoi(str)
	if err != nil {
		LogTool.LogDebug("轉int出錯")
	}
	return num
}

// StrToInt32 - 字串轉 Int32
func StrToInt32(str string) (int32, error) {
	num, err := strconv.ParseInt(str, 10, 32) //轉完可能變int64
	if err != nil {
		LogTool.LogDebug("轉int32出錯")
	}
	return int32(num), err
}

// StrToInt64 - 字串轉 Int64
func StrToInt64(str string) int64 {
	num, err := strconv.ParseInt(str, 10, 64) //轉完可能變int64
	if err != nil {
		LogTool.LogDebug("轉int64出錯%v", err)
	}
	return int64(num)
}

// GetStringEnd - 取得字串最後一碼字
func GetStringEnd(str string) string {
	strlen := len(str)
	if strlen == 0 {
		return "14"
	}
	return str[strlen-1 : strlen]
}

// InterfaceToString - Interface轉字串
func InterfaceToString(val interface{}) (res string) {
	if val == nil {
		return ""
	}

	switch v := val.(type) {
	case float64:
		res = strconv.FormatFloat(val.(float64), 'f', 6, 64)
	case float32:
		res = strconv.FormatFloat(float64(val.(float32)), 'f', 6, 32)
	case int:
		res = strconv.FormatInt(int64(val.(int)), 10)
	case int32:
		res = strconv.FormatInt(int64(val.(int32)), 10)
	case int64:
		res = strconv.FormatInt(val.(int64), 10)
	case uint:
		res = strconv.FormatUint(uint64(val.(uint)), 10)
	case uint64:
		res = strconv.FormatUint(val.(uint64), 10)
	case uint32:
		res = strconv.FormatUint(uint64(val.(uint32)), 10)
	case json.Number:
		res = val.(json.Number).String()
	case string:
		res = val.(string)
	case []byte:
		res = string(v)
	default:
		res = ""
	}
	return
}

// DataConvert - 資料轉換
func DataConvert(from interface{}, dst interface{}) error {
	if reflect.ValueOf(from).Kind() != reflect.Ptr {
		return errors.New("from need ptr.")
	}
	if reflect.ValueOf(dst).Kind() != reflect.Ptr {
		return errors.New("dst need ptr.")
	}

	tmpStr, err := jsontool.JsonMarshal(from)

	if err := jsontool.JsonUnmarshal(tmpStr, dst); err != nil {
		LogTool.LogError("DataConvert JsonUnmarshal Error", err)
		return err
	}

	_, err2 := jsontool.JsonMarshal(dst)

	if err2 != nil {
		return err
	}

	return nil
}

// ParseStrToArrayInt32 - ex: 123,555 -> [123,555]
func ParseStrToArrayInt32(target string, sep string) []int32 {
	target = strings.TrimSpace(target)
	if target == "" {
		return nil
	}

	ids := strings.Split(target, sep)
	result := make([]int32, len(ids))
	for index, id := range ids {
		id, err := strconv.ParseInt(strings.TrimSpace(id), 10, 32)

		if err != nil {
			return nil
		}

		result[index] = int32(id)
	}
	return result
}

// ParseStrToArrayStr - ex: "嘎抓","打妹" -> ["嘎抓","打妹"]
func ParseStrToArrayStr(target string, sep string) []string {
	target = strings.TrimSpace(target)
	if target == "" {
		return nil
	}
	return strings.Split(target, sep)
}
