package vdlog

import (
	"os"
	"testing"
	"time"
)

func TestLogglySink(t *testing.T) {
	evnt := Event{
		Level:     LevelInfo,
		Payload:   "Hello from vdlog",
		Facility:  "vdlogUnitTest",
		Timestamp: time.Now(),
		Line:      2,
		Filename:  "/var/www/localhost/index.php",
	}
	evnt.prepare()
	var token string

	token = os.Getenv("LOGGLY_TOKEN")
	if token == "" {
		t.Skip("set LOGGLY_TOKEN environment to run this test")
	}

	logglySink := createLogglySync(token, true)
	err := logglySink(evnt)
	if err != nil {
		t.Error(err)
	}
}
