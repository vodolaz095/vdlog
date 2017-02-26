// +build linux

package vdlog

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/satori/go.uuid"
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

func createJournaldSink(host string, port int, tcp, local bool) func(e Event) error {
	return func(e Event) error {
		e.prepare()
		var journaldPriority int
		switch e.Level {
		case LevelError:
			journaldPriority = 3
			break
		case LevelWarn:
			journaldPriority = 4
			break
		case LevelInfo:
			journaldPriority = 5
			break
		case LevelVerbose:
			journaldPriority = 6
			break
		case LevelDebug:
			journaldPriority = 7
			break
		default:
			journaldPriority = 5
		}
		var cmd *exec.Cmd
		if local {
			cmd = exec.Command("/bin/logger", "--journald")
		} else {
			var protoParams string
			if tcp {
				protoParams = "--tcp"
			} else {
				protoParams = "--udp"
			}
			cmd = exec.Command("/bin/logger",
				"--journald",
				fmt.Sprintf("--server %s", host),
				fmt.Sprintf("--port %v", port),
				protoParams,
			)
		}
		stdin, err := cmd.StdinPipe()
		if err != nil {
			return err
		}
		//http://man7.org/linux/man-pages/man7/systemd.journal-fields.7.html
		fmt.Fprintf(stdin, "MESSAGE_ID=%s\n", uuid.NewV4().String())
		fmt.Fprintf(stdin, "CODE_FILE=%s\n", e.Filename)
		fmt.Fprintf(stdin, "CODE_LINE=%v\n", e.Line)
		fmt.Fprintf(stdin, "SYSLOG_IDENTIFIER=%s\n", e.Facility)
		fmt.Fprintf(stdin, "MESSAGE=%q\n", e.Payload)
		fmt.Fprintf(stdin, "PRIORITY=%d\n", journaldPriority)
		fmt.Fprintf(stdin, "SYSLOG_PID=%d\n", os.Getpid())
		fmt.Fprintf(stdin, "OBJECT_PID=%d\n", os.Getpid())
		fmt.Fprintf(stdin, "VDLOG_LEVEL=%s\n", e.LevelString)
		fmt.Fprintf(stdin, "VDLOG_CALLED=%s\n", e.Called)
		err = stdin.Close()
		if err != nil {
			return err
		}
		err = cmd.Run()
		if err != nil {
			return err
		}
		return nil
	}
}

//LogToLocalJournald allows to send messages directly to journald daemon via logger command
func LogToLocalJournald() {
	AddSink("journald local", createJournaldSink("localhost", 514, true, true))
}

//LogToRemoteJournaldViaTCP allows to send messages directly to remote journald daemon via logger command and TCP
func LogToRemoteJournaldViaTCP(server string, port int) {
	AddSink(fmt.Sprintf("journald tcp(%s:%v)", server, port), createJournaldSink(server, port, true, false))
}

//LogToRemoteJournaldViaUDP allows to send messages directly to remote journald daemon via logger command and UDP
func LogToRemoteJournaldViaUDP(server string, port int) {
	AddSink(fmt.Sprintf("journald udp(%s:%v)", server, port), createJournaldSink(server, port, false, false))
}
