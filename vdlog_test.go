package vdlog

import (
	"fmt"
	"log"
	"strings"
	"testing"
	"time"
)

func TestGlobalLog(t *testing.T) {
	var evnt []Event
	AddSink("TestGlobalLogSink", func(e Event) error {
		evnt = append(evnt, e)
		return nil
	})
	EmitError("test", fmt.Errorf("test %s", "error"), H{"error": "error"})
	EmitWarn("test", H{"warn": "warn"})
	EmitInfo("test", H{"info": "info"})
	EmitVerbose("test", H{"verbose": "verbose"})
	EmitDebug("test", H{"debug": "debug"})
	EmitSilly("test", H{"silly": "silly"})

	log.SetOutput(CreateIoWriter(LevelError, "test"))
	log.Printf("testing %s", "ioWriterLog")
	log.SetPrefix("kuku ")
	log.Printf("testing %s", "ioWriterLog")
	time.Sleep(100 * time.Millisecond)

	if len(evnt) != 8 {
		t.Errorf("Wrong number of events emitted - %v instead of 8", len(evnt))
	}

	if evnt[0].Level != LevelError || evnt[0].Error.Error() != "test error" || evnt[0].Type != "test" {
		t.Error("Wrong Error behavior")
	}

	if evnt[1].Level != LevelWarn || (evnt[1].Metadata["warn"]).(string) != "warn" || evnt[1].Type != "test" {
		t.Error("Wrong Warn behavior")
	}
	if evnt[2].Level != LevelInfo || (evnt[2].Metadata["info"]).(string) != "info" || evnt[2].Type != "test" {
		t.Error("Wrong Info behavior")
	}
	if evnt[3].Level != LevelVerbose || (evnt[3].Metadata["verbose"]).(string) != "verbose" || evnt[3].Type != "test" {
		t.Error("Wrong Error message")
	}
	if evnt[4].Level != LevelDebug || (evnt[4].Metadata["debug"]).(string) != "debug" || evnt[4].Type != "test" {
		t.Error("Wrong Error message")
	}
	if evnt[5].Level != LevelSilly || (evnt[5].Metadata["silly"]).(string) != "silly" || evnt[5].Type != "test" {
		t.Error("Wrong Error message")
	}

	if evnt[6].Level != LevelError || !strings.Contains(evnt[6].Metadata["message"].(string), "testing ioWriterLog") || evnt[6].Type != "test" {
		fmt.Println(evnt[6].Metadata)
		t.Error("Wrong Error message")
	}

	if evnt[7].Level != LevelError || !strings.Contains(evnt[7].Metadata["message"].(string), "kuku testing ioWriterLog") || evnt[7].Type != "test" {
		fmt.Println(evnt[7].Metadata)
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
	logger.EmitError(fmt.Errorf("test %s", "error"), H{"error": "error"})
	logger.EmitWarn(H{"warn": "warn"})
	logger.EmitInfo(H{"info": "info"})
	logger.EmitVerbose(H{"verbose": "verbose"})
	logger.EmitDebug(H{"debug": "debug"})
	logger.EmitSilly(H{"silly": "silly"})

	log.SetOutput(CreateIoWriter(LevelError, "test"))
	log.Printf("testing %s", "ioWriterLog")
	log.SetPrefix("kuku ")
	log.Printf("testing %s", "ioWriterLog")
	time.Sleep(100 * time.Millisecond)

	if len(evnt) != 8 {
		t.Errorf("Wrong number of events emitted - %v instead of 8", len(evnt))
	}

	if evnt[0].Level != LevelError || evnt[0].Error.Error() != "test error" || evnt[0].Type != "TestLoggerLog" {
		t.Error("Wrong Error behavior")
	}

	if evnt[1].Level != LevelWarn || (evnt[1].Metadata["warn"]).(string) != "warn" || evnt[1].Type != "TestLoggerLog" {
		t.Error("Wrong Warn behavior")
	}
	if evnt[2].Level != LevelInfo || (evnt[2].Metadata["info"]).(string) != "info" || evnt[2].Type != "TestLoggerLog" {
		t.Error("Wrong Info behavior")
	}
	if evnt[3].Level != LevelVerbose || (evnt[3].Metadata["verbose"]).(string) != "verbose" || evnt[3].Type != "TestLoggerLog" {
		t.Error("Wrong Error message")
	}
	if evnt[4].Level != LevelDebug || (evnt[4].Metadata["debug"]).(string) != "debug" || evnt[4].Type != "TestLoggerLog" {
		t.Error("Wrong Error message")
	}
	if evnt[5].Level != LevelSilly || (evnt[5].Metadata["silly"]).(string) != "silly" || evnt[5].Type != "TestLoggerLog" {
		t.Error("Wrong Error message")
	}

	if evnt[6].Level != LevelError || !strings.Contains(evnt[6].Metadata["message"].(string), "testing ioWriterLog") || evnt[6].Type != "test" {
		fmt.Println(evnt[6].Metadata)
		t.Error("Wrong Error message")
	}

	if evnt[7].Level != LevelError || !strings.Contains(evnt[7].Metadata["message"].(string), "kuku testing ioWriterLog") || evnt[7].Type != "test" {
		fmt.Println(evnt[7].Metadata)
		t.Error("Wrong Error message")
	}
}

//*/
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
	 * Logging to Journalctl on local server (works only in linux!)
	 */
	LogToLocalJournald()

	/*
	 * Logging to Journalctl on remote server (works only in linux!)
	 */
	LogToRemoteJournaldViaTCP("logger.example.org", 514)
	LogToRemoteJournaldViaUDP("logger.example.org", 514)

	/*
	 * Logging to Loggly.com
	 */
	LogToLoggly("{YOUR LOGGLY TOKEN PLS}", true, LevelSilly) //true = https, false = http

	/*
	 * Send notifications to telegram channel/group chat/personal chat
	 */

	LogToTelegram("286759464:AAFRalklssMW9hsZ592O8CxZo63QU7KM7d0", "-1001055587116", LevelInfo)

	/*
	 * Add custom sink for storing events
	 * Currently, this sink outputs to STDOUT only events from `feedback`
	 * facility with level lower and including the `LevelInfo`
	 * If Payload equals to `bad`, error is returned
	 */
	AddSink("feedback", func(e Event) error {
		// we ignore events not related for feedback facility
		if e.Type != "feedback" {
			return nil
		}
		//we ignore events of low priority
		if e.Level > LevelInfo {
			return nil
		}
		//check if payload is the proper one
		if e.Metadata["payload"].(string) == "bad" {
			return fmt.Errorf("bad event")
		}

		//start pretty printing
		fmt.Println("===================")
		fmt.Printf("%v seconds ago event with level %s occurred!\n",
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

		//Output pretty printed JSON representation of message (slice of bytes converted to string)
		fmt.Println("Indented JSON of event:", string(e.ToIndentedJSON()))

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
	EmitError("test", fmt.Errorf("test %s", "error"), H{"error": "error"})
	EmitWarn("test", H{"warn": "warn"})
	EmitInfo("test", H{"info": "info"})
	EmitVerbose("test", H{"verbose": "verbose"})
	EmitDebug("test", H{"debug": "debug"})
	EmitSilly("test", H{"silly": "silly"})

	/*
	 * Using custom logger for `feedback` facility
	 */
	feedbackLogger := New("feedback")
	feedbackLogger.EmitError(fmt.Errorf("test %s", "error"), H{"error": "error"})
	feedbackLogger.EmitWarn(H{"warn": "warn"})
	feedbackLogger.EmitInfo(H{"info": "info"})
	feedbackLogger.EmitVerbose(H{"verbose": "verbose"})
	feedbackLogger.EmitDebug(H{"debug": "debug"})
	feedbackLogger.EmitSilly(H{"silly": "silly"})

	/*
	 * Using popular https://godoc.org/log package
	 */
	log.SetOutput(CreateIoWriter(LevelError, "test"))
	log.Printf("testing %s", "ioWriterLog")

	//wait until all events are processed
	FlushLogs()
}
