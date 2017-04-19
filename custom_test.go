package vdlog

import (
	"sync"
	"testing"
	"time"
)

var eventsLogDrain []Event
var mutex *sync.Mutex

func init() {
	mutex = &sync.Mutex{}
	eventsLogDrain = make([]Event, 0)
	AddSink("custom", customLogger)
}

func customLogger(e Event) error {
	if e.Type == "custom" {
		mutex.Lock()
		time.Sleep(10 * time.Millisecond)
		eventsLogDrain = append(eventsLogDrain, e)
		mutex.Unlock()
	}
	return nil
}

func TestCustomLoggerSync(t *testing.T) {
	for i := 0; i < 100; i++ {
		EmitInfo("custom", H{"CustomLogIteration": i})
	}
	time.Sleep(2 * time.Second)
	eventsCreated := len(eventsLogDrain)
	if eventsCreated != 100 {
		t.Errorf("There is %v events instead of 100", eventsCreated)
	}
	for k, v := range eventsLogDrain {
		if v.Metadata["CustomLogIteration"] != k {
			t.Errorf("Wrong event order, it have to be %s instead of %v", v.Metadata["CustomLogIteration"], k)
		}
	}
}

func TestCustomLoggerAsync(t *testing.T) {
	for i := 0; i < 100; i++ {
		go func(a int) {
			EmitInfo("custom", H{"CustomLogIteration": a})
		}(i)
	}
	time.Sleep(2 * time.Second)
	eventsCreated := len(eventsLogDrain)
	if eventsCreated != 200 {
		t.Errorf("There is %v events instead of 200", eventsCreated)
	}
}
