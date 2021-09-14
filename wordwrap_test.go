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
