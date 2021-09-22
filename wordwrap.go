// Package wordwrap provides a word-wrapping function that returns the wrapped
// lines as a slice of strings.
package wordwrap

import (
	"unicode"

	"github.com/mattn/go-runewidth"
)

// Undefined means that the break point is not defined.
const undefined int = -1

// WordWrap wraps a string, where each "line" of the wrapped string is an item
// of a slice of strings. A limit of zero or less is treated as an infinite
// limit.
//
// If a word is longer than the limit would be allowed even after breaking on
// word separators, then the word itself will be wrapped.
//
// Whitespace such as spaces and tabs are considered word separators that can
// excluded from the output. Hyphens are also considered word separators, but
// will be included at the end of a line.
func WordWrap(s string, limit int) (lines []string) {
	if limit <= 0 {
		return []string{s}
	}

	var (
		startpoint    int = 0
		currentLen    int = 0
		hyphenpoint   int = undefined
		spacepoint    int = undefined
		lastCharIndex int
	)
	reset := func() {
		hyphenpoint = undefined
		spacepoint = undefined
	}

	for i, char := range s {
		charWidth := runewidth.RuneWidth(char)
		if unicode.IsSpace(char) {
			spacepoint = i
		} else if char == '-' {
			hyphenpoint = i
		}
		currentLen += charWidth

		if currentLen > limit {
			var endpoint int
			switch {
			case hyphenpoint != undefined:
				lines = append(lines, s[startpoint:hyphenpoint+1])
				endpoint = hyphenpoint + 1
			case spacepoint != undefined:
				lines = append(lines, s[startpoint:spacepoint])
				endpoint = spacepoint + 1
			default:
				lines = append(lines, s[startpoint:lastCharIndex+1])
				endpoint = lastCharIndex + 1
			}
			remainder := s[endpoint : i+1]
			currentLen = runewidth.StringWidth(remainder)
			startpoint = endpoint
			reset()
		}

		lastCharIndex = i
	}

	if trail := s[startpoint:]; trail != "" {
		lines = append(lines, trail)
	}
	return
}
