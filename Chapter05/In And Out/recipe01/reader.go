package main

import (
	"fmt"
	"os"
)

func main() {

	for {
		data := make([]byte, 8)
		n, err := os.Stdin.Read(data)
		if err == nil && n > 0 {
			process(data)
		} else {
			break
		}
	}

}

func process(data []byte) {
	fmt.Printf("Received: %X 	%s\n", data, string(data))
}
