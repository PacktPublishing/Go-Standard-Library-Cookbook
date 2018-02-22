package main

import (
	"fmt"
	"regexp"
)

const refString = "Mary had a little lamb"

func main() {
	regex := regexp.MustCompile("l[a-z]+")
	out := regex.ReplaceAllString(refString, "replacement")
	fmt.Println(out)
}
