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
	}

	err := journaldSink(evnt)
	if err != nil {
		t.Error(err)
	}
}
