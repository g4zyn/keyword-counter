package wc

import (
	"strings"
	"unicode"
)

// WordCount
type WordCount map[string]int

// CountWords
func CountWords(content string) WordCount {
	// function to detect word separators.
	f := func(r rune) bool { return !unicode.IsLetter(r) }
	// split content into slice of words.
	words := strings.FieldsFunc(content, f)

	wc := make(WordCount)
	for _, word := range words {
		if count, ok := wc[word]; ok {
			wc[word] = count + 1
		} else {
			wc[word] = 0
		}
	}

	return wc
}
