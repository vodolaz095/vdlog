vdlog
======================
[![Build Status](https://travis-ci.org/vodolaz095/vdlog.png?branch=master)](https://travis-ci.org/vodolaz095/vdlog)
[![GoDoc](https://godoc.org/gopkg.in/vodolaz095/vdlog.v2?status.svg)](https://godoc.org/gopkg.in/vodolaz095/vdlog.v2)
[![Go Report Card](https://goreportcard.com/badge/github.com/vodolaz095/vdlog)](https://goreportcard.com/report/github.com/vodolaz095/vdlog)


Modular, simple, NodeJS Winston inspired logger for golang

Installing
======================

```shell

   go get gopkg.in/vodolaz095/vdlog.v2

```


Full example of usage
======================

```go

	package main
	
	import (
		"time"
		"gopkg.in/vodolaz095/vdlog.v2"
	)

	func main() {
		/*
		 *  Set Console Verbosity level
		 */
		vdlog.SetConsoleVerbosity(vdlog.LevelSilly) //highest verbosity
		vdlog.SetConsoleVerbosity(vdlog.LevelDebug)
		vdlog.SetConsoleVerbosity(vdlog.LevelInfo)
		vdlog.SetConsoleVerbosity(vdlog.LevelWarn)
		vdlog.SetConsoleVerbosity(vdlog.LevelError) //lowest verbosity
	
		/*
		 * Enable output to local file
		 */
		//LogErrorsToFile outputs errors only
		vdlog.LogErrorsToFile("/var/log/my_vdlog_errors.log")
		//LogNormalToFile outputs events from error to debug levels
		vdlog.LogNormalToFile("/var/log/my_vdlog.log")
		//We can log defined level ranges to file
		vdlog.LogToFile("/var/log/onlyInfoAndWarn.log", vdlog.LevelWarn, vdlog.LevelInfo)
	
		/*
		 * Add custom sink for storing events
		 * Currently, this sink outputs to STDOUT only events from `feedback`
		 * facility with level lower and including the `LevelInfo`
		 * If Payload equals to `bad`, error is returned
		 */
		vdlog.AddSink("feedback", func(e vdlog.Event) error {
			// we ignore events not related for feedback facility
			if e.Facility != "feedback" {
				return nil
			}
			//we ignore events of low priority
			if e.Level > vdlog.LevelInfo {
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
			// Dec 08 23:49:32 TestLoggerLog VERBOSE <File: /home/vodolaz095/projects/go/src/bitbucket.org/vodolaz095/vdlog/vdlog_test.go:61>: verbose verbose
	
			//Output JSON representation of message (slice of bytes converted to string)
			fmt.Println("JSON of event:", string(e.ToJSON()))
	
			fmt.Println("===================")
	
			//Sink processed event properly
			return nil
		})
	
		/*
		 * Add function to report sink misbehaviour - i.e. when it returns error
		 */
		vdlog.BrokenSinkReporter = func(brokenSinkName string, eventThatCloggedIt vdlog.Event, errorRecievedFromSink error) {
			fmt.Printf("Sink %s is broken by event %s with error %s", brokenSinkName, eventThatCloggedIt.String(), errorRecievedFromSink.Error())
			panic("broken sink")
		}
	
		/*
		 * Using global logger ("f" as last letter means formating like fmt.Printf)
		 */
		vdlog.Sillyf("testFacility", "testing %s", "test")
		vdlog.Verbosef("testFacility", "testing %s", "test")
		vdlog.Debugf("testFacility", "testing %s", "test")
		vdlog.Infof("testFacility", "testing %s", "test")
		vdlog.Warnf("testFacility", "testing %s", "test")
		vdlog.Errorf("testFacility", "testing %s", "test")
		vdlog.Error("testFacility", "Simple string")
	
		/*
		 * Using custom logger for `feedback` facility
		 */
		feedbackLogger := vdlog.New("feedback")
		feedbackLogger.Sillyf("testing %s", "test")
		feedbackLogger.Verbosef("testing %s", "test")
		feedbackLogger.Debugf("testing %s", "test")
		feedbackLogger.Infof("testing %s", "test")
		feedbackLogger.Warnf("testing %s", "test")
		feedbackLogger.Errorf("testing %s", "test")
		feedbackLogger.Error("Simple string")

		/*
		 * Using popular https://godoc.org/log package
		 */
		log.SetOutput(CreateIoWriter(LevelError, "test"))
		log.Printf("testing %s", "ioWriterLog")

		/*
		 * Logging to Journalctl on local server (works only in linux!)
		 */
		vdlog.LogToLocalJournald()

		/*
		 * Logging to Journalctl on remote server (works only in linux!)
		 */
		vdlog.LogToRemoteJournaldViaTCP("logger.example.org", 514)
		vdlog.LogToRemoteJournaldViaUDP("logger.example.org", 514)

		//wait until all events are processed
		time.Sleep(100*time.Millisecond)
	}

```


The MIT License (MIT)
==============================

Copyright (c) 2016 Ostroumov Anatolij <ostroumov095 at gmail dot com>

Permission is hereby granted, free of charge, to any person obtaining a copy of
this software and associated documentation files (the "Software"), to deal in
the Software without restriction, including without limitation the rights to
use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of
the Software, and to permit persons to whom the Software is furnished to do so,
subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS
FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR
COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER
IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN
CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.



