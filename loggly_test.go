package vdlog

import (
	"os"
	"testing"
	"time"
)

func TestLogglySink(t *testing.T) {
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
	var token string

	token = os.Getenv("LOGGLY_TOKEN")
	if token == "" {
		t.Skip("set LOGGLY_TOKEN environment to run this test")
	}

	logglySink := createLogglySync(token, true, LevelSilly)
	err := logglySink(evnt)
	if err != nil {
		t.Error(err)
	}
}
