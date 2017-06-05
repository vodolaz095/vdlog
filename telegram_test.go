package vdlog

import (
	"os"
	"testing"
	"time"
)

func TestTelegramSink(t *testing.T) {
	telegramToken := os.Getenv("TELEGRAM_BOT_TOKEN")
	if telegramToken == "" {
		t.Skip("Set environment variable TELEGRAM_BOT_TOKEN to run this test")
		return
	}
	telegramChatID := os.Getenv("TELEGRAM_CHAT_ID")
	if telegramChatID == "" {
		t.Skip("Set environment variable TELEGRAM_CHAT_ID to run this test")
		return
	}
	send := createTelegramSink(telegramToken, telegramChatID, LevelInfo)

	evnt := Event{
		Level: LevelInfo,
		Metadata: H{
			"Payload": "Hello from vdlog",
		},
		Type:      "vdlogUnitTest",
		Timestamp: time.Now(),
		Line:      2,
		Filename:  "/var/www/localhost/index.php",
	}

	err := send(evnt)
	if err != nil {
		t.Error(err)
	}
}
