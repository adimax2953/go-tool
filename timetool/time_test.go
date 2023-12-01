package timetool_test

import (
	"testing"
	"time"

	"github.com/adimax2953/go-tool/timetool"
	LogTool "github.com/adimax2953/log-tool"
)

func TestQQ(t *testing.T) {

	LogTool.LogInfo("當日", timetool.GetDurationUntilMidnight())
	LogTool.LogInfo("到次月1日", timetool.GetDurationUntilNextMonth())
	currentTime := time.Now()
	targetTime := currentTime.AddDate(1, 1, 20)

	LogTool.LogInfo("自訂義時間", timetool.GetDurationUntil(targetTime))
}
