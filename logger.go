package vdlog

//Logger is a instance for emitting log events bound to facility
type Logger struct {
	Facility string
}

//Silly makes Logger emit Event with LevelSilly priority
func (l *Logger) Silly(format string, data ...interface{}) {
	log(LevelSilly, l.Facility, format, data...)
}

//Debug makes Logger emit Event with LevelVerbose priority
func (l *Logger) Debug(format string, data ...interface{}) {
	log(LevelDebug, l.Facility, format, data...)
}

//Verbose makes Logger emit Event with LevelDebug priority
func (l *Logger) Verbose(format string, data ...interface{}) {
	log(LevelVerbose, l.Facility, format, data...)
}

//Info makes Logger emit Event with LevelInfo priority
func (l *Logger) Info(format string, data ...interface{}) {
	log(LevelInfo, l.Facility, format, data...)
}

//Warn makes Logger emit Event with LevelWarn priority
func (l *Logger) Warn(format string, data ...interface{}) {
	log(LevelWarn, l.Facility, format, data...)
}

//Error makes Logger emit Event with LevelError priori
func (l *Logger) Error(format string, data ...interface{}) {
	log(LevelError, l.Facility, format, data...)
}

//Log makes Logger emit Event with all things customizeble
func (l *Logger) log(level EventLevel, format string, data ...interface{}) {
	log(level, l.Facility, format, data...)
}

//New creates logger with facility being set
func New(facility string) Logger {
	l := Logger{Facility: facility}
	return l
}
