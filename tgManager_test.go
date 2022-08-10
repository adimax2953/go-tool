package gotool_test

import (
	"fmt"
	"testing"

	gotool "github.com/adimax2953/go-tool"
	LogTool "github.com/adimax2953/log-tool"
	logType "github.com/adimax2953/log-tool/logType"
)

func Test_SendtoTG(t *testing.T) {

	TgbotChatID := -603254809
	TgbotToken := "5222610499:AAGNsiLffs9jxR0X1xfQpPi2MPV0HBRnxtw"

	msg := fmt.Sprintf("\n事件：" + "山豬開工了阿")
	//for true {
	LogTool.LogDebug("山豬開工了阿")
	LogTool.LogBase(logType.Config, "山豬開工了阿")

	gotool.SendToTG(TgbotChatID, TgbotToken, msg)
	//}
}
