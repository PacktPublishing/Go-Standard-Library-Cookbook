package main

import (
	"fmt"
	"strings"
)

const selectBase = "SELECT * FROM user WHERE %s "

var refStringSlice = []string{
	" FIRST_NAME = 'Jack' ",
	" INSURANCE_NO = 333444555 ",
	" EFFECTIVE_FROM = SYSDATE "}

func main() {

	sentence := strings.Join(refStringSlice, "AND")
	fmt.Printf(selectBase+"\n", sentence)

}
