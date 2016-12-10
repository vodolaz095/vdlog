package vdlog

import (
	"fmt"
	"testing"
	"time"
)

func TestGlobalLog(t *testing.T) {
	var evnt []Event
	AddSink("TestGlobalLogSink", func(e Event) error {
		evnt = append(evnt, e)
		return nil
	})
	Error("test", "error %s", "error")
	Warn("test", "warn %s", "warn")
	Info("test", "info %s", "info")
	Verbose("test", "verbose %s", "verbose")
	Debug("test", "debug %s", "debug")
	Silly("test", "silly %s", "silly")

	time.Sleep(100 * time.Millisecond)

	if len(evnt) != 6 {
		t.Errorf("Wrong number of events emmited - %v instead of %v", len(evnt), 6)
	}

	if evnt[0].Level != LevelError || evnt[0].Payload != "error error" || evnt[0].Facility != "test" {
		t.Error("Wrong Error behavior")
	}

	if evnt[1].Level != LevelWarn || evnt[1].Payload != "warn warn" || evnt[1].Facility != "test" {
		t.Error("Wrong Warn behavior")
	}
	if evnt[2].Level != LevelInfo || evnt[2].Payload != "info info" || evnt[2].Facility != "test" {
		t.Error("Wrong Info behavior")
	}
	if evnt[3].Level != LevelVerbose || evnt[3].Payload != "verbose verbose" || evnt[3].Facility != "test" {
		t.Error("Wrong Error message")
	}
	if evnt[4].Level != LevelDebug || evnt[4].Payload != "debug debug" || evnt[4].Facility != "test" {
		t.Error("Wrong Error message")
	}
	if evnt[5].Level != LevelSilly || evnt[5].Payload != "silly silly" || evnt[5].Facility != "test" {
		t.Error("Wrong Error message")
	}
}

func TestLoggerLog(t *testing.T) {
	var evnt []Event
	AddSink("TestLoggerLogSink", func(e Event) error {
		evnt = append(evnt, e)
		return nil
	})

	logger := New("TestLoggerLog")
	logger.Error("error error") //funny, go vet complains on it :-)
	logger.Warn("warn %s", "warn")
	logger.Info("info %s", "info")
	logger.Verbose("verbose %s", "verbose")
	logger.Debug("debug %s", "debug")
	logger.Silly("silly %s", "silly")

	time.Sleep(100 * time.Millisecond)

	if len(evnt) != 6 {
		t.Errorf("Wrong number of events emmited - %v instead of %v", len(evnt), 6)
	}

	if evnt[0].Level != LevelError || evnt[0].Payload != "error error" || evnt[0].Facility != "TestLoggerLog" {
		t.Error("Wrong Error behavior")
	}

	if evnt[1].Level != LevelWarn || evnt[1].Payload != "warn warn" || evnt[1].Facility != "TestLoggerLog" {
		t.Error("Wrong Warn behavior")
	}
	if evnt[2].Level != LevelInfo || evnt[2].Payload != "info info" || evnt[2].Facility != "TestLoggerLog" {
		t.Error("Wrong Info behavior")
	}
	if evnt[3].Level != LevelVerbose || evnt[3].Payload != "verbose verbose" || evnt[3].Facility != "TestLoggerLog" {
		t.Error("Wrong Error message")
	}
	if evnt[4].Level != LevelDebug || evnt[4].Payload != "debug debug" || evnt[4].Facility != "TestLoggerLog" {
		t.Error("Wrong Error message")
	}
	if evnt[5].Level != LevelSilly || evnt[5].Payload != "silly silly" || evnt[5].Facility != "TestLoggerLog" {
		t.Error("Wrong Error message")
	}
}

func Example() {
	/*
	 *  Set Console Verbosity level
	 */
	SetConsoleVerbosity(LevelSilly) //highest verbosity
	SetConsoleVerbosity(LevelDebug)
	SetConsoleVerbosity(LevelInfo)
	SetConsoleVerbosity(LevelWarn)
	SetConsoleVerbosity(LevelError) //lowest verbosity

	/*
	 * Enable output to local file
	 */
	//LogErrorsToFile outputs errors only
	LogErrorsToFile("/var/log/my_vdlog_errors.log")
	//LogNormalToFile outputs events from error to debug levels
	LogNormalToFile("/var/log/my_log")
	//We can log defined level ranges to file
	LogToFile("/var/log/onlyInfoAndWarn.log", LevelWarn, LevelInfo)

	/*
	 * Add custom sink for storing events
	 * Currently, this sink outputs to STDOUT only events from `feedback`
	 * facility with level lower and including the `LevelInfo`
	 * If Payload equals to `bad`, error is returned
	 */
	AddSink("feedback", func(e Event) error {
		// we ignore events not related for feedback facility
		if e.Facility != "feedback" {
			return nil
		}
		//we ignore events of low priority
		if e.Level > LevelInfo {
			return nil
		}
		//check if payload is the proper one
		if e.Payload == "bad" {
			return fmt.Errorf("bad event")
		}

		//start pretty printing
		fmt.Println("===================")
		fmt.Printf("%v seconds ago event with level %s occured!\n",
			e.Ago().Seconds(),
			e.GetLevelString())

		//Output event as string
		fmt.Println(e.String())
		//will output something like
		// Dec 08 23:49:32 TestLoggerLog VERBOSE : verbose verbose

		//Output event as string with caller information - i.e.
		//where in source code does the message was called
		fmt.Println(e.StringWithCaller())
		//will output something like
		// Dec 08 23:49:32 TestLoggerLog VERBOSE <File: /home/vodolaz095/projects/go/src/bitbucket.org/vodolaz095/vdlog/vdlog_test.go Line:61>: verbose verbose

		//Output JSON representation of message (slice of bytes converted to string)
		fmt.Println("JSON of event:", string(e.ToJSON()))

		fmt.Println("===================")

		//Sink processed event properly
		return nil
	})

	/*
	 * Add function to report sink misbehaviour - i.e. when it returns error
	 */
	BrokenSinkReporter = func(brokenSinkName string, eventThatCloggedIt Event, errorRecievedFromSink error) {
		fmt.Printf("Sink %s is broken by event %s with error %s", brokenSinkName, eventThatCloggedIt.String(), errorRecievedFromSink.Error())
		panic("broken sink")
	}

	/*
	 * Using global logger
	 */
	Silly("testFacility", "testing %s", "test")
	Verbose("testFacility", "testing %s", "test")
	Debug("testFacility", "testing %s", "test")
	Info("testFacility", "testing %s", "test")
	Warn("testFacility", "testing %s", "test")
	Error("testFacility", "testing %s", "test")
	Error("testFacility", "Simple string")

	/*
	 * Using custom logger for `feedback` facility
	 */
	feedbackLogger := New("feedback")
	feedbackLogger.Silly("testing %s", "test")
	feedbackLogger.Verbose("testing %s", "test")
	feedbackLogger.Debug("testing %s", "test")
	feedbackLogger.Info("testing %s", "test")
	feedbackLogger.Warn("testing %s", "test")
	feedbackLogger.Error("testing %s", "test")
	feedbackLogger.Error("Simple string")

	//wait until all events are processed
	time.Sleep(100 * time.Millisecond)
}
