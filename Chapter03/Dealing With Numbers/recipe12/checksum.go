package main

import (
	"crypto/md5"
	"fmt"
	"io"
	"os"
)

var content = "This is content to check"

func main() {

	checksum := MD5(content)
	checksum2 := FileMD5("content.dat")

	fmt.Printf("Checksum 1: %s\n", checksum)
	fmt.Printf("Checksum 2: %s\n", checksum2)
	if checksum == checksum2 {
		fmt.Println("Content matches!!!")
	}

}

// MD5 creates the md5
// hash for given content encoded in
// hex string
func MD5(data string) string {
	h := md5.Sum([]byte(data))
	return fmt.Sprintf("%x", h)
}

// FileMD5 creates hex encoded md5 hash
// of file content
func FileMD5(path string) string {
	h := md5.New()
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	_, err = io.Copy(h, f)
	if err != nil {
		panic(err)
	}
	return fmt.Sprintf("%x", h.Sum(nil))
}
