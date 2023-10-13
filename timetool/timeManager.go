package timetool

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

//	轉化成"2006-01-02 15:04:05"的時間模版
//
// 輸入為time.Now().Unix()
func TimeStamptoDateTime(timestamp int64) string {
	//timestamp := time.Now().Unix()
	datetime := time.Unix(timestamp, 0).Format(DateTimeLayout)
	return datetime
}

// 現在時間 格式為"2006-01-02"
func DateFromNow() string {
	return time.Now().Format(DateLayout)
}

// 現在時間 格式為"15:04:05"
func TimeFromNow() string {
	return time.Now().Format(TimeLayout)
}

// 現在時間 格式為"2006-01-02 15:04:05"
func DateTimeFromNow() string {
	return time.Now().Format(DateTimeLayout)
}

// 透過time.Now().Unix()出來的秒數轉為 "2006-01-02"
func DateFromTimeStamp(timestamp int64) string {
	return time.Unix(timestamp, 0).Format(DateLayout)
}

// 透過time.Now().Unix()出來的秒數轉為 "15:04:05"
func TimeFromTimeStamp(timestamp int64) string {
	return time.Unix(timestamp, 0).Format(TimeLayout)
}

// 透過time.Now().Unix()出來的秒數轉為 "2006-01-02" "15:04:05"
func DateTimeFromTimeStamp(timestamp int64) (string, string) {
	return time.Unix(timestamp, 0).Format(DateLayout), time.Unix(timestamp, 0).Format(TimeLayout)
}

// 透過time.Now()出來的年份跟週數
func GetWeek() (y, w int) {
	datetime := time.Now().Format(DateLayout)
	loc, _ := time.LoadLocation("Local")
	tmp, _ := time.ParseInLocation(DateLayout, datetime, loc)
	return tmp.ISOWeek()
}

func GetBetweenDates(start string, end string) []string {
	start = start[:10]
	end = end[:10]
	startTime, _ := time.Parse("2006-01-02", start)
	endTime, _ := time.Parse("2006-01-02", end)
	dates := make([]string, 0)
	dates = append(dates, startTime.Format("2006-01-02"))
	i := 0
	for {
		i++
		queryTime := startTime.AddDate(0, 0, i)
		if queryTime.After(endTime) {
			break
		}
		dates = append(dates, queryTime.Format("2006-01-02"))
	}
	return dates
}

func GetBetweenTimes(start string, end string) []string {
	startTime, _ := time.Parse("2006-01-02 15:04:05", start)
	endTime, _ := time.Parse("2006-01-02 15:04:05", end)
	dates := make([]string, 0)
	dates = append(dates, startTime.Format("2006-01-02 15:04:05"))
	i := 0
	var queryTime time.Time = startTime
	for {
		i++
		queryTime = queryTime.Add(30 * time.Minute)
		if queryTime.After(endTime) || queryTime.Equal(endTime) {
			break
		}
		dates = append(dates, queryTime.Format("2006-01-02 15:04:05"))
	}
	return dates
}

// GetDurationUntilMidnight -取得到當日0點的時間差(秒)
func GetDurationUntilMidnight() int {
	// 取得目前時間
	currentTime := time.Now()

	// 取得今天的午夜時間
	midnight := time.Date(
		currentTime.Year(),  // 使用目前年份
		currentTime.Month(), // 使用目前月份
		currentTime.Day(),   // 使用當前日期
		0, 0, 0, 0,          // 凌晨的時間
		currentTime.Location(),
	)

	// 計算時間差
	durationUntilMidnight := midnight.Sub(currentTime)

	// 將時間差轉換為秒
	secondsUntilMidnight := durationUntilMidnight.Seconds()
	return int(secondsUntilMidnight)
}

// GetDurationUntilNextMonth -取得到次月1日0點的時間差(秒)
func GetDurationUntilNextMonth() int {
	// 取得目前時間
	currentTime := time.Now()

	// 取得下個月1號的凌晨時間
	nextMonth := currentTime.AddDate(0, 1, 0)
	nextMonthMidnight := time.Date(
		nextMonth.Year(),  // 下個月的年份
		nextMonth.Month(), // 下個月的月份
		1,                 // 下個月1號
		0, 0, 0, 0,        // 凌晨的時間
		currentTime.Location(),
	)

	// 計算時間差
	durationUntilNextMonth := nextMonthMidnight.Sub(currentTime)

	// 將時間差轉換為秒
	secondsUntilNextMonth := durationUntilNextMonth.Seconds()
	return int(secondsUntilNextMonth)

}

// GetDurationUntil -取得到指定時間的時間差(秒)
func GetDurationUntil(targetTime time.Time) int {
	// 取得目前時間
	currentTime := time.Now()

	// 計算時間差
	durationUntilTargetTime := targetTime.Sub(currentTime)

	// 將時間差轉換為秒
	secondsUntilTargetTime := durationUntilTargetTime.Seconds()

	return int(secondsUntilTargetTime)
}
