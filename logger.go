package vdlog

import (
	"fmt"
)

//Logger is a instance for emitting vdlogEntryPoint events bound to facility
type Logger struct {
	Facility string
}

//Silly makes Logger emit Event with LevelSilly priority
func (l *Logger) Silly(data ...interface{}) {
	vdlogEntryPoint(LevelSilly, l.Facility, fmt.Sprint(data...))
}

//Sillyf makes Logger emit Event with LevelSilly priority with defined format
func (l *Logger) Sillyf(format string, data ...interface{}) {
	vdlogEntryPoint(LevelSilly, l.Facility, format, data...)
}

//Debug makes Logger emit Event with LevelVerbose priority
func (l *Logger) Debug(data ...interface{}) {
	vdlogEntryPoint(LevelDebug, l.Facility, fmt.Sprint(data...))
}

//Debugf makes Logger emit Event with LevelVerbose priority with defined format
func (l *Logger) Debugf(format string, data ...interface{}) {
	vdlogEntryPoint(LevelDebug, l.Facility, format, data...)
}

//Verbose makes Logger emit Event with LevelDebug priority
func (l *Logger) Verbose(data ...interface{}) {
	vdlogEntryPoint(LevelVerbose, l.Facility, fmt.Sprint(data...))
}

//Verbosef makes Logger emit Event with LevelDebug priority
func (l *Logger) Verbosef(format string, data ...interface{}) {
	vdlogEntryPoint(LevelVerbose, l.Facility, format, data...)
}

//Info makes Logger emit Event with LevelInfo priority
func (l *Logger) Info(data ...interface{}) {
	vdlogEntryPoint(LevelInfo, l.Facility, fmt.Sprint(data...))
}

//Infof makes Logger emit Event with LevelInfo priority with defined format
func (l *Logger) Infof(format string, data ...interface{}) {
	vdlogEntryPoint(LevelInfo, l.Facility, format, data...)
}

//Warn makes Logger emit Event with LevelWarn priority
func (l *Logger) Warn(data ...interface{}) {
	vdlogEntryPoint(LevelWarn, l.Facility, fmt.Sprint(data...))
}

//Warnf makes Logger emit Event with LevelWarn priority with defined format
func (l *Logger) Warnf(format string, data ...interface{}) {
	vdlogEntryPoint(LevelWarn, l.Facility, format, data...)
}

//Error makes Logger emit Event with LevelError priority
func (l *Logger) Error(data ...interface{}) {
	vdlogEntryPoint(LevelError, l.Facility, fmt.Sprint(data...))
}

//Errorf makes Logger emit Event with LevelError priority with defined format
func (l *Logger) Errorf(format string, data ...interface{}) {
	vdlogEntryPoint(LevelError, l.Facility, format, data...)
}

//Print allows to print anything as LevelInfo event with payload created by fmt.Print
func (l *Logger) Print(v ...interface{}) {
	vdlogEntryPoint(LevelInfo, l.Facility, fmt.Sprint(v))
}

//Println allows to print anything as LevelInfo event with payload created by fmt.Println
func (l *Logger) Println(v ...interface{}) {
	vdlogEntryPoint(LevelInfo, l.Facility, fmt.Sprintln(v))
}

//Printf allows to print anything as LevelInfo event with payload created by fmt.Printf
func (l *Logger) Printf(format string, v ...interface{}) {
	vdlogEntryPoint(LevelInfo, l.Facility, fmt.Sprintf(format, v))
}

//Log makes Logger emit Event with all things customizable
func (l *Logger) vdlogEntryPoint(level EventLevel, format string, data ...interface{}) {
	vdlogEntryPoint(level, l.Facility, format, data...)
}

//New creates logger with facility being set
func New(facility string) Logger {
	l := Logger{Facility: facility}
	return l
}
