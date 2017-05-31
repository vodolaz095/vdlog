package vdlog

//https://stackoverflow.com/a/36623663/1885921
//https://stackoverflow.com/questions/33126743/how-do-i-join-my-bot-to-the-channel/36623663#36623663

import (
	"fmt"
	"net/http"
	"net/url"
)

const telegramUrlTemplate = "https://api.telegram.org/bot%s/sendMessage"

func createTelegramSink(botToken, chatId string, level EventLevel) func(e Event) error {
	return func(e Event) error {
		if e.Level > level {
			return nil
		}
		telegramUrl := fmt.Sprintf(telegramUrlTemplate, botToken)
		body := url.Values{}
		body.Set("text", string(e.ToIndentedJSON())) //TODO better formating
		body.Set("chat_id", chatId)
		resp, err := http.PostForm(telegramUrl, body)
		defer resp.Body.Close()
		if err != nil {
			return err
		}
		return nil
	}
}

//LogToTelegram allows to send events to telegram channel/chat using telegram bot api
func LogToTelegram(botToken, chatId string, level EventLevel) {
	AddSink("telegram", createTelegramSink(botToken, chatId, level))
}
