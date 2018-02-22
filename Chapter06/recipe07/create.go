package main

import (
	"os"
)

func main() {

	f, err := os.Create("created.file")
	if err != nil {
		panic(err)
	}
	f.Close()

	f, err = os.OpenFile("created.byopen", os.O_CREATE|os.O_APPEND, os.ModePerm)
	if err != nil {
		panic(err)
	}
	f.Close()

	err = os.Mkdir("createdDir", 0777)
	if err != nil {
		panic(err)
	}

	err = os.MkdirAll("sampleDir/path1/path2", 0777)
	if err != nil {
		panic(err)
	}

}
