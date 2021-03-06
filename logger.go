package vdlog

//Logger is a instance for emitting vdlogEntryPoint events bound to facility
type Logger struct {
	Facility string
}

//New creates logger with facility being set
func New(facility string) Logger {
	l := Logger{Facility: facility}
	return l
}

//EmitSilly makes Logger emit Event with LevelSilly priority
func (l *Logger) EmitSilly(data H) {
	vdlogEntryPoint(LevelSilly, l.Facility, data, nil)
}

//EmitDebug makes Logger emit Event with LevelDebug priority
func (l *Logger) EmitDebug(data H) {
	vdlogEntryPoint(LevelDebug, l.Facility, data, nil)
}

//EmitVerbose makes Logger emit Event with LevelVerbose priority
func (l *Logger) EmitVerbose(data H) {
	vdlogEntryPoint(LevelVerbose, l.Facility, data, nil)
}

//EmitInfo makes Logger emit Event with LevelInfo priority
func (l *Logger) EmitInfo(data H) {
	vdlogEntryPoint(LevelInfo, l.Facility, data, nil)
}

//EmitWarn makes Logger emit Event with LevelWarn priority
func (l *Logger) EmitWarn(data H) {
	vdlogEntryPoint(LevelWarn, l.Facility, data, nil)
}

//EmitError makes Logger emit Event with LevelError priority
func (l *Logger) EmitError(err error, data H) {
	vdlogEntryPoint(LevelError, l.Facility, data, err)
}
