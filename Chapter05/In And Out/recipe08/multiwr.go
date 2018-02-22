package main

import "io"
import "bytes"
import "os"
import "fmt"

func main() {

	buf := bytes.NewBuffer([]byte{})
	f, err := os.OpenFile("sample.txt", os.O_CREATE|os.O_RDWR, os.ModePerm)
	if err != nil {
		panic(err)
	}
	wr := io.MultiWriter(buf, f)
	_, err = io.WriteString(wr, "Hello, Go is awesome!")
	if err != nil {
		panic(err)
	}

	fmt.Println("Content of buffer: " + buf.String())
}
