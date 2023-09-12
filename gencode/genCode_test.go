package gencode_test

import (
	"testing"

	"github.com/adimax2953/go-tool/gencode"
	LogTool "github.com/adimax2953/log-tool"
)

func TestQQ(t *testing.T) {
	LogTool.LogInfo(gencode.GenCode(300, 200, 4))
}
