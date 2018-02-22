package main

import (
	"fmt"
	"strings"
)

const selectBase = "SELECT * FROM user WHERE "

var refStringSlice = []string{
	" FIRST_NAME = 'Jack' ",
	" INSURANCE_NO = 333444555 ",
	" EFFECTIVE_FROM = SYSDATE "}

type JoinFunc func(piece string) string

func main() {

	jF := func(p string) string {
		if strings.Contains(p, "INSURANCE") {
			return "OR"
		}

		return "AND"
	}
	result := JoinWithFunc(refStringSlice, jF)
	fmt.Println(selectBase + result)
}

func JoinWithFunc(refStringSlice []string, joinFunc JoinFunc) string {
	concatenate := refStringSlice[0]
	for _, val := range refStringSlice[1:] {
		concatenate = concatenate + joinFunc(val) + val
	}
	return concatenate
}
