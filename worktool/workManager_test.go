package worktool_test

import (
	"testing"

	gotool "github.com/adimax2953/go-tool"
	timetool "github.com/adimax2953/go-tool/timetool"
	"github.com/adimax2953/go-tool/worktool"
	LogTool "github.com/adimax2953/log-tool"
)

func Test_WorkPool(t *testing.T) {

	return
	runTimes := 1000000

	wp := worktool.NewWorkPool(runTimes)
	// for i := 0; i < runTimes; i++ {
	// 	wp.SubmitTask(test)
	// }

	p, _ := wp.NewWorkPoolWithFunc(runTimes, func(i interface{}) {
		test2(i)
	})
	for i := 0; i < runTimes; i++ {
		_ = p.Invoke(int32(i))
	}
	defer p.Release()
	defer wp.Release()
}
func test() {
	LogTool.LogInfo("Hello World!")
}
func test2(i interface{}) {
	y, w := timetool.GetWeek()
	LogTool.LogInfo("Hello", gotool.Decode62To10(gotool.Encode10To62(int64(y))), gotool.Decode62To10(gotool.Encode10To62(int64(w))))
	LogTool.LogInfo("Hello", gotool.Encode10To62(int64(y)), gotool.Encode10To62(int64(w)))
}
