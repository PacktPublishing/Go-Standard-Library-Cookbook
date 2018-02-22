package main

import (
	"fmt"
	"math"
)

func main() {

	ln := math.Log(math.E)
	fmt.Printf("Ln(E) = %.4f\n", ln)

	log10 := math.Log10(-100)
	fmt.Printf("Log10(10) = %.4f\n", log10)

	log2 := math.Log2(2)
	fmt.Printf("Log2(2) = %.4f\n", log2)

	log_3_6 := Log(3, 6)
	fmt.Printf("Log3(6) = %.4f\n", log_3_6)

}

// Log computes the logarithm of
// base > 1 and x greater 0
func Log(base, x float64) float64 {
	return math.Log(x) / math.Log(base)
}
