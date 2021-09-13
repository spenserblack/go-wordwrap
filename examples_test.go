package wordwrap_test

import (
	"fmt"

	"github.com/spenserblack/go-wordwrap"
)

// Example is an example for the whole package.
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

// ExampleWordWrap_space is an example for wrapping at spaces.
func ExampleWordWrap_space() {
	lines := wordwrap.WordWrap("we wrap at spaces", 9)
	for _, line := range lines {
		fmt.Println(line)
	}
	// Output:
	// we wrap
	// at spaces
}

// ExampleWordWrap_hyphen is an example for wrapping at hyphens.
func ExampleWordWrap_hyphen() {
	lines := wordwrap.WordWrap("hyphenated-words", 15)
	for _, line := range lines {
		fmt.Println(line)
	}
	// Output:
	// hyphenated-
	// at spaces
}

// ExampleWordWrap_long is an example for wrapping at long words.
func ExampleWordWrap_long() {
	lines := wordwrap.WordWrap("longwordsarewrapped", 10)
	for _, line := range lines {
		fmt.Println(line)
	}
	// Output:
	// longwordsa
	// rewrapped
}
