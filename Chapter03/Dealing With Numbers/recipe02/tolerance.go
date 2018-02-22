package main

import (
	"fmt"
	"math"
)

const da = 0.29999999999999998889776975374843459576368331909180
const db = 0.3

func main() {

	daStr := fmt.Sprintf("%.10f", da)
	dbStr := fmt.Sprintf("%.10f", db)

	// While formating the number to string
	// it is rounded to 3.
	fmt.Printf("Strings %s = %s equals: %v \n", daStr, dbStr, dbStr == daStr)

	// Numbers are not equal
	fmt.Printf("Number equals: %v \n", db == da)

	// As the precision of float representation
	// is limited. For the float comparison it is
	// better to use comparison with some tolerance.
	fmt.Printf("Number equals with TOLERANCE: %v \n", Equals(da, db))

}

const TOLERANCE = 1e-8

// Equals compares the floating point numbers
// with tolerance 1e-8
func Equals(numA, numB float64) bool {
	delta := math.Abs(numA - numB)
	if delta < TOLERANCE {
		return true
	}
	return false
}
