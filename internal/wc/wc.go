package wc

import (
	"strings"
	"unicode"
)

// CountWords
func CountWords(content string) map[string]int {
	// function to detect word separators.
	ff := func(r rune) bool { return !unicode.IsLetter(r) }
	// split content into slice of words.
	words := strings.FieldsFunc(content, ff)

	results := make(map[string]int)
	for _, word := range words {
		if res, ok := results[word]; ok {
			results[word] = res + 1
		} else {
			results[word] = 0
		}
	}

	return results
}
