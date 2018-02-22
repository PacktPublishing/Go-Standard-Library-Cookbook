package main

import (
	"fmt"
	"strings"
)

const refString = "Mary had a little lamb"

func main() {
	replacer := strings.NewReplacer("lamb", "wolf", "Mary", "Jack")
	out := replacer.Replace(refString)
	fmt.Println(out)
}
