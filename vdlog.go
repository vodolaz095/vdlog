package vdlog

import (
	"fmt"
	"io"
	"os"
	"regexp"
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

//AddSink allows to add custom events' sink by defined event processing function
func AddSink(name string, sink func(e Event) error) {
	sinks[name] = sink
}

//BrokenSinkReporter is a function being called when any of sinks is broken
var BrokenSinkReporter func(brokenSinkName string, eventThatCloggedIt Event, errorRecievedFromSink error)

//IoWriterSink is a struct that implements io.Writer for usage for https://godoc.org/log#SetOutput with Level and Facility defined
type IoWriterSink struct {
	Level    EventLevel
	Facility string
}

//Write just sends any slice of bytes as payload of new event with prepending timestamp removed
func (i IoWriterSink) Write(p []byte) (n int, err error) {
	n = len(p)
	r, err := regexp.Compile(`\d{4}(\/\d\d){2}\s\d{1,2}\:\d{1,2}\:\d{1,2}\s`)
	p = r.ReplaceAll(p, []byte(""))
	vdlogEntryPoint(i.Level, i.Facility, "%s", string(p))
	return
}

//CreateIoWriter creates io.Writer struct with level and facility defined to be used with https://godoc.org/log#SetOutput
func CreateIoWriter(level EventLevel, facility string) io.Writer {
	return IoWriterSink{level, facility}
}
