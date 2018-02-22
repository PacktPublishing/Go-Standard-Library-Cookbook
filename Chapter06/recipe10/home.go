package main

import (
	"fmt"
	"log"
	"os/user"
)

func main() {
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("The user home directory: " + usr.HomeDir)
}
