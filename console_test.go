package vdlog

import (
	"fmt"
	"testing"
	"time"
)

func TestLog(t *testing.T) {
	SetConsoleVerbosity(LevelSilly)

	EmitSilly("testFacility", H{"testing": "test"})
	EmitVerbose("testFacility", H{"testing": "test"})
	EmitDebug("testFacility", H{"testing": "test"})
	EmitInfo("testFacility", H{"testing": "test"})
	EmitWarn("testFacility", H{"testing": "test"})
	EmitError("testFacility", fmt.Errorf("some test %s", "error"), H{"testing": "test"})

	time.Sleep(100 * time.Millisecond)
}
