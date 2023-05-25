package tgbottool_test

import (
	"fmt"
	"testing"

	tg "github.com/adimax2953/go-tool/tgbottool"
)

func Test_SendtoTG(t *testing.T) {
	return
	TgbotChatID := -603254809
	TgbotToken := "5222610499:AAEdDXEcLnNvQnHx07w38wU3sysONI0-MzU"

	msg := fmt.Sprintf("\n事件：" + "山豬開工了阿")

	tg.SendToTG(TgbotChatID, TgbotToken, "Prod", msg)

	TgbotChatID = -544039489
	tg.SendToTG(TgbotChatID, TgbotToken, "Prod", msg)

	TgbotChatID = -662611117
	tg.SendToTG(TgbotChatID, TgbotToken, "Prod", msg)

	TgbotChatID = -800990157
	tg.SendToTG(TgbotChatID, TgbotToken, "Prod", msg)
}
