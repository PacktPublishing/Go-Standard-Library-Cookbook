package main

import (
	"fmt"
	"os"
)

func main() {

	f, err := os.Open("test.file")
	if err != nil {
		panic(err)
	}
	fi, err := f.Stat()
	if err != nil {
		panic(err)
	}

	fmt.Printf("File name: %v\n", fi.Name())
	fmt.Printf("Is Directory: %t\n", fi.IsDir())
	fmt.Printf("Size: %d\n", fi.Size())
	fmt.Printf("Mode: %v\n", fi.Mode())

}
