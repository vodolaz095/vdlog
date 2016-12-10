package vdlog

import (
	"fmt"
	"os"
	"runtime"
	"time"
)

var spine chan Event

var sinks map[string](func(e Event) error)

func init() {
	sinks = make(map[string](func(e Event) error), 0)
	SetConsoleVerbosity(LevelInfo)
	AddSink("STDOUT", consoleSink)
	AddSink("STDERR", consoleErrorSink)
	spine = make(chan Event, 100)
	go func() {
		for {
			event := <-spine
			for name, sink := range sinks {
				err := sink(event)
				if err != nil {
					if BrokenSinkReporter != nil {
						BrokenSinkReporter(name, event, err)
					} else {
						fmt.Fprintf(os.Stderr, "Sink %s failed with error '%s' while processing '%s'!\n", name, err.Error(), event.StringWithCaller())
					}
				}
			}
		}
	}()
}

//log is an internal function used for making event objects and sending them to spine channel
func log(level EventLevel, facility, format string, data ...interface{}) {
	_, file, line, _ := runtime.Caller(2)
	spine <- Event{
		Level:     level,
		Facility:  facility,
		Timestamp: time.Now(),
		Filename:  file,
		Line:      line,
		Payload:   fmt.Sprintf(format, data...),
	}
}

//AddSink allows to add custom events' sink by defined event processing function
func AddSink(name string, sink func(e Event) error) {
	sinks[name] = sink
}

//BrokenSinkReporter is a function being called when any of sinks is broken
var BrokenSinkReporter func(brokenSinkName string, eventThatCloggedIt Event, errorRecievedFromSink error)
