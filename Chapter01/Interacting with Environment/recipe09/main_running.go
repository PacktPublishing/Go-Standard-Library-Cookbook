package main

import (
	"fmt"
	"os/exec"
	"runtime"
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

	// No process state is returned
	// till the process finish.
	fmt.Printf("Process state for running process: %v\n", proc.ProcessState)

	// The PID could be obtain
	// event for the running process
	fmt.Printf("PID of running process: %d\n\n", proc.Process.Pid)
}
