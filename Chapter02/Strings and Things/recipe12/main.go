package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func main() {

	text := "Hi! Go is awesome."
	text = Indent(text, 6)
	fmt.Println(text)

	text = Unindent(text, 3)
	fmt.Println(text)

	text = Unindent(text, 10)
	fmt.Println(text)

	text = IndentByRune(text, 10, '.')
	fmt.Println(text)

}

// Indent indenting the input by given indent and rune
func IndentByRune(input string, indent int, r rune) string {
	return strings.Repeat(string(r), indent) + input
}

// Indent indenting the input by given indent
func Indent(input string, indent int) string {
	padding := indent + len(input)
	return fmt.Sprintf("% "+strconv.Itoa(padding)+"s", input)
}

// Unindent unindenting the input string. In case the
// input is indented by less than "indent" spaces
// the min of this both is removed.
func Unindent(input string, indent int) string {

	count := 0
	for _, val := range input {
		if unicode.IsSpace(val) {
			count++
		}
		if count == indent || !unicode.IsSpace(val) {
			break
		}
	}

	return input[count:]
}
