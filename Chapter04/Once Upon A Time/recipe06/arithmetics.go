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
	t := time.Date(2017, 11, 30, 11, 10, 20, 0, l)
	fmt.Printf("Default date is: %v\n", t)

	// Add 3 days
	r1 := t.Add(72 * time.Hour)
	fmt.Printf("Default date +3HRS is: %v\n", r1)

	// Subtract 3 days
	r1 = t.Add(-72 * time.Hour)
	fmt.Printf("Default date -3HRS is: %v\n", r1)

	// More comfortable api
	// to add days/months/years
	r1 = t.AddDate(1, 3, 2)
	fmt.Printf("Default date +1YR +3MTH +2D is: %v\n", r1)
}
