package vdlog

import (
	"bytes"
	"fmt"
	"os/exec"
)

// Mapping between journalctl and vdlog levels

//    journald            vdlog
//    0: emerg            0:LevelError
//    1: alert            0:LevelError
//    2: crit             0:LevelError
//    3: err              0:LevelError
//    4: warning          1:LevelWarn
//    5: notice           2:LevelInfo
//    6: info             3:LevelVerbose
//    7: debug            4:LevelDebug

func journaldSink(e Event) error {
	//	var journaldPriority int

	//	switch e.Level {
	//	case LevelError:
	//		journaldPriority = 3
	//		break
	//	case LevelWarn:
	//		journaldPriority = 4
	//		break
	//	case LevelInfo:
	//		journaldPriority = 5
	//		break
	//	case LevelVerbose:
	//		journaldPriority = 6
	//		break
	//	case LevelDebug:
	//		journaldPriority = 7
	//		break
	//	default:
	//		journaldPriority = 5
	//	}

	cmd := exec.Command("logger",
		fmt.Sprintf("--tag %s", e.Facility),
		//		fmt.Sprintf("-p %", 5),
		fmt.Sprintf("%q", e.Payload),
	)
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	fmt.Println(cmd.Args, "Stdout", stdout.String(), "Stderr", stderr.String())
	return cmd.Run()
}

//LogToJournald allows to send messages directly to journald daemon via logger command
func LogToJournald() {
	AddSink("journald", journaldSink)
}
