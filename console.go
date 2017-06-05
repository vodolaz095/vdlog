package vdlog

import (
	"fmt"
	"os"

	"github.com/fatih/color"
)

var consoleLogLevelTrigger EventLevel

var useJSONinConsole = false

//SetConsoleVerbosity allows to set verbosity level for console log sink
func SetConsoleVerbosity(level EventLevel) {
	consoleLogLevelTrigger = level
}

//SetConsoleJSON makes console sink output properly formatted json
func SetConsoleJSON() {
	useJSONinConsole = true
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
	var err error
	if useJSONinConsole {
		_, err = fmt.Fprintln(os.Stderr, string(e.ToIndentedJSON()))
	} else {
		_, err = fmt.Fprintln(os.Stderr, e.String())
	}

	color.Unset()
	return err
}

func consoleErrorSink(e Event) error {
	if e.Level > LevelError {
		return nil
	}
	var err error
	color.Set(color.FgRed)
	if useJSONinConsole {
		_, err = fmt.Fprintln(os.Stderr, string(e.ToIndentedJSON()))
	} else {
		_, err = fmt.Fprintln(os.Stderr, e.String())
	}
	color.Unset()
	return err
}
