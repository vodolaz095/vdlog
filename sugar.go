package vdlog

import "fmt"

/*
 * Syntax sugar - wrappers around log function
 */

//Silly makes Logger emit Event with LevelSilly priority
func Silly(facility string, data ...interface{}) {
	log(LevelSilly, facility, fmt.Sprint(data...))
}

//Sillyf makes Logger emit Event with LevelSilly priority with defined format
func Sillyf(facility string, format string, data ...interface{}) {
	log(LevelSilly, facility, format, data...)
}

//Debug makes Logger emit Event with LevelVerbose priority
func Debug(facility string, data ...interface{}) {
	log(LevelDebug, facility, fmt.Sprint(data...))
}

//Debugf makes Logger emit Event with LevelVerbose priority with defined format
func Debugf(facility string, format string, data ...interface{}) {
	log(LevelDebug, facility, format, data...)
}

//Verbose makes Logger emit Event with LevelDebug priority
func Verbose(facility string, data ...interface{}) {
	log(LevelVerbose, facility, fmt.Sprint(data...))
}

//Verbosef makes Logger emit Event with LevelDebug priority
func Verbosef(facility string, format string, data ...interface{}) {
	log(LevelVerbose, facility, format, data...)
}

//Info makes Logger emit Event with LevelInfo priority
func Info(facility string, data ...interface{}) {
	log(LevelInfo, facility, fmt.Sprint(data...))
}

//Infof makes Logger emit Event with LevelInfo priority with defined format
func Infof(facility string, format string, data ...interface{}) {
	log(LevelInfo, facility, format, data...)
}

//Warn makes Logger emit Event with LevelWarn priority
func Warn(facility string, data ...interface{}) {
	log(LevelWarn, facility, fmt.Sprint(data...))
}

//Warnf makes Logger emit Event with LevelWarn priority with defined format
func Warnf(facility string, format string, data ...interface{}) {
	log(LevelWarn, facility, format, data...)
}

//Error makes Logger emit Event with LevelError priority
func Error(facility string, data ...interface{}) {
	log(LevelError, facility, fmt.Sprint(data...))
}

//Errorf makes Logger emit Event with LevelError priority with defined format
func Errorf(facility string, format string, data ...interface{}) {
	log(LevelError, facility, format, data...)
}

//Print allows to print anything as LevelInfo event with payload created by fmt.Print
func Print(facility string, v ...interface{}) {
	log(LevelInfo, facility, fmt.Sprint(v))
}

//Println allows to print anything as LevelInfo event with payload created by fmt.Println
func Println(facility string, v ...interface{}) {
	log(LevelInfo, facility, fmt.Sprintln(v))
}

//Printf allows to print anything as LevelInfo event with payload created by fmt.Printf
func Printf(facility string, format string, v ...interface{}) {
	log(LevelInfo, facility, fmt.Sprintf(format, v))
}
