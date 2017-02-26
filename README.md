vdlog
======================
[![Build Status](https://travis-ci.org/vodolaz095/vdlog.png?branch=master)](https://travis-ci.org/vodolaz095/vdlog)
[![GoDoc](https://godoc.org/gopkg.in/vodolaz095/vdlog.v2?status.svg)](https://godoc.org/gopkg.in/vodolaz095/vdlog.v2)
[![Go Report Card](https://goreportcard.com/badge/github.com/vodolaz095/vdlog)](https://goreportcard.com/report/github.com/vodolaz095/vdlog)


Package assumes that there are this event levels:

- Error is a level for errors that can wake you on 4 hours past midnight

- Warn is a level for an unexpected technical or business event happened, customers may be affected, but probably no immediate human intervention is required. On call people won't be called immediately, but support personnel will want to review these issues asap to understand what the impact is. Basically any issue that needs to be tracked but may not require immediate intervention. For example, CPU load is higher than usual, but bot critical.

- Info are for things we want to see at high volume in case we need to forensically analyze an issue. System lifecycle events (system start, stop) go here. "Session" lifecycle events (login, logout, etc.) go here. Significant boundary events should be considered as well (e.g. database calls, remote API calls). Typical business exceptions can go here (e.g. login failed due to bad credentials). Any other event you think you'll need to see in production at high volume goes here.

- Verbose just about everything that doesn't make the "info" cut... any message that is helpful in tracking the flow through the system and isolating issues, especially during the development and QA phases. We use "debug" level logs for entry/exit of most non-trivial methods and marking interesting events and decision points inside methods.

- Debug is for extremely detailed and potentially high volume logs that you don't typically want enabled even during normal development. Examples include dumping a full object hierarchy, logging some state during every iteration of a large loop, etc.

- Silly is putting every fart to log.

After events emitted, they are all send to buffered channel called `spine`.
Than, each event is processed via separate goroutines
using Sink functions applied to this event.
You can define sink function easily, see the source of `console.go` and `file.go` files.
If Sink function returns error, it is processed by `BrokenSinkReporter` function.


Log storage support
======================

From the box - STDERR, local file, [Journald](https://wiki.archlinux.org/index.php/Systemd), 
[Loggly](https://loggly.com). 

With little effort - anything.
You just need to implement simple function that will process events

```go

		vdlog.AddSink("feedback", func(e vdlog.Event) error {
			fmt.Println("Feedback", e.String)
		})


```

How does logged events looks like (in JSON format)?
======================

```javascript

    {
      //unique UUID of event, the same reported to all sinks
      "uuid": "990d28fd-4461-480a-be7e-08abacc9bdeb",


      "timestamp": "2017-02-26T19:45:13.398512249+03:00",


      "level": 2,
      "levelString": "INFO",


      "facility": "vdlogUnitTest",

      "payload": "Hello from vdlog",

      "hostname":"server.local",
	
      //Process ID
      "pid": 11337,

      //Where does the logger function was actually called
      "filename": "/var/www/localhost/index.php",
      "line": 2,
      "called": "/var/www/localhost/index.php:2"
    }

```

Installing
======================

```shell

   go get gopkg.in/vodolaz095/vdlog.v2

```

Basic example
======================

```go

	package main

	import (
		"time"
		"gopkg.in/vodolaz095/vdlog.v2"
	)

	func main(){
		vdlog.SetConsoleVerbosity(LevelSilly)

		vdlog.Sillyf("testFacility", "testing %s", "test")
		vdlog.Verbosef("testFacility", "testing %s", "test")
		vdlog.Debugf("testFacility", "testing %s", "test")
		vdlog.Infof("testFacility", "testing %s", "test")
		vdlog.Warnf("testFacility", "testing %s", "test")
		vdlog.Errorf("testFacility", "testing %s", "test")
		vdlog.Error("testFacility", "Simple string")

		//wait until all events are processed
		time.Sleep(100*time.Millisecond)
	}


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
		 * Logging to Journalctl on local server (works only in linux!)
		 */
		vdlog.LogToLocalJournald()

		/*
		 * Logging to Journalctl on remote server (works only in linux!)
		 */
		vdlog.LogToRemoteJournaldViaTCP("logger.example.org", 514)
		vdlog.LogToRemoteJournaldViaUDP("logger.example.org", 514)

		/*
		 * Logging to Loggly.com
		 */
		vdlog.LogToLoggly("{YOUR LOGGLY TOKEN PLS}", true) //true = https, false = http

	
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



