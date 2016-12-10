package vdlog

/*
 * Syntax sugar - wrappers around log function
 */

//Silly emits Event with LevelSilly priority
func Silly(facility, format string, data ...interface{}) {
	log(LevelSilly, facility, format, data...)
}

//Verbose emits Event with LevelVerbose priority
func Verbose(facility, format string, data ...interface{}) {
	log(LevelVerbose, facility, format, data...)
}

//Debug emits Event with LevelDebug priority
func Debug(facility, format string, data ...interface{}) {
	log(LevelDebug, facility, format, data...)
}

//Info emits Event with LevelInfo priority
func Info(facility, format string, data ...interface{}) {
	log(LevelInfo, facility, format, data...)
}

//Warn emits Event with LevelWarn priority
func Warn(facility, format string, data ...interface{}) {
	log(LevelWarn, facility, format, data...)
}

//Error emits Event with LevelError priority
func Error(facility, format string, data ...interface{}) {
	log(LevelError, facility, format, data...)
}
