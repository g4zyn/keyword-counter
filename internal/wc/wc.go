package wc

import (
	"strings"
	"unicode"
)

// Result
type Result map[string]int

// CountWords
func CountWords(content string) Result {
	// function to detect word separators.
	f := func(r rune) bool { return !unicode.IsLetter(r) }
	// split content into slice of words.
	words := strings.FieldsFunc(content, f)

	res := make(Result)
	for _, word := range words {
		if count, ok := res[word]; ok {
			res[word] = count + 1
		} else {
			res[word] = 0
		}
	}

	return res
}
