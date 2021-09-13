// Package wordwrap is the main package.
package wordwrap

import (
	"unicode"

	"github.com/mattn/go-runewidth"
)

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

	breakpoints := make([]breakpoint, 0)
	var bp breakpoint
	var currentLen int
	reset := func() {
		bp = nil
		currentLen = 0
	}
	reset()

	runes := []rune(s)
	for i, char := range runes {
		if unicode.IsSpace(char) {
			bp = spacepoint(i)
		} else if char == '-' {
			bp = hyphenpoint(i)
		}
		currentLen += runewidth.RuneWidth(char)

		if currentLen >= limit {
			if bp == nil {
				bp = defaultpoint(i)
			}
			breakpoints = append(breakpoints, bp)
			reset()
		}
	}

	if len(breakpoints) == 0 {
		return []string{s}
	}

	lines = append(lines, string(runes[:breakpoints[0].Start()]))
	if len(breakpoints) == 1 {
		lines = append(lines, string(runes[breakpoints[0].End():]))
		return
	}

	for i, bp := range breakpoints[1:] {
		prev := breakpoints[i]
		lines = append(lines, string(runes[prev.End():bp.Start()]))
	}

	lines = append(lines, string(runes[breakpoints[len(breakpoints)-1].End():]))

	return
}

// Breakpoint represents where a string should be broken to wrap.
type breakpoint interface {
	// Start is where the break should start.
	Start() int
	// End is where the break should end.
	End() int
}

// Spacepoint is a breakpoint triggered by whitespace.
type spacepoint int

func (p spacepoint) Start() int {
	return int(p)
}
func (p spacepoint) End() int {
	return int(p) + 1
}

// Hyphenpoint is a breakpoint triggered by a hyphen.
type hyphenpoint int

func (p hyphenpoint) Start() int {
	return int(p) + 1
}
func (p hyphenpoint) End() int {
	return int(p) + 1
}

// Defaultpoint occurs when there is no available trigger for breakage.
type defaultpoint int

func (p defaultpoint) Start() int {
	return int(p) + 1
}
func (p defaultpoint) End() int {
	return int(p) + 1
}
