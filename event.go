package vdlog

import (
	"encoding/json"
	"fmt"
	"time"
)

//EventLevel is a type describing event's level
type EventLevel uint8

const (
	//LevelError is a level for erros that can wake you on 4 hours past midnight
	LevelError = EventLevel(iota)

	//LevelWarn is a level for an unexpected technical or business event happened, customers may be affected, but probably no immediate human intervention is required. On call people won't be called immediately, but support personnel will want to review these issues asap to understand what the impact is. Basically any issue that needs to be tracked but may not require immediate intervention.
	LevelWarn

	//LevelInfo are for things we want to see at high volume in case we need to forensically analyze an issue. System lifecycle events (system start, stop) go here. "Session" lifecycle events (login, logout, etc.) go here. Significant boundary events should be considered as well (e.g. database calls, remote API calls). Typical business exceptions can go here (e.g. login failed due to bad credentials). Any other event you think you'll need to see in production at high volume goes here.
	LevelInfo

	//LevelVerbose just about everything that doesn't make the "info" cut... any message that is helpful in tracking the flow through the system and isolating issues, especially during the development and QA phases. We use "debug" level logs for entry/exit of most non-trivial methods and marking interesting events and decision points inside methods.
	LevelVerbose

	//LevelDebug is for extremely detailed and potentially high volume logs that you don't typically want enabled even during normal development. Examples include dumping a full object hierarchy, logging some state during every iteration of a large loop, etc.
	LevelDebug

	//LevelSilly is putting every fart to log
	LevelSilly

	//EventDateFormat is a constant for formatting date output
	EventDateFormat = "Jan 02 15:04:05"
)

//Event represents anything to be logged
type Event struct {
	Level     EventLevel `json:"level"`
	Facility  string     `json:"facility"`
	Payload   string     `json:"payload"`
	Filename  string     `json:"filename"`
	Line      int        `json:"line"`
	Timestamp time.Time  `json:"timestamp"`
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

//String returns string representation of event without caller information
func (e *Event) String() string {
	return fmt.Sprintf("%s %s %s : %s", e.Timestamp.Format(EventDateFormat), e.Facility, e.GetLevelString(), e.Payload)
}

//StringWithCaller returns string representation of an event with information where it was called in code
func (e *Event) StringWithCaller() string {
	return fmt.Sprintf("%s %s %s <File: %s Line:%v>: %s", e.Timestamp.Format(EventDateFormat), e.Facility, e.GetLevelString(), e.Filename, e.Line, e.Payload)
}

//ToJSON returns json representation of event
func (e *Event) ToJSON() (ret []byte) {
	ret, _ = json.Marshal(e)
	return
}
