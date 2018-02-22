package main

import (
	"fmt"
	"regexp"
)

const refString = "Mary*had,a%little_lamb"

func main() {

	words := regexp.MustCompile("[*,%_]{1}").Split(refString, -1)
	for idx, word := range words {
		fmt.Printf("Word %d is: %s\n", idx, word)
	}

}
