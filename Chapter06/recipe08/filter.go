package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {

	for i := 1; i <= 6; i++ {
		_, err := os.Create(fmt.Sprintf("./test.file%d", i))
		if err != nil {
			fmt.Println(err)
		}
	}

	m, err := filepath.Glob("test.file[1-3]")
	if err != nil {
		panic(err)
	}

	for _, val := range m {
		fmt.Println(val)
	}

	// Cleanup
	for i := 1; i <= 6; i++ {
		err := os.Remove(fmt.Sprintf("./test.file%d", i))
		if err != nil {
			fmt.Println(err)
		}
	}
}
