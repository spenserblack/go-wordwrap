package wordwrap

import "testing"

// TestNoLimit tests that a limit of zero or less is considered as an infinite
// maximum length.
func TestNoLimit(t *testing.T) {
	in := "this is a test string"
	out := WordWrap(in, 0)

	if l := len(out); l != 1 {
		t.Fatalf(`len = %d, want 1`, l)
	}
	if word := out[0]; word != in {
		t.Errorf(`word = %q, want %q`, word, in)
	}
}

// TestLimitGreaterThanString tests that a limit that is greater than the input
// string will not break the string up.
func TestLimitGreaterThanString(t *testing.T) {
	in := "this is a test string"
	out := WordWrap(in, 100)

	if l := len(out); l != 1 {
		t.Fatalf(`len = %d, want 1`, l)
	}
	if word := out[0]; word != in {
		t.Errorf(`word = %q, want %q`, word, in)
	}
}

// TestSpacesWithLimit tests that a word with a limit and spaces will be
// wrapped at spaces to fit in that limit.
func TestSpacesWithLimit(t *testing.T) {
	in := "this is a test string"
	out := WordWrap(in, 6)
	expected := []string{"this", "is a", "test", "string"}

	if l := len(out); l > len(expected) {
		t.Errorf(`extra lines after expected output: %v`, out[len(expected):])
	} else if l < len(expected) {
		t.Fatalf(`not enough lines: %#v`, out)
	}

	for i, v := range expected {
		if line := out[i]; line != v {
			t.Errorf(`line %d = %q, want %q`, i+1, line, v)
		}
	}
}

// TestHyphens tests that a word would be wrapped at a hyphen, keeping the
// hyphen in the output.
func TestHyphens(t *testing.T) {
	in := "test-string"
	out := WordWrap(in, 6)

	if l := len(out); l != 2 {
		t.Fatalf(`len = %d, want 2`, l)
	}
	if line, want := out[0], "test-"; line != want {
		t.Errorf(`line 1 = %q, want %q`, line, want)
	}
	if line, want := out[1], "string"; line != want {
		t.Errorf(`line 2 = %q, want %q`, line, want)
	}
}

// TestTooLong tests that a word that is too long to fit within the limit
// will be broken at the limit.
func TestTooLong(t *testing.T) {
	in := "test string"
	out := WordWrap(in, 5)
	expected := []string{"test", "strin", "g"}

	if l := len(out); l > len(expected) {
		t.Errorf(`extra lines after expected output: %v`, out[len(expected):])
	} else if l < len(expected) {
		t.Fatalf(`not enough lines: %#v`, out)
	}

	for i, v := range expected {
		if line := out[i]; line != v {
			t.Errorf(`line %d = %q, want %q`, i+1, line, v)
		}
	}
}

// TestEndOfString tests that a wrap does not occur at the end of a string when
// the length of the string is a multiple of the limit.
//
// In other words, a wrap should not occur when the last line's length equals
// the limit.
func TestEndOfString(t *testing.T) {
	in := "12345"
	out := WordWrap(in, 5)

	if l := len(out); l != 1 {
		t.Errorf(`length = %d, want 1`, l)
	}
	if out[0] != in {
		t.Errorf(`out[0] = %q, want %q`, out[0], in)
	}
}

// TestNoWrapBeforeSpace tests that a wrap does not occur on a space when the
// wrap could occur at a later space and still be within the limit.
//
// In other words, if a space would be the character to put the line over the
// limit, it is *that* space that should be wrapped, not the previous.
func TestNoWrapBeforeSpace(t *testing.T) {
	in := "1 3 5 7 9"
	out := WordWrap(in, 5)
	expected := []string{"1 3 5", "7 9"}

	if l := len(out); l > len(expected) {
		t.Errorf(`extra lines after expected output: %#v`, out[len(expected):])
	} else if l < len(expected) {
		t.Fatalf(`not enough lines: %#v`, out)
	}

	for index, want := range expected {
		if line := out[index]; line != want {
			t.Errorf(`line %d = %q, want %q`, index+1, line, want)
		}
	}
}
