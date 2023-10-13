package randtool_test

import (
	"fmt"
	"testing"

	gotool "github.com/adimax2953/go-tool"
	"github.com/adimax2953/go-tool/randtool"
	LogTool "github.com/adimax2953/log-tool"
	"github.com/shopspring/decimal"
)

func Test_Math(t *testing.T) {
	LogTool.LogDebug("", gotool.Int32ToStr(randtool.GetRandom(10))+
		gotool.Int32ToStr(randtool.GetRandom(10))+
		gotool.Int32ToStr(randtool.GetRandom(10))+
		gotool.Int32ToStr(randtool.GetRandom(10))+
		gotool.Int32ToStr(randtool.GetRandom(10))+
		gotool.Int32ToStr(randtool.GetRandom(10)))
	return
	var runtimes int64 = 1000000
	var win [randtool.NMaxHit]int64 = [randtool.NMaxHit]int64{0, 0, 0, 0, 0}
	var wintimes int64 = 0
	for i := 0; i < int(runtimes); i++ {

		idx := randtool.Lottery([]int64{0, 0, 0, 0})
		win[idx]++
		wintimes++
		if wintimes == runtimes {
			LogTool.LogDebug(fmt.Sprintf("開1 %d 次", win[1]), decimal.NewFromInt(win[1]).Div(decimal.NewFromInt(runtimes)).String())
			LogTool.LogDebug(fmt.Sprintf("開2 %d 次", win[2]), decimal.NewFromInt(win[2]).Div(decimal.NewFromInt(runtimes)).String())
			LogTool.LogDebug(fmt.Sprintf("開3 %d 次", win[3]), decimal.NewFromInt(win[3]).Div(decimal.NewFromInt(runtimes)).String())
			LogTool.LogDebug(fmt.Sprintf("開4 %d 次", win[4]), decimal.NewFromInt(win[4]).Div(decimal.NewFromInt(runtimes)).String())

			a := decimal.NewFromInt(win[1]).Div(decimal.NewFromInt(runtimes))
			b := decimal.NewFromInt(win[2]).Div(decimal.NewFromInt(runtimes))
			c := decimal.NewFromInt(win[3]).Div(decimal.NewFromInt(runtimes))
			d := decimal.NewFromInt(win[4]).Div(decimal.NewFromInt(runtimes))
			LogTool.LogDebug("", fmt.Sprintf("總和 %s ", a.Add(b).Add(c).Add(d).String()))
			return
		}
	}
}
