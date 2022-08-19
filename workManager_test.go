package gotool_test

import (
	"testing"

	gotool "github.com/adimax2953/go-tool"
	LogTool "github.com/adimax2953/log-tool"
)

func Test_WorkPool(t *testing.T) {
	runTimes := 1000000

	wp := gotool.NewWorkPool(runTimes)
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
	LogTool.LogInfo("Hello World!", i)
}
