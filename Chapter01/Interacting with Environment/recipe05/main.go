package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}

	// Path to executable file
	fmt.Println(ex)

	// Resolve the direcotry
	// of the executable
	exPath := filepath.Dir(ex)
	fmt.Println("Executable path :" + exPath)

	// Use EvalSymlinks to get
	// the real path.
	realPath, err := filepath.EvalSymlinks(exPath)
	if err != nil {
		panic(err)
	}
	fmt.Println("Symlink evaluated:" + realPath)
}
