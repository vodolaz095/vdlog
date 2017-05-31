package vdlog

//https://github.com/segmentio/go-loggly/blob/master/loggly.go

import (
	"bytes"
	"fmt"
	"net/http"
)

//const logglyURL = "logs-01.loggly.com/inputs/TOKEN/tag/http/"

func createLogglySync(token string, secure bool, logglyLogLevelTrigger EventLevel) func(e Event) error {
	return func(e Event) error {
		if e.Level > logglyLogLevelTrigger {
			return nil
		}

		var logglyURL string
		if secure {
			logglyURL = fmt.Sprintf("https://logs-01.loggly.com/inputs/%s/tag/http", token)
		} else {
			logglyURL = fmt.Sprintf("http://logs-01.loggly.com/inputs/%s/tag/http", token)
		}

		body := bytes.NewBuffer(e.ToJSON())
		resp, err := http.Post(logglyURL, "application/json; charset=utf-8", body)
		defer resp.Body.Close()
		if err != nil {
			return err
		}
		return nil
	}
}

//LogToLoggly allows to send messages to Loggly.com, if secure is true, https is used, which can increase security but reduce bandwidth
func LogToLoggly(token string, secure bool, level EventLevel) {
	AddSink("loggly", createLogglySync(token, secure, level))
}
