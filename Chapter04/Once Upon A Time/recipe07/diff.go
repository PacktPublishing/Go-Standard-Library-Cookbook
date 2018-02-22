package main

import (
	"fmt"
	"time"
)

func main() {

	l, err := time.LoadLocation("Europe/Vienna")
	if err != nil {
		panic(err)
	}

	t := time.Date(2000, 1, 1, 0, 0, 0, 0, l)
	t2 := time.Date(2000, 1, 3, 0, 0, 0, 0, l)
	fmt.Printf("First Default date is %v\n", t)
	fmt.Printf("Second Default date is %v\n", t2)

	dur := t2.Sub(t)
	fmt.Printf("The duration between t and t2 is %v\n", dur)

	dur = time.Since(t)
	fmt.Printf("The duration between now and t is %v\n", dur)

	dur = time.Until(t)
	fmt.Printf("The duration between t and now is %v\n", dur)

}
