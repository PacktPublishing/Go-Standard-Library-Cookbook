package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	// Create the channel where the received
	// signal would be sent. The Notify
	// will not block when the signal
	// is sent and the channel is not ready.
	// So it is better to
	// create buffered channel.
	sChan := make(chan os.Signal, 1)

	// Notify will catch the
	// given signals and send
	// the os.Signal value
	// through the sChan
	signal.Notify(sChan,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT,
		syscall.SIGKILL)

	// Create channel to wait till the
	// signal is handled.
	exitChan := make(chan int)
	go func() {
		signal := <-sChan
		switch signal {
		case syscall.SIGHUP:
			fmt.Println("The calling terminal has been closed")
			exitChan <- 0

		case syscall.SIGINT:
			fmt.Println("The process has been interrupted by CTRL+C")
			exitChan <- 1

		case syscall.SIGTERM:
			fmt.Println("kill SIGTERM was executed for process")
			exitChan <- 1

		case syscall.SIGKILL:
			fmt.Println("SIGKILL handler")
			exitChan <- 1

		case syscall.SIGQUIT:
			fmt.Println("kill SIGQUIT was executed for process")
			exitChan <- 1
		}
	}()

	code := <-exitChan
	os.Exit(code)
}
