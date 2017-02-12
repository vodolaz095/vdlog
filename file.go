package vdlog

import (
	"fmt"
	"os"
)

func generateFileSink(filename string, minLevel, maxLevel EventLevel) func(e Event) error {
	return func(e Event) error {
		if e.Level < minLevel {
			return nil
		}
		if e.Level > maxLevel {
			return nil
		}

		file, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
		if err != nil {
			return err
		}
		defer file.Close()
		_, err = fmt.Fprintln(file, e.String())
		return err
	}
}

//LogErrorsToFile outputs only errors to file provided
func LogErrorsToFile(filename string) {
	AddSink(filename, generateFileSink(filename, LevelError, LevelError))
}

//LogNormalToFile outputs only WARN---DEBUG messages to file provided
func LogNormalToFile(filename string) {
	AddSink(filename, generateFileSink(filename, LevelWarn, LevelDebug))
}

//LogToFile allows to output desired range of events by level intp file provided
func LogToFile(filename string, minLevel, maxLevel EventLevel) {
	AddSink(filename, generateFileSink(filename, minLevel, maxLevel))
}
