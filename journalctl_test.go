// +build linux

package vdlog

import (
	"os"
	"testing"
	"time"
)

func TestJournaldSink(t *testing.T) {
	_, err:= os.Stat("/bin/logger")

	if os.IsNotExist(err) {
		t.Skip("unable to find /bin/logger binary - cannot perform the test")
	}


	evnt := Event{
		Level:     LevelInfo,
		Payload:   "Hello from vdlog",
		Facility:  "vdlogUnitTest",
		Timestamp: time.Now(),
		Line:      2,
		Filename:  "/var/www/localhost/index.php",
	}

	localJournaldSink := createJournaldSink("localhost", 514, true, true)
	err = localJournaldSink(evnt)
	if err != nil {
		t.Error(err)
	}
}
