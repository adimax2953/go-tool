package gotool_test

import (
	"testing"

	gotool "github.com/adimax2953/go-tool"
	LogTool "github.com/adimax2953/log-tool"
)

func Test_WorkPool(t *testing.T) {

	for i := 0; i < 1; i++ {
		test2(i)
	}

	return

	runTimes := 1000000
	wp := gotool.NewWorkPool(runTimes)
	p, _ := wp.NewWorkPoolWithFunc(runTimes, func(i interface{}) {
		test2(i)
	})
	for i := 0; i < runTimes; i++ {
		_ = p.Invoke(int(i))
	}
	defer p.Release()
	defer wp.Release()
}
func test() {
	LogTool.LogInfo("Hello World!")
}

var str string = "z"

func test2(i interface{}) {
	str = gotool.Base62Increment(str)
	LogTool.LogDebug(str)
}
