package vdlog

import (
	"encoding/json"
	"fmt"
	"os"
	"runtime"
	"time"

	"github.com/satori/go.uuid"
)

//EventLevel is a type describing event's level
type EventLevel uint8

const (
	//LevelError is a level for errors that can wake you on 4 hours past midnight
	LevelError = EventLevel(iota)

	//LevelWarn is a level for an unexpected technical or business event happened, customers may be affected, but probably no immediate human intervention is required. On call people won't be called immediately, but support personnel will want to review these issues asap to understand what the impact is. Basically any issue that needs to be tracked but may not require immediate intervention.
	LevelWarn

	//LevelInfo are for things we want to see at high volume in case we need to forensically analyze an issue. System lifecycle events (system start, stop) go here. "Session" lifecycle events (login, logout, etc.) go here. Significant boundary events should be considered as well (e.g. database calls, remote API calls). Typical business exceptions can go here (e.g. login failed due to bad credentials). Any other event you think you'll need to see in production at high volume goes here.
	LevelInfo

	//LevelVerbose just about everything that doesn't make the "info" cut... any message that is helpful in tracking the flow through the system and isolating issues, especially during the development and QA phases. We use "debug" level logs for entry/exit of most non-trivial methods and marking interesting events and decision points inside methods.
	LevelVerbose

	//LevelDebug is for extremely detailed and potentially high volume logs that you don't typically want enabled even during normal development. Examples include dumping a full object hierarchy, logging some state during every iteration of a large loop, etc. String representation of event has information where was it called in code.
	LevelDebug

	//LevelSilly is putting every fart to log. String representation of event has information where was it called in code
	LevelSilly

	//EventDateFormat is a constant for formatting date output
	EventDateFormat = "Jan 02 15:04:05"

	//EventDateFormatMilli is a constant for formatting date output including milliseconds
	EventDateFormatMilli = "Jan 02 15:04:05.000"
)

//H is a type for storing metadata of event
type H map[string]interface{}

//Event represents anything to be logged
type Event struct {
	UUID        string     `json:"uuid" xml:"uuid"`
	Level       EventLevel `json:"level" xml:"level"`
	LevelString string     `json:"levelString" xml:"levelString"`
	Type        string     `json:"type" xml:"type"`
	Filename    string     `json:"filename" xml:"filename"`
	Line        int        `json:"line" xml:"line"`
	Called      string     `json:"called" xml:"called"`
	Hostname    string     `json:"hostname" xml:"hostname"`
	Pid         int        `json:"pid" xml:"pid"`
	Timestamp   time.Time  `json:"timestamp" xml:"timestamp"`
	Metadata    H          `json:"metadata" xml:"metadata"`
	Error       error      `json:"error" xml:"error"`
}

//vdlogEntryPoint is an internal function used for making event objects and sending them to spine channel
func vdlogEntryPoint(level EventLevel, eventType string, payload H, err error) {
	_, file, line, _ := runtime.Caller(2)
	evnt := Event{
		UUID:      uuid.NewV4().String(),
		Pid:       os.Getpid(),
		Level:     level,
		Timestamp: time.Now(),
		Filename:  file,
		Line:      line,
		Type:      eventType,
		Metadata:  payload,
	}
	hostname, _ := os.Hostname()
	evnt.Hostname = hostname
	evnt.LevelString = evnt.GetLevelString()
	evnt.Called = fmt.Sprintf("%s:%v", evnt.Filename, evnt.Line)
	if err != nil {
		evnt.Error = err
	} else {
		evnt.Error = nil
	}
	spine <- evnt
}

//Ago returns how long ago does the event was fired
func (e *Event) Ago() time.Duration {
	return time.Since(e.Timestamp)
}

//GetLevelString returns string representation of event level
func (e *Event) GetLevelString() (ret string) {
	switch e.Level {
	case LevelSilly:
		ret = "SILLY"
		break
	case LevelDebug:
		ret = "DEBUG"
		break
	case LevelVerbose:
		ret = "VERBOSE"
		break
	case LevelInfo:
		ret = "INFO"
		break
	case LevelWarn:
		ret = "WARN"
		break
	case LevelError:
		ret = "ERROR"
		break
	}
	return
}

//StringWithCaller returns string representation of an event with information where it was called in code and exactly when (to milliseconds)
func (e *Event) StringWithCaller() string {
	return fmt.Sprintf("%s %s %s <File: %s:%v>: %s", e.Timestamp.Format(EventDateFormatMilli), e.Type, e.GetLevelString(), e.Filename, e.Line, fmt.Sprint(e.Metadata))
}

//StringWithoutCaller returns string representation of an event without information where it was called in code and exactly when (to milliseconds)
func (e *Event) StringWithoutCaller() string {
	return fmt.Sprintf("%s %s %s : %s", e.Timestamp.Format(EventDateFormat), e.Type, e.GetLevelString(), fmt.Sprint(e.Metadata))
}

//String returns string representation of event. If even is of LevelDebug and LevelSilly, it has caller information where it was called in code
func (e *Event) String() string {
	if e.Level >= LevelDebug {
		return e.StringWithCaller()
	}
	return e.StringWithoutCaller()
}

//ToJSON returns json representation of event
func (e *Event) ToJSON() (ret []byte) {
	ret, _ = json.Marshal(e)
	return
}

//ToIndentedJSON returns pretty formated json representation of event
func (e *Event) ToIndentedJSON() (ret []byte) {
	ret, _ = json.MarshalIndent(e, " ", "  ")
	return

}
