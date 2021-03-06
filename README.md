vdlog
======================
[![Build Status](https://travis-ci.org/vodolaz095/vdlog.png?branch=master)](https://travis-ci.org/vodolaz095/vdlog)
[![GoDoc](https://godoc.org/gopkg.in/vodolaz095/vdlog.v2?status.svg)](https://godoc.org/gopkg.in/vodolaz095/vdlog.v2)
[![Go Report Card](https://goreportcard.com/badge/github.com/vodolaz095/vdlog)](https://goreportcard.com/report/github.com/vodolaz095/vdlog)

Structured logger for Go.

Info
======================
Package assumes that there are this event levels:

- Error is a level for errors that can wake you on 4 hours past midnight

- Warn is a level for an unexpected technical or business event happened, customers may be affected, but probably no immediate human intervention is required. On call people won't be called immediately, but support personnel will want to review these issues asap to understand what the impact is. Basically any issue that needs to be tracked but may not require immediate intervention. For example, CPU load is higher than usual, but bot critical.

- Info are for things we want to see at high volume in case we need to forensically analyze an issue. System lifecycle events (system start, stop) go here. "Session" lifecycle events (login, logout, etc.) go here. Significant boundary events should be considered as well (e.g. database calls, remote API calls). Typical business exceptions can go here (e.g. login failed due to bad credentials). Any other event you think you'll need to see in production at high volume goes here.

- Verbose just about everything that doesn't make the "info" cut... any message that is helpful in tracking the flow through the system and isolating issues, especially during the development and QA phases. We use "debug" level logs for entry/exit of most non-trivial methods and marking interesting events and decision points inside methods.

- Debug is for extremely detailed and potentially high volume logs that you don't typically want enabled even during normal development. Examples include dumping a full object hierarchy, logging some state during every iteration of a large loop, etc.

- Silly is putting every fart to log.

After events emitted, they are all send to buffered channel called `spine`.
Than, each event is processed via separate goroutines using Sink functions applied to this event.
You can define sink function easily, see the source of `console.go` and `file.go` files.
If Sink function returns error, it is processed by `BrokenSinkReporter` function.


Log storage support
======================

From the box - STDERR, local file, [Journald](https://wiki.archlinux.org/index.php/Systemd), 
[Loggly](https://loggly.com),  [Telegram](https://core.telegram.org/) channel or group via bot.

With little effort - anything.
You just need to implement simple function that will process events, like this `AddSink` one

```go

	package main

	import (
	  	"fmt"
		"gopkg.in/vodolaz095/vdlog.v3"
	)

	func main(){
		// it is custom sink for events of type `feedback`
		vdlog.AddSink("feedback", func(e vdlog.Event) error {
			fmt.Println("Feedback", e.String())
			return nil
		})

		vdlog.EmitError("test", fmt.Errorf("test %s", "error"), vdlog.H{"error": "error"})
		vdlog.EmitWarn("test", vdlog.H{"warn": "warn"})
		
		//this even will be processed by our defined sink `feedback`
		vdlog.EmitInfo("feedback", vdlog.H{"info": "info"})
		
		vdlog.EmitVerbose("test", vdlog.H{"verbose": "verbose"})
		vdlog.EmitDebug("test", vdlog.H{"debug": "debug"})
		vdlog.EmitSilly("test", vdlog.H{"silly": "silly"})

		//wait until all events are processed
		vdlog.FlushLogs()
	}


```


How to send messages  to Telegram channel or group
======================

1. create telegram bot - see [https://core.telegram.org/bots](https://core.telegram.org/bots).
2. invite this bot as admin of channel or as member of group you want to recieve log entries too.
3. set up log sink using your bot token (something like this - `286759464:AAFRalklssMW9hsZ592O8CxZo63QU7KM7d0`)
4. Get channel/group id using [console telegram client](https://github.com/vysheng/tg) and `channel_info` or `chat_info` commands
5. Find proper id for channel. There is a [crazy solution](https://stackoverflow.com/a/39943226/1885921). For public/private channel you need to prepend `-100` to its id for notifications to be delivered.
6. Enable sink for delivering events

```go

	vdlog.LogToTelegram("286759464:AAFRalklssMW9hsZ592O8CxZo63QU7KM7d0", "-1001055587116", vdlog.LevelInfo)

```




How does logged events looks like (in JSON format)?
======================

```json

      {
        //unique UUID of event, the same reported to all sinks
        "uuid": "990d28fd-4461-480a-be7e-08abacc9bdeb",
        "timestamp": "2017-02-26T19:45:13.398512249+03:00",
  
        //type and level
        "level": 2,
        "levelString": "INFO",
        "type": "vdlogUnitTest",
  
        //event payload, metadata
        "metadata": {
          "someText":"text",
          "someNumber":10,
          "someArray":[0,"1","b"]
        },
  
        //Server info
        "hostname":"server.local",
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

   go get gopkg.in/vodolaz095/vdlog.v3

```

Basic example
======================

```go

	package main

	import (
		"fmt"
		"gopkg.in/vodolaz095/vdlog.v3"
	)

	func main(){
		vdlog.SetConsoleVerbosity(vdlog.LevelSilly)
		vdlog.SetConsoleJSON() //for pretty printing json

		vdlog.EmitError("test", fmt.Errorf("test %s", "error"), vdlog.H{"error": "error"})
		vdlog.EmitWarn("test", vdlog.H{"warn": "warn"})
		vdlog.EmitInfo("feedback", vdlog.H{"info": "info"})
		vdlog.EmitVerbose("test", vdlog.H{"verbose": "verbose"})
		vdlog.EmitDebug("test", vdlog.H{"debug": "debug"})
		vdlog.EmitSilly("test", vdlog.H{"silly": "silly"})


		//wait until all events are processed
		vdlog.FlushLogs()
	}


```


Full example of usage
======================

```go

	package main
	
	import (
	  "fmt"
	  "log"
  	  "gopkg.in/vodolaz095/vdlog.v3"
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
		 * Using global logger
		 */
		vdlog.EmitError("test", fmt.Errorf("test %s", "error"), vdlog.H{"error": "error"})
		vdlog.EmitWarn("test", vdlog.H{"warn": "warn"})
		vdlog.EmitInfo("test", vdlog.H{"info": "info"})
		vdlog.EmitVerbose("test", vdlog.H{"verbose": "verbose"})
		vdlog.EmitDebug("test",vdlog. H{"debug": "debug"})
		vdlog.EmitSilly("test", vdlog.H{"silly": "silly"})

	
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
		log.SetOutput(vdlog.CreateIoWriter(vdlog.LevelError, "test"))
		log.Printf("testing %s", "ioWriterLog")

		//wait until all events are processed
		vdlog.FlushLogs()
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



