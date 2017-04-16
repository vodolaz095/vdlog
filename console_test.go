package vdlog

import (
	"testing"
	"time"
)

func TestLog(t *testing.T) {
	SetConsoleVerbosity(LevelSilly)

	Silly("testFacility", "testing", "Silly", 1, 2, 3)
	Sillyf("testFacility", "testing %s", "Sillyf")

	Verbose("testFacility", "testing", "Verbose", 1, 2, 3)
	Verbosef("testFacility", "testing %s", "Verbosef")

	Debug("testFacility", "testing", "Debug", 1, 2, 3)
	Debugf("testFacility", "testing %s", "Debugf")

	Info("testFacility", "testing", "info", 1, 2, 3)
	Infof("testFacility", "testing %s", "Infof")

	Warn("testFacility", "testing", "warn", 1, 2, 3)
	Warnf("testFacility", "testing %s", "Warn")

	Error("testFacility", "testing", "Error", 1, 2, 3)

	var evnt Event
	evnt.Level = LevelInfo
	evnt.Facility = "testFacility"
	evnt.Payload = "testing"
	evnt.Emit()

	time.Sleep(100 * time.Millisecond)
}
