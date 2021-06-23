// Package acronym converts a statement to an acronym
// https://golang.org/doc/effective_go.html#commentary
package acronym

import (
	"strings"
	"unicode"
)

// Abbreviate converts a provided statement to an acronym
func Abbreviate(s string) string {
	acronym := ""  // final output
	currWord := "" // current word being built
	for _, c := range s {
		isLetter := unicode.IsLetter(c)
		if isLetter || string(c) == "'" {
			currWord += string(c)
		} else if len(currWord) > 0 {
			acronym += string([]rune(currWord)[0:1])
			currWord = ""
		}
	}

	// include last word in acronym
	acronym += string([]rune(currWord)[0:1])

	return strings.ToUpper(acronym)
}
