package vdlog

import (
	"fmt"
	"os"

	"github.com/fatih/color"
)

var consoleLogLevelTrigger EventLevel

//SetConsoleVerbosity allows to set verbosity level for console log sink
func SetConsoleVerbosity(level EventLevel) {
	consoleLogLevelTrigger = level
}

func consoleSink(e Event) error {
	if e.Level > consoleLogLevelTrigger {
		return nil
	}

	switch e.Level {
	case LevelSilly:
		color.Set(color.FgMagenta)
		break
	case LevelDebug:
		color.Set(color.FgBlue)
		break
	case LevelVerbose:
		color.Set(color.FgCyan)
		break
	case LevelInfo:
		color.Set(color.FgGreen)
		break
	case LevelWarn:
		color.Set(color.FgYellow)
		break
	case LevelError:
		return nil
	}
	_, err := fmt.Fprintf(os.Stderr, "%s\n", e.String())
	color.Unset()
	return err
}

func consoleErrorSink(e Event) error {
	if e.Level > LevelError {
		return nil
	}
	color.Set(color.FgRed)
	_, err := fmt.Fprintf(os.Stderr, "%s\n", e.String())
	color.Unset()
	return err
}
