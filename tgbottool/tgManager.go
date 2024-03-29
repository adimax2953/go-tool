package tgbottool

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"

	timetool "github.com/adimax2953/go-tool/timetool"

	LogTool "github.com/adimax2953/log-tool"
)

// SendTextToTelegramChat -發送訊息到TelegramChat
func SendTextToTelegramChat(chatId int, text string, Token string) (string, error) {

	var telegramApi string = "https://api.telegram.org/bot" + Token + "/sendMessage"

	response, err := http.PostForm(
		telegramApi,
		url.Values{
			"chat_id": {strconv.Itoa(chatId)},
			"text":    {text},
		})

	if err != nil {
		LogTool.LogError(fmt.Sprintf("error when posting text to the chat: %s", err.Error()))
		return "", err
	}
	defer response.Body.Close()

	var bodyBytes, errRead = ioutil.ReadAll(response.Body)
	if errRead != nil {
		LogTool.LogError(fmt.Sprintf("error in parsing telegram answer %s", errRead.Error()))
		return "", err
	}
	bodyString := string(bodyBytes)
	LogTool.LogDebug(fmt.Sprintf("Body of Telegram Response: %s", bodyString))

	return bodyString, nil
}

func SendGameEventToTG(TgbotChatID int, TgbotToken, game, env, reson string) {
	msg := fmt.Sprintf("\n" + game + "遊戲服務器\n環境：" + env + "\n發生時間：" + timetool.TimeNowStr() + "\n版本號：" + "")
	msg += reson
	LogTool.LogSystem(msg)
	SendTextToTelegramChat(TgbotChatID, msg, TgbotToken)
}

func SendToTG(TgbotChatID int, TgbotToken, env, reson string) {
	msg := fmt.Sprintf("\nXXX遊戲服務器\n環境：" + env + "\n發生時間：" + timetool.TimeNowStr() + "\n版本號：" + "")
	msg += reson
	LogTool.LogSystem(msg)
	SendTextToTelegramChat(TgbotChatID, msg, TgbotToken)
}
