package wordwrap_test

import (
	"fmt"

	"github.com/spenserblack/go-wordwrap"
)

// Words are wrapped at whitespace and at hyphens. Wrapping whitespace is
// trimmed, hyphens are kept. If a word is too long to fit in the limit, it
// will be broken at the limit.
func Example() {
	lines := wordwrap.WordWrap("this test-string has been wrapped successfully", 10)
	for _, line := range lines {
		fmt.Println(line)
	}
	// Output:
	// this
	// test-
	// string
	// has been
	// wrapped
	// successful
	// ly
}

// Words are wrapped at spaces.
func ExampleWordWrap_space() {
	lines := wordwrap.WordWrap("we wrap at spaces", 9)
	for _, line := range lines {
		fmt.Println(line)
	}
	// Output:
	// we wrap
	// at spaces
}

// Words are wrapped at hyphens, and the hyphens are kept at the ends of lines.
func ExampleWordWrap_hyphen() {
	lines := wordwrap.WordWrap("hyphenated-words", 15)
	for _, line := range lines {
		fmt.Println(line)
	}
	// Output:
	// hyphenated-
	// at spaces
}

// Words that are too long for the limit will be broken.
func ExampleWordWrap_long() {
	lines := wordwrap.WordWrap("longwordsarewrapped", 10)
	for _, line := range lines {
		fmt.Println(line)
	}
	// Output:
	// longwordsa
	// rewrapped
}
