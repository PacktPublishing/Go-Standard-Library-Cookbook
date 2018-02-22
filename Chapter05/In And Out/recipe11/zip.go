package main

import (
	"archive/zip"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
)

func main() {

	var buff bytes.Buffer

	// Compress content
	zipW := zip.NewWriter(&buff)
	f, err := zipW.Create("newfile.txt")
	if err != nil {
		panic(err)
	}
	_, err = f.Write([]byte("This is my file content"))
	if err != nil {
		panic(err)
	}
	err = zipW.Close()
	if err != nil {
		panic(err)
	}

	//Write output to file
	err = ioutil.WriteFile("data.zip", buff.Bytes(), os.ModePerm)
	if err != nil {
		panic(err)
	}

	// Decompress the content
	zipR, err := zip.OpenReader("data.zip")
	if err != nil {
		panic(err)
	}

	for _, file := range zipR.File {
		fmt.Println("File " + file.Name + " contains:")
		r, err := file.Open()
		if err != nil {
			log.Fatal(err)
		}
		_, err = io.Copy(os.Stdout, r)
		if err != nil {
			panic(err)
		}
		err = r.Close()
		if err != nil {
			panic(err)
		}
		fmt.Println()
	}

}
