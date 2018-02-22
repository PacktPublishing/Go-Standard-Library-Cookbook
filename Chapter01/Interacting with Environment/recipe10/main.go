package main

import (
	"bufio"
	"fmt"
	"io"
	"os/exec"
	"time"
)

func main() {
	cmd := []string{"go", "run", "sample.go"}

	// The command line tool
	// "ping" is executed for
	// 2 seconds
	proc := exec.Command(cmd[0], cmd[1], cmd[2])

	// The process input is obtained
	// in form of io.WriteCloser. The underlying
	// implementation use the os.Pipe
	stdin, _ := proc.StdinPipe()
	defer stdin.Close()

	// For debugging purposes we watch the
	// output of the executed process
	stdout, _ := proc.StdoutPipe()
	defer stdout.Close()

	go func() {
		s := bufio.NewScanner(stdout)
		for s.Scan() {
			fmt.Println("Program says:" + s.Text())
		}
	}()

	// Start the process
	proc.Start()

	// Now the the following lines
	// are written to child
	// process standard input
	fmt.Println("Writing input")
	io.WriteString(stdin, "Hello\n")
	io.WriteString(stdin, "Golang\n")
	io.WriteString(stdin, "is awesome\n")

	time.Sleep(time.Second * 2)

	proc.Process.Kill()

}
