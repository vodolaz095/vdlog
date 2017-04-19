package vdlog

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"
	"time"
)

func TestLogToFile(t *testing.T) {
	filename := fmt.Sprintf("%v%s%v", os.TempDir(), string(os.PathSeparator), "vdlog.log")
	os.Remove(filename)

	LogToFile(filename, LevelError, LevelSilly)
	EmitSilly("testFacility", H{"testing": "test"})
	EmitVerbose("testFacility", H{"testing": "test"})
	EmitDebug("testFacility", H{"testing": "test"})
	EmitInfo("testFacility", H{"testing": "test"})
	EmitWarn("testFacility", H{"testing": "test"})
	EmitError("testFacility", fmt.Errorf("some test %s", "error"), H{"testing": "test"})

	time.Sleep(100 * time.Millisecond)
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		t.Errorf("Error reading file with log sink - %s - %s", filename, err)
	}
	if len(data) == 0 {
		t.Error("Empty data was saved")
	}
	os.Remove(filename)
}
