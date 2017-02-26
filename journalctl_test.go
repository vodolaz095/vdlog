// +build linux

package vdlog

import (
	"testing"
	"time"
)

func TestJournaldSink(t *testing.T) {
	evnt := Event{
		Level:     LevelInfo,
		Payload:   "Hello from vdlog",
		Facility:  "vdlogUnitTest",
		Timestamp: time.Now(),
		Line:      2,
		Filename:  "/var/www/localhost/index.php",
	}

	localJournaldSink := createJournaldSink("localhost", 514, true, true)
	err := localJournaldSink(evnt)
	if err != nil {
		t.Error(err)
	}
}
