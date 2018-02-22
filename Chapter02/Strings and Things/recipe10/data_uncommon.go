package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func main() {

	file, err := os.Open("data_uncommon.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.Comma = ';'

	for {
		record, e := reader.Read()
		if e != nil {
			fmt.Println(e)
			break
		}
		fmt.Println(record)
	}
}
