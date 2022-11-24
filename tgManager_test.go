package gotool_test

import (
	"fmt"
	"testing"

	gotool "github.com/adimax2953/go-tool"
)

func Test_SendtoTG(t *testing.T) {

	TgbotChatID := -603254809
	TgbotToken := "5222610499:AAEdDXEcLnNvQnHx07w38wU3sysONI0-MzU"

	msg := fmt.Sprintf("\n事件：" + "山豬開工了阿")

	gotool.SendToTG(TgbotChatID, TgbotToken, "Prod", msg)

	TgbotChatID = -544039489
	gotool.SendToTG(TgbotChatID, TgbotToken, "Prod", msg)

	TgbotChatID = -662611117
	gotool.SendToTG(TgbotChatID, TgbotToken, "Prod", msg)

	TgbotChatID = -800990157
	gotool.SendToTG(TgbotChatID, TgbotToken, "Prod", msg)
}
