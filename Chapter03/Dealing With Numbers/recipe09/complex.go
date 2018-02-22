package main

import (
	"fmt"
	"math/cmplx"
)

func main() {

	// complex numbers are
	// defined as real and imaginary
	// part defined by float64
	a := complex(2, 3)

	fmt.Printf("Real part: %f \n", real(a))
	fmt.Printf("Complex part: %f \n", imag(a))

	b := complex(6, 4)

	// All common
	// operators are useful
	c := a - b
	fmt.Printf("Difference : %v\n", c)
	c = a + b
	fmt.Printf("Sum : %v\n", c)
	c = a * b
	fmt.Printf("Product : %v\n", c)
	c = a / b
	fmt.Printf("Product : %v\n", c)

	conjugate := cmplx.Conj(a)
	fmt.Println("Complex number a's conjugate : ", conjugate)

	cos := cmplx.Cos(b)
	fmt.Println("Cosine of b : ", cos)

}
