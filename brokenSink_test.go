package vdlog

import (
	"fmt"
	"testing"
	"time"
)

func cloggedSink(e Event) error {
	if e.Type == "toilet" {
		return fmt.Errorf("oops")
	}
	return nil
}

func TestBrokenSync(t *testing.T) {
	called := false

	AddSink("CloggedSink", cloggedSink)

	BrokenSinkReporter = func(brokenSinkName string, eventThatCloggedIt Event, errorRecievedFromSink error) {
		called = true
		if brokenSinkName != "CloggedSink" {
			t.Errorf("BrokenSinkReporter was fired with wrong name of broken sink - %s instead of CloggedSink", brokenSinkName)
		}
		if eventThatCloggedIt.Metadata["message"] != "Большая кала попалась!" {
			t.Errorf("BrokenSinkReporter was fired with wrong event - %s", eventThatCloggedIt.StringWithCaller())
		}
		if errorRecievedFromSink.Error() != "oops" {
			t.Errorf("BrokenSinkReporter was fired with wrong error - %s instead of oops", errorRecievedFromSink.Error())
		}
		fmt.Println("Попалась!")
	}
	EmitInfo("toilet", H{"message": "Большая кала попалась!"})
	time.Sleep(time.Second)
	if !called {
		t.Errorf("BrokenSinkReporter was not called")
	}
}
