/*

Package vdlog is a simple and modular logger inspired by
Winston for NodeJS (https://npmjs.org/package/winston).

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

See minimal example for module usage:


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

See full example for module usage:
*/
package vdlog
