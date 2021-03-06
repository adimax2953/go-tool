package gotool

import (
	"fmt"
	"time"
)

const (
	DateLayout     = "2006-01-02"
	TimeLayout     = "15:04:05"
	DateTimeLayout = "2006-01-02 15:04:05"
)

// 日期格式轉時間戳
func TimestrToTimestamp(time_str string, flag int) int64 {
	var t int64
	loc, _ := time.LoadLocation("Local")
	if flag == 1 {
		t1, _ := time.ParseInLocation("2006.01.02 15:04:05", time_str, loc)
		t = t1.Unix()
	} else if flag == 2 {
		t1, _ := time.ParseInLocation("2006-01-02 15:04", time_str, loc)
		t = t1.Unix()
	} else if flag == 3 {
		t1, _ := time.ParseInLocation("2006-01-02", time_str, loc)
		t = t1.Unix()
	} else if flag == 4 {
		t1, _ := time.ParseInLocation("2006.01.02", time_str, loc)
		t = t1.Unix()
	} else {
		t1, _ := time.ParseInLocation("2006-01-02 15:04:05", time_str, loc)
		t = t1.Unix()
	}
	return t
}

// TimeNowStr 輸出格式為 2019/11/4 20:15:26
func TimeNowStr() string {
	n := time.Now()
	timeStr := fmt.Sprintf("%d/%d/%d %d:%d:%d", n.Year(), n.Month(), n.Day(), n.Hour(), n.Minute(), n.Second())
	return timeStr
}

//  轉化成"2006-01-02 15:04:05"的時間模版
// 輸入為time.Now().Unix()
func TimeFormatDate(timestamp int64) string {
	//timestamp := time.Now().Unix()
	datetime := time.Unix(timestamp, 0).Format(DateTimeLayout)
	return datetime
}

// 現在時間 格式為""2006-01-02""
func DateFromNow() string {
	return time.Now().Format(DateLayout)
}

// 透過time.Now().Unix()出來的秒數轉為 "2006-01-02"
func DateFromTimeStamp(timestamp int64) string {
	return time.Unix(timestamp, 0).Format(DateLayout)
}
