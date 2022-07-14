package gotool_test

import (
	gotool "github.com/adimax2953/go-tool"
	"testing"
	"fmt"
)

func Test_SendtoTG(t *testing.T)  {
	
		TgbotChatID := -603254809
		TgbotToken := "5222610499:AAGNsiLffs9jxR0X1xfQpPi2MPV0HBRnxtw"
	
	msg := fmt.Sprintf("\n事件：" + "山豬想睡覺了阿")
	gotool.SendToTG(TgbotChatID,TgbotToken,msg)
}
