package vdlog

/*
 * Syntax sugar - wrappers around vdlogEntryPoint function
 */

//EmitSilly makes Logger emit Event with LevelSilly priority
func EmitSilly(facility string, data H) {
	vdlogEntryPoint(LevelSilly, facility, data, nil)
}

//EmitDebug makes Logger emit Event with LevelDebug priority
func EmitDebug(facility string, data H) {
	vdlogEntryPoint(LevelDebug, facility, data, nil)
}

//EmitVerbose makes Logger emit Event with LevelVerbose priority
func EmitVerbose(facility string, data H) {
	vdlogEntryPoint(LevelVerbose, facility, data, nil)
}

//EmitInfo makes Logger emit Event with LevelInfo priority
func EmitInfo(facility string, data H) {
	vdlogEntryPoint(LevelInfo, facility, data, nil)
}

//EmitWarn makes Logger emit Event with LevelWarn priority
func EmitWarn(facility string, data H) {
	vdlogEntryPoint(LevelWarn, facility, data, nil)
}

//EmitError makes Logger emit Event with LevelError priority
func EmitError(facility string, err error, data H) {
	vdlogEntryPoint(LevelError, facility, data, err)
}
