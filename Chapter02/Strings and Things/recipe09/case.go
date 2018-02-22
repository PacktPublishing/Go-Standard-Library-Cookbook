package main

import (
	"fmt"
	"strings"
	"unicode"
)

const email = "ExamPle@domain.com"
const name = "isaac newton"
const upc = "upc"
const i = "i"

const snakeCase = "first_name"

func main() {

	// For comparing the user input
	// sometimes it is better to
	// compare the input in a same
	// case.
	input := "Example@domain.com"
	input = strings.ToLower(input)
	emailToCompare := strings.ToLower(email)
	matches := input == emailToCompare
	fmt.Printf("Email matches: %t\n", matches)

	upcCode := strings.ToUpper(upc)
	fmt.Println("UPPER case: " + upcCode)

	// This digraph has different upper case and
	// title case.
	str := "ǳ"
	fmt.Printf("%s in upper: %s and title: %s \n",
		str,
		strings.ToUpper(str),
		strings.ToTitle(str))

	// Use of XXXSpecial function
	title := strings.ToTitle(i)
	titleTurk := strings.ToTitleSpecial(unicode.TurkishCase, i)
	if title != titleTurk {
		fmt.Printf("ToTitle is defferent: %#U vs. %#U \n",
			title[0],
			[]rune(titleTurk)[0])
	}

	// In some cases the input
	// needs to be corrected in case.
	correctNameCase := strings.Title(name)
	fmt.Println("Corrected name: " + correctNameCase)

	// Converting the snake case
	// to camel case with use of
	// Title and ToLower functions.
	firstNameCamel := toCamelCase(snakeCase)
	fmt.Println("Camel case: " + firstNameCamel)

}

func toCamelCase(input string) string {
	titleSpace := strings.Title(strings.Replace(input, "_", " ", -1))
	camel := strings.Replace(titleSpace, " ", "", -1)
	return strings.ToLower(camel[:1]) + camel[1:]
}
