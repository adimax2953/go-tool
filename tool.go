package gotool

import (
	"math/rand"

	LogTool "github.com/adimax2953/log-tool"
)

// 取亂數 1~num
func RanInt(num int) int {
	if num < 0 {
		LogTool.LogDebug("傳入了負數 %d", num)
		return 0
	}

	if num == 0 {
		num++
	}
	rndInt := rand.Intn(num) + 1
	return rndInt
}
