package main

import (
	"fmt"
	"os/exec"
	"runtime"
	"time"
)

func main() {

	var cmd string
	if runtime.GOOS == "windows" {
		cmd = "timeout"
	} else {
		cmd = "sleep"
	}

	proc := exec.Command(cmd, "1")
	proc.Start()

	// Wait function will
	// wait till the process ends.
	proc.Wait()

	// After the process terminates
	// the *os.ProcessState contains
	// simple information
	// about the process run
	fmt.Printf("PID: %d\n", proc.ProcessState.Pid())
	fmt.Printf("Process took: %dms\n", proc.ProcessState.SystemTime()/time.Microsecond)
	fmt.Printf("Exited sucessfuly : %t\n", proc.ProcessState.Success())
}
