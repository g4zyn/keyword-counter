package wc

import (
	"strings"
	"unicode"

	"github.com/mgajin/keyword-counter/internal/result"
)

// CountWords
func CountWords(content string) result.Summary {
	// function to detect word separators.
	f := func(r rune) bool { return !unicode.IsLetter(r) }
	// split content into slice of words.
	words := strings.FieldsFunc(content, f)

	summary := make(result.Summary)
	for _, word := range words {
		if res, ok := summary[word]; ok {
			summary[word] = res + 1
		} else {
			summary[word] = 0
		}
	}

	return summary
}
