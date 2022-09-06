package gotool_test

import (
	"fmt"
	"testing"
	"time"

	gotool "github.com/adimax2953/go-tool"
	LogTool "github.com/adimax2953/log-tool"
)

func Test_SendtoTG(t *testing.T) {

	TgbotChatID := -603254809
	TgbotToken := "5222610499:AAGNsiLffs9jxR0X1xfQpPi2MPV0HBRnxtw"

	msg := fmt.Sprintf("\n事件：" + "山豬開工了阿")
	//for true {
	LogTool.LogDebug("山豬開工了阿")

	gotool.SendToTG(TgbotChatID, TgbotToken, "Prod", msg)

	LogTool.LogDebug("", time.Now().UnixMilli())
	//}
}
