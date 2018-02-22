package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

func main() {

	// Writing binary values
	buf := bytes.NewBuffer([]byte{})
	if err := binary.Write(buf, binary.BigEndian, 1.004); err != nil {
		panic(err)
	}
	if err := binary.Write(buf, binary.BigEndian, []byte("Hello")); err != nil {
		panic(err)
	}

	// Reading the written values
	var num float64
	if err := binary.Read(buf, binary.BigEndian, &num); err != nil {
		panic(err)
	}
	fmt.Printf("float64: %.3f\n", num)
	greeting := make([]byte, 5)
	if err := binary.Read(buf, binary.BigEndian, &greeting); err != nil {
		panic(err)
	}
	fmt.Printf("string: %s\n", string(greeting))
}
