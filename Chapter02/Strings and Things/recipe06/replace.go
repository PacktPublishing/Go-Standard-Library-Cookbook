package main

import (
	"fmt"
	"strings"
)

const refString = "Mary had a little lamb"
const refStringTwo = "lamb lamb lamb lamb"

func main() {
	out := strings.Replace(refString, "lamb", "wolf", -1)
	fmt.Println(out)

	out = strings.Replace(refStringTwo, "lamb", "wolf", 2)
	fmt.Println(out)
}
