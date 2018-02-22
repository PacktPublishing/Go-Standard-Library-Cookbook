package main

import (
	"bytes"
	"fmt"
)

func main() {
	strings := []string{"This ", "is ", "even ",
		"more ", "performant "}
	buffer := bytes.Buffer{}
	for _, val := range strings {
		buffer.WriteString(val)
	}

	fmt.Println(buffer.String())
}
