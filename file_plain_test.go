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
	Sillyf("testFacility", "testing %s", "test")
	Verbosef("testFacility", "testing %s", "test")
	Debugf("testFacility", "testing %s", "test")
	Infof("testFacility", "testing %s", "test")
	Warnf("testFacility", "testing %s", "test")

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
