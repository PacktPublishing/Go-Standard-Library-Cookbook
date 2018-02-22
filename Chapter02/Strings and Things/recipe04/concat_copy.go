package main

import (
	"fmt"
)

func main() {

	strings := []string{"This ", "is ", "even ",
		"more ", "performant "}

	bs := make([]byte, 100)
	bl := 0

	for _, val := range strings {
		bl += copy(bs[bl:], []byte(val))
	}

	fmt.Println(string(bs[:]))

}
