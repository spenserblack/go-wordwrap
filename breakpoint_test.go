package wordwrap

import "testing"

// TestDefaultpointLen tests that a break point in the middle of a word would
// have length 0.
func TestDefaultpointLen(t *testing.T) {
	if l := defaultpoint(1).Len(); l != 0 {
		t.Fatalf(`Len = %d, want 0`, l)
	}
}
