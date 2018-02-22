package main

import (
	"fmt"
	"log"
	"os"
)

func main() {

	key := "DB_CONN"
	connStr, ex := os.LookupEnv(key)
	if !ex {
		log.Printf("The env variable %s is not set.\n", key)
	}
	fmt.Println(connStr)
}
