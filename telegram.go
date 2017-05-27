package vdlog

//https://stackoverflow.com/a/36623663/1885921

//curl -v \
//-X POST \
//-d text="A message from your bot" \
//-d chat_id="-1001071402342" \
//https://api.telegram.org/bot286759464:AAFRalklssMW9hsZ592O8CxZo63QU7KM7d0/sendMessage

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
		body:= url.Values{}
		body.Set("text", e.String()) //TODO better formating
		body.Set("chat_id",chatId)
		resp, err := http.PostForm(telegramUrl, body)
		defer resp.Body.Close()
		if err != nil {
			return err
		}
		return nil
	}
}
