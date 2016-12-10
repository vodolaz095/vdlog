package vdlog

import (
	"testing"
	"time"
)

func TestLog(t *testing.T) {
	SetConsoleVerbosity(LevelSilly)

	Silly("testFacility", "testing %s", "test")
	Verbose("testFacility", "testing %s", "test")
	Debug("testFacility", "testing %s", "test")
	Info("testFacility", "testing %s", "test")
	Warn("testFacility", "testing %s", "test")
	Error("testFacility", "testing %s", "test")
	Error("testFacility", "Simple string")

	time.Sleep(100 * time.Millisecond)
}
